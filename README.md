# learning_record
### 学习记录

 

#### golang websocket develop
- websocket js实现
- golang写服务器响应<br>
## main code
### 库相关
1.[这是应该是一个比较简单的 websocket 应用。](https://blog.csdn.net/lhtzbj12/article/details/79357336)<br>
2.[当然这个库貌似比较旧 ，进一步的话可以研究这个库:](https://github.com/gorilla/websocket)
### websocket 教程
1.[阮一峰 WebSocket 教程](http://www.ruanyifeng.com/blog/2017/05/websocket.html)<br>
2.[用 Go 编写一个简单的 WebSocket 推送服务](https://juejin.im/post/5ac0e11f518825555e5dfd1c)<br>
3.[Golang代码搜集-基于websocket+vue.js的简易聊天室](https://blog.csdn.net/lhtzbj12/article/details/79357336)<br>
4.[基于golang的websocket](https://www.jianshu.com/p/65ef71ddb910)<br>
5.[golang echo with gorilla/websocket](http://go-echo.org/cookbook/websocket/)<br>
6.[golang websocket simple example](https://blog.csdn.net/wangshubo1989/article/details/79140278)
<br>
<br>
### 问题记录
`1.net/url里的url.URL是一个结构体 初始化时需要填好相应的字段`
![net/url包url.URL参数](https://github.com/CHIKITCHONG/learning_record/blob/master/WechatIMG20.png)
<br>

Browse activity  原生浏览器<br>
json:"first_name,omitempty" validate:"omitempty,min=2" 意思是为空则不输出

## 2018-10-16
```<br>

#传和返回结构体都是以*结构体的方式

eg:

type Account struct {
	svc *account.Service
}

func NewAccount(svc *account.Service, ar *echo.Group)
```
#### yaml 简述

![yaml](https://github.com/CHIKITCHONG/learning_record/blob/master/20181016.png)
### omitempty 简述
![moitempty](https://github.com/CHIKITCHONG/learning_record/blob/master/20181016-2.png)
### patch 方法简述
![patch](https://github.com/CHIKITCHONG/learning_record/blob/master/20181016-3.png)
## echo 的一些用法记录
```
--1.nocontent 的作用
return c.NoContent(http.StatusOK)`
# 文档描述
NoContent sends a response with no body and a status code.
NoContent(code int) error

--2.c.Request().UserAgent(),    
# 文档描述
// Request() *http.Request


--3.c.RealIP()
# 文档描述
// RealIP returns the client's network address based on `X-Forwarded-For`
```
## 函数中的指针引用
```
type Tools struct {
}

func NewTools(tr *echo.Group) {
	t := &Tools{}   	#去掉{}:Cannot take the address of 'Tools' 

```

## golang strust json 序列化操作的一些问题记录

```
// Product _
type Product struct {
    Name      string  `json:"name"`
    ProductID int64   `json:"-"` // 表示不进行序列化
    Number    int     `json:"number"`
    Price     float64 `json:"price"`
    IsOnSale  bool    `json:"is_on_sale,string"`
}
 
// 序列化过后，可以看见
   {"name":"Xiao mi 6","number":10000,"price":2499,"is_on_sale":"false"}
```
## 20181018
![tupian](https://github.com/CHIKITCHONG/learning_record/blob/master/20181018-1.png)
```
service/user_test.go

根据service.NewUser的路有要求：NewUser(svc *user.Service, ur *echo.Group)
需要一个新的user的service服务，和一个已分组的echo路由
user.New(nil, tt.udb, tt.rbac, tt.auth)是已经初始化的user的service对象，rg是echo的r路由的分组
```
![tupian](https://github.com/CHIKITCHONG/learning_record/blob/master/20181018-2.png)
```
strconv.Atoic.Param("id"),这样捕获的id之后，strconv.Atoic.Param将只接受数字类型，
如果不是的话，就会报错
```
`http://tool.oschina.net/codeformat/json`在线格式化代码


## 20181023
[psql的一些命令](https://blog.csdn.net/smstong/article/details/17138355)
## 20181024
#### 微信小程序不支持PATCH方法，请注意
## 20181025
#### ![tupian](https://github.com/CHIKITCHONG/learning_record/blob/master/20181025-1.jpg)
### redis 取出time.duration 类型可用其.second方法来转换为float64类型来判断是否超时
![tupian](https://github.com/CHIKITCHONG/learning_record/blob/master/20181018-1.png)可用
![tupian](https://github.com/CHIKITCHONG/learning_record/blob/master/20181018-1.png)
## 20181029
golang websocket 微服务<br>
https://blog.csdn.net/yueguanghaidao/article/details/46334483<br>
```
一个专门做推送的库，是哔哩哔哩现在用的弹幕库：
https://github.com/Terry-Mao/goim
https://www.kancloud.cn/liupengjie/go/706202
```
## 2018-10-31
一些题：
#### Short Declairation
```
# Assuming x is declared and y isnot declared, which clauses below are correct?
x; _ := f()
x, _ = f()
x, y := f()
x, y = f()
```
#### Subslice
```
# what will be printed when the code is executed?
package main

import (
   "fmt"
)
func main(){
    s := []int{1,2,3}
    ss := s [1:]
    ss =append(ss, 4)
    
    for _, v :range ss{
        ss[i] += 10
    }
    
    for i := range ss{
        ss[i] += 10
    }
    fmt.Println(s)
}
```
#### Map Ok-Idiom

```
# correct two mistakes in lineA and B

package main

func main(){
    var m map[string]int  //A
    m["a"] = 1
    if v := m["b"]; v!= nil {	//B
    	Println(v)
    }
}
```

#### Pointers

```
# Fill in the blanks in lineA and B,to assure the printed out is "foo"

package main

type S struct {
    m string
}

func f() *S {
    return  _  		//A 
}

func main() {
    p := _   		//B
    print(p, m)		//print "foo"
}
// Todo
```

#### Break Outer Loop
```
# Modify the code below, to exit the outer for loop?

package main 

func main() {
    for i := 0; i < 3; i ++ {
	for j := 0; j < 3; j++ {
	    print(i, ",", j, " ")
	    break
	}
	println()
}
```

#### Defer Stack
```
# Correct a mistake about defer in the code below

package main

import (
    "lo/loutil"
    "os"
)
func main() {
    f, err := os.Open("file")
    defer f.Close()
    if err != nil {
    	return
    }
    b, err := ioutil.ReadAil(f)
    println(string(b))
}
```
## 2018-11-2

post请求 $ curl 127.0.0.1:5342/ws/push -d "id=1&content=hello12"

- string转成int： 
    * int, err := strconv.Atoi(string)
- string转成int64： 
    * int64, err := strconv.ParseInt(string, 10, 64)
- int转成string： 
    * string := strconv.Itoa(int)
- int64转成string： 
    * string := strconv.FormatInt(int64,10)

## 2018-1103
- go在mac上的安装与配置

#### 1
- brew install go之后
    * go env 查看一下go环境看看是否安装好
#### 2
- vim ~/.bash_profile
```
#GOPATH
export GOPATH=/Users/michael2008s/Develop/golang

#GOPATH bin
PATH="$GOPATH/bin:$PATH"
```
#### 3
- $ source ~/.bash_profile
- 完成

## 201811-04

`golang判断key是否在map中`
```
if _, ok := map[key]; ok {
//存在
}

另外golang也没有提供item是否在array当中的判断方法,如果程序里面频繁用到了这种判断,可以将array转化为以array当中的成员为key的map再用上面的方法进行判断,这样会提高判断的效率.
```
#### go语言选择语句 switch case

```
switch i { 
    case 0: 
        fmt.Printf("0") 
    case 1: 
        fmt.Printf("1") 
    case 2: 
        fallthrough 
    case 3: 
        fmt.Printf("3") 
    case 4, 5, 6: 
        fmt.Printf("4, 5, 6") 
    default: 
        fmt.Printf("Default") 
}

运行上面的案例，将会得到如下结果： 
 i = 0时，输出0； 
 i = 1时，输出1； 
 i = 2时，输出3；

 i = 3时，输出3； 

 i = 4时，输出4, 5, 6； 
 i = 5时，输出4, 5, 6； 
 i = 6时，输出4, 5, 6； 
 i = 其他任意值时，输出Default
```

#### 2018-11-06

##### curl 127.0.0.1:5342/ws/push?id=1 -d "orderNumber=100&storekeeperID=1&ActualPrice=666&OrderStatus=888&PaymentStatus=999&StoreName=深圳湾1号&DistributorName=xxx"


#### 2018-11-07
`让外网可以访问本地的好软件：https://ngrok.com/docs#global`

#### 2018-11-08
linux的grep命令`https://blog.csdn.net/tenfyguo/article/details/6387786`
网关限速`https://www.krakend.io/`
docker镜像下载 https://hub.docker.com/
