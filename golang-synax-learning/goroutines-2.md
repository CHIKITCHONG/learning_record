## 关于golang的goroutines

### 当主线程 goroutines 在没有结束之前，go 会跑 spinner 子线程
```
package main

import (
	"fmt"
	"time"
)

func main() {
	//go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
		}
	}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
	}

```
[goroutines](https://wizardforcel.gitbooks.io/gopl-zh/ch8/ch8-01.html)
![说明](https://github.com/CHIKITCHONG/learning_record/blob/master/img-src/WechatIMG1.png)
