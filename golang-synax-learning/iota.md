# golang iota用法记录

### 1, iota << 1

####  << 1 是表示位移，下一位根据上一位的位置来显示 iota 的值（移动当前iota位）

```go
package main
import (
    "fmt"
)
type BitFlag int
const (
  // iota为0，1左移0位 = 1
    Active BitFlag = 1 << iota
  // Send <=> Active <=> 1 << iota，此时iota为1，1左移1位 = 2
    Send
  // Receive <=> Send <=> 1 << iota，此时iota为2，1左移2位 = 4
    Receive
)
func main() {
    fmt.Println(Active, Send, Receive)
}


const (
        bit00 uint32 = 5 << iota           //bit00=5
        bit01                              //bit01=10
        bit02															 //bit02=20
    )

fmt.Printf("bit00 = %d, bit01 = %d, bit02 = %d\n", bit00, bit01, bit02)

//bit00 = 5, bit01 = 10, bit02 = 20
```

### 2. iota是在编译的时候，编译器根据代码中iota与const关键字的位置动态替换的。

```go
package main

import (
    "fmt"
)

const (
    //e=0,f=0,g=0
    e, f, g = iota, iota, iota
)

func main() {
    fmt.Println(e, f, g)
}
```

### 3.可以将iota理解为const语句的行索引

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println(iota)
}

// 编译错误：undefined: iota.
// iota是预先声明的标识符，但是只能作用在const常量声明里。
// 我怎么觉得iota这东西是go的私生子，只能被关在某个地方，不同于true/false等这些兄弟，不能访问它。
```

