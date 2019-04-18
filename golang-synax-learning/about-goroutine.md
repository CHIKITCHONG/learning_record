### 关于 goroutine

摘自 <https://studygolang.com/articles/12972?fr=sidebar>, 侵删

经常会看到以下了代码：

```
package main

import (
    "fmt"
    "time"
)

func main(){
    for i := 0; i < 100 ; i++{
        go fmt.Println(i)
    }
    time.Sleep(time.Second)
}
```

主线程为了等待goroutine都运行完毕，不得不在程序的末尾使用`time.Sleep()` 来睡眠一段时间，等待其他线程充分运行。对于简单的代码，100个for循环可以在1秒之内运行完毕，`time.Sleep()` 也可以达到想要的效果。

但是对于实际生活的大多数场景来说，1秒是不够的，并且大部分时候我们都无法预知for循环内代码运行时间的长短。这时候就不能使用`time.Sleep()` 来完成等待操作了。

可以考虑使用管道来完成上述操作：

```
func main() {
    c := make(chan bool, 100)
    for i := 0; i < 100; i++ {
        go func(i int) {
            fmt.Println(i)
            c <- true
        }(i)
    }

    for i := 0; i < 100; i++ {
        <-c
    }
}
```

首先可以肯定的是使用管道是能达到我们的目的的，而且不但能达到目的，还能十分完美的达到目的。

但是管道在这里显得有些大材小用，因为它被设计出来不仅仅只是在这里用作简单的同步处理，在这里使用管道实际上是不合适的。而且假设我们有一万、十万甚至更多的for循环，也要申请同样数量大小的管道出来，对内存也是不小的开销。

对于这种情况，go语言中有一个其他的工具`sync.WaitGroup` 能更加方便的帮助我们达到这个目的。

`WaitGroup` 对象内部有一个计数器，最初从0开始，它有三个方法：`Add(), Done(), Wait()` 用来控制计数器的数量。`Add(n)` 把计数器设置为`n` ，`Done()` 每次把计数器`-1` ，`wait()` 会阻塞代码的运行，直到计数器地值减为0。

使用`WaitGroup` 将上述代码可以修改为：

```
func main() {
    wg := sync.WaitGroup{}
    wg.Add(100)
    for i := 0; i < 100; i++ {
        go func(i int) {
            fmt.Println(i)
            wg.Done()
        }(i)
    }
    wg.Wait()
}
```

这里首先把`wg` 计数设置为100， 每个for循环运行完毕都把计数器减一，主函数中使用`Wait()` 一直阻塞，直到wg为零——也就是所有的100个for循环都运行完毕。相对于使用管道来说，`WaitGroup` 轻巧了许多。

### 注意：

###  计数器不能为负值

我们不能使用`Add()` 给`wg` 设置一个负值，否则代码将会报错：

```
panic: sync: negative WaitGroup counter

goroutine 1 [running]:
sync.(*WaitGroup).Add(0xc042008230, 0xffffffffffffff9c)
    D:/Go/src/sync/waitgroup.go:75 +0x1d0
main.main()
    D:/code/go/src/test-src/2-Package/sync/waitgroup/main.go:10 +0x54
```

同样使用`Done()` 也要特别注意不要把计数器设置成负数了。

### WaitGroup对象不是一个引用类型

WaitGroup对象不是一个引用类型，在通过函数传值的时候需要使用地址：

```
func main() {
    wg := sync.WaitGroup{}
    wg.Add(100)
    for i := 0; i < 100; i++ {
        go f(i, &wg)
    }
    wg.Wait()
}

// 一定要通过指针传值，不然进程会进入死锁状态
func f(i int, wg *sync.WaitGroup) { 
    fmt.Println(i)
    wg.Done()
}
```

### 这里面补充一下, golang 的 select 语法

（作者：chaors，链接：https://www.jianshu.com/p/2a1146dc42c3，来源：简书，侵删）

首先，我们来从官方文档看一下有关select的描述：

A "select" statement chooses which of a set of possible send or receive operations will proceed. It looks similar to a "switch" statement but with the cases all referring to communication operations.
 一个select语句用来选择哪个case中的发送或接收操作可以被立即执行。它类似于switch语句，但是它的case涉及到channel有关的I/O操作。

或者换一种说法，select就是用来监听和channel有关的IO操作，当 IO 操作发生时，触发相应的动作。

### 基本用法

```
//select基本用法
select {
case <- chan1:
// 如果chan1成功读到数据，则进行该case处理语句
case chan2 <- 1:
// 如果成功向chan2写入数据，则进行该case处理语句
default:
// 如果上面都没有成功，则进入default处理流程
```



