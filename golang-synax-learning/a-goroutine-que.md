# 一个用 goroutines 解决的有趣的题目

## 实现阻塞读且并发安全的map

GO里面MAP如何实现key不存在 get操作等待 直到key存在或者超时，保证并发安全，且需要实现以下接口：

```
type sp interface {
    Out(key string, val interface{})  //存入key /val，如果该key读取的goroutine挂起，则唤醒。此方法不会阻塞，时刻都可以立即执行并返回
    Rd(key string, timeout time.Duration) interface{}  //读取一个key，如果key不存在阻塞，等待key存在或者超时
}
```

解析：
看到阻塞协程第一个想到的就是channel，题目中要求并发安全，那么必须用锁，还要实现多个goroutine读的时候如果值不存在则阻塞，直到写入值，那么每个键值需要有一个阻塞goroutine 的 channel。
实现如下：

```
type Map struct {
    c   map[string]*entry
    rmx *sync.RWMutex
}
type entry struct {
    ch      chan struct{}
    value   interface{}
    isExist bool
}

func (m *Map) Out(key string, val interface{}) {
    m.rmx.Lock()
    defer m.rmx.Unlock()
    if e, ok := m.c[key]; ok {
        e.value = val
        e.isExist = true
        close(e.ch)
    } else {
        e = &entry{ch: make(chan struct{}), isExist: true,value:val}
        m.c[key] = e
        close(e.ch)
    }
}

func (m *Map) Rd(key string, timeout time.Duration) interface{} {
    m.rmx.Lock()
    if e, ok := m.c[key]; ok && e.isExist {
        m.rmx.Unlock()
        return e.value
    } else if !ok {
        e = &entry{ch: make(chan struct{}), isExist: false}
        m.c[key] = e
        m.rmx.Unlock()
        fmt.Println("协程阻塞 -> ", key)
        select {
        case -e.ch:
            return e.value
        case -time.After(timeout):
            fmt.Println("协程超时 -> ", key)
            return nil
        }
    } else {
        m.rmx.Unlock()
        fmt.Println("协程阻塞 -> ", key)
        select {
        case -e.ch:
            return e.value
        case -time.After(timeout):
            fmt.Println("协程超时 -> ", key)
            return nil
        }
    }
}
```

### 思路想改装一下: 不断循环,来执行读写操作

```go
package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

/*
实现阻塞读且并发安全的map
GO里面MAP如何实现key不存在 get操作等待 直到key存在或者超时，保证并发安全，且需要实现以下接口：
*/

type sp interface {
	Out(key string, val interface{})  //存入key /val，如果该key读取的goroutine挂起，则唤醒。此方法不会阻塞，时刻都可以立即执行并返回
	Rd(key string, timeout time.Duration) interface{}  //读取一个key，如果key不存在阻塞，等待key存在或者超时
}

type Map struct {
	c      map[string]*entry
	rmx    *sync.RWMutex
	msg_ch chan *ResultAdd
}

type entry struct {
	ch      chan struct{}
	value   interface{}
	isExist bool
}

type ResultAdd struct {
	key   string
	value interface{}
}

func (m *Map) Out(key string, val interface{}) {
	m.rmx.Lock()
	defer m.rmx.Unlock()
	if e, ok := m.c[key]; ok {
		e.value = val
		e.isExist = true
		close(e.ch)
	} else {
		e = &entry{ch: make(chan struct{}), isExist: true, value: val}
		m.c[key] = e
		close(e.ch)
	}
}

func (m *Map) Rd(key string, timeout time.Duration) interface{} {
	m.rmx.Lock()
	if e, ok := m.c[key]; ok && e.isExist {
		m.rmx.Unlock()
		return e.value
	} else if !ok {
		e = &entry{ch: make(chan struct{}), isExist: false}
		m.c[key] = e
		m.rmx.Unlock()
		fmt.Println("协程阻塞 -> ", key)
		select {
		case <-e.ch:
			return e.value
		case <-time.After(timeout):
			fmt.Println("协程超时 -> ", key)
			// 发送消息到添加管道
			t := &ResultAdd{"hello", e}
			m.msg_ch <- t
			return nil
		}
		return nil
	} else {
		m.rmx.Unlock()
		fmt.Println("协程阻塞 -> ", key)
		select {
		case <-e.ch:
			return e.value
		case <-time.After(timeout):
			fmt.Println("协程超时 -> ", key)
			return nil
		}
	}
}

func main() {
	var lock sync.RWMutex
	s := Map{c: map[string]*entry{}, rmx: &lock}
	s.c = make(map[string]*entry)
	s.c["hello"] = &entry{
		ch:      make(chan struct{}),
		value:   1,
		isExist: true,
	}

	s.msg_ch = make(chan *ResultAdd, 1)
	
	s.Rd("hello2", 5*time.Second)

  for {
			log.Println("aaa")
			select {
			case v := <-s.msg_ch:
				//s.Out(v.key, v.value)
				fmt.Println("result is here", v.key, v.value)
			}
  }
}

// 之后会看到输出报错：
协程阻塞 ->  hello2
2019/04/19 16:22:57 aaa
协程超时 ->  hello2
2019/04/19 16:22:57 aaa
result is here hello &{0xc0000600c0 <nil> false}
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
	/Users/chikitchong/Develop/interview-go/src/q010.go:102 +0x1b1
```

## 原因就是：一个主线程不能有两个阻塞：

### 第一个阻塞 for,第二个阻塞在select case <- xxx,

```go
其实这个问题类似于，主线程两个阻塞
package main

func main() {
	for {
		select {
		
		}
	}
}
```

## 解决的办法就是放在携程里

```go
package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

/*
实现阻塞读且并发安全的map
GO里面MAP如何实现key不存在 get操作等待 直到key存在或者超时，保证并发安全，且需要实现以下接口：
*/

type sp interface {
	Out(key string, val interface{})  //存入key /val，如果该key读取的goroutine挂起，则唤醒。此方法不会阻塞，时刻都可以立即执行并返回
	Rd(key string, timeout time.Duration) interface{}  //读取一个key，如果key不存在阻塞，等待key存在或者超时
}

type Map struct {
	c      map[string]*entry
	rmx    *sync.RWMutex
	msg_ch chan *ResultAdd
}

type entry struct {
	ch      chan struct{}
	value   interface{}
	isExist bool
}

type ResultAdd struct {
	key   string
	value interface{}
}

func (m *Map) Out(key string, val interface{}) {
	m.rmx.Lock()
	defer m.rmx.Unlock()
	if e, ok := m.c[key]; ok {
		e.value = val
		e.isExist = true
		close(e.ch)
	} else {
		e = &entry{ch: make(chan struct{}), isExist: true, value: val}
		m.c[key] = e
		close(e.ch)
	}
}

func (m *Map) Rd(key string, timeout time.Duration) interface{} {
	m.rmx.Lock()
	if e, ok := m.c[key]; ok && e.isExist {
		m.rmx.Unlock()
		return e.value
	} else if !ok {
		e = &entry{ch: make(chan struct{}), isExist: false}
		m.c[key] = e
		m.rmx.Unlock()
		fmt.Println("协程阻塞 -> ", key)
		select {
		case <-e.ch:
			return e.value
		case <-time.After(timeout):
			fmt.Println("协程超时 -> ", key)
			// 发送消息到添加管道
			t := &ResultAdd{"hello", e}
			m.msg_ch <- t
			return nil
		}
		return nil
	} else {
		m.rmx.Unlock()
		fmt.Println("协程阻塞 -> ", key)
		select {
		case <-e.ch:
			return e.value
		case <-time.After(timeout):
			fmt.Println("协程超时 -> ", key)
			return nil
		}
	}
}

func main() {
	var lock sync.RWMutex
	s := Map{c: map[string]*entry{}, rmx: &lock}
	s.c = make(map[string]*entry)
	s.c["hello"] = &entry{
		ch:      make(chan struct{}),
		value:   1,
		isExist: true,
	}

	s.msg_ch = make(chan *ResultAdd, 1)

	s.Rd("hello2", 5*time.Second)
	go func() {
		for {
			log.Println("aaa")
			select {
			case v := <-s.msg_ch:
				//s.Out(v.key, v.value)
				fmt.Println("result is here", v.key, v.value)
			}
		}
	}()
	for ;true;{
		
	}
}
```

