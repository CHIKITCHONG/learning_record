# 关于 new、make 的区别

### new

```
这是一个用来分配内存的内建函数，它并不初始化内存，只是将其置零。也就是说，new(T)会为T类型的新项目，分配被置零的存储，并且返回它的地址，一个类型为*T的值。在Go的术语中，其返回一个指向新分配的类型为T的指针，这个指针指向的内容的值为零（zero value）。注意并不是指针为零。

### e.g.

// 主要用于初始化结构体
	很多时候，零值并不是一个好主意，我们需要做一些初始化。考虑如下结构体：

    type Rect struct {

    x, y float64

    width, height float64

    }

	零值的Rect并没有多大用处，我们以下方式进行初始化：

	rect3 := &Rect{0, 0, 100, 200}
	rect4 := &Rect{width: 100, height: 200}
```

### make

```
内建函数make(T, args)与new(T)的用途不一样。它只用来创建slice，map和channel，并且返回一个初始化的(而不是置零)，类型为T的值（而不是*T）
```



## 一道题目

```go
package main

import "fmt"

type Param map[string]interface{}

type Show struct {
	db Param
}

// 因为new无法创建字典(没有分配储存结构)
func main1() {
	s := new(Show)
	s.db["RMB"] = 10000
}

// 以下是正确的写法
func main2() {
	s := new(Show)
	s.db = make(map[string]interface{})
	s.db["RMB"] = 10000
	fmt.Println("正确的字典是：", s.db)
}

func main() {
	main1()
    // main1 运行会报错:
    // panic: assignment to entry in nil map
	// goroutine 1 [running]:
    // main.main1(...)
	main2()
}
```

