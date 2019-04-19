## 这里主要收集一些与 go 相关的东西

### go rune 简要分析

rune在golang中是int32的别名，在各个方面都与int32相同。
被用来区分字符值和整数值。

```
s:="hello你好"
fmt.Println(len(s))//输出长度为11
fmt.Println(len([]rune(s)))//输出长度为7
s="你好"
fmt.Println(len(s))//输出长度为6
fmt.Println(len([]rune(s)))//输出长度为2
s="你"
fmt.Println([]byte(s))//输出长度为6
fmt.Println(rune('你'))//输出20320
--------------------- 
作者：中国流浪猫 
来源：CSDN 
原文：https://blog.csdn.net/a41888313/article/details/78946911 
版权声明：本文为博主原创文章，转载请附上博文链接！
```

通过上述代码可以将rune理解为 一个 可以表示unicode 编码的值int 的值，称为码点（code point）。只不过go语言把这个码点抽象为rune。
最后，rune 就相当于java中 char类型，只不过go作为国际化语言，char 只有2个字节，go本身是utf-8编码的，char不能满足要求，所有go就有rune这个类型。