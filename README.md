# learning_record

## 问题记录
### 2018-10-16
### golang strust json 序列化操作的一些问题记录

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
```
service/user_test.go

根据service.NewUser的路有要求：NewUser(svc *user.Service, ur *echo.Group)
需要一个新的user的service服务，和一个已分组的echo路由
user.New(nil, tt.udb, tt.rbac, tt.auth)是已经初始化的user的service对象，rg是echo的r路由的分组
```
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


#### 2018-11-19
`git 回退`
1.git log查看日志,拿到commit编号
2.git reset --hard 4b67d6c153666ad57f4311380e8f51b915a73d70


`sql limit与offset`
![pic](https://github.com/CHIKITCHONG/learning_record/blob/master/2018-11-19.png)

#### 2018-11-20
##### 阿里云oss 封禁小程序请求问题
###### 小程序请求图片时，访问阿里云oss资源遭到403forbidden 
![pic](https://github.com/CHIKITCHONG/learning_record/blob/master/2018-11-23-1.jpg)<br>
`解决方法是：`
https://www.zhihu.com/question/63977821
在阿里云设置二级域名映射所有请求，并打开https服务

```
不仅小程序屏蔽阿里云OSS，微信也屏蔽OSS，除了图片可正常打开，其他的word、excel、rar等文件路径都被屏蔽。解决办法就是在阿里云OSS--bucket--域名管理，增加自定义的二级域名路径代替“http://XXXX.oss-cn-hangzhou.aliyuncs.com/”，这样微信就不会屏蔽了。
```
![pic2](https://github.com/CHIKITCHONG/learning_record/blob/master/2018-11-23_2.jpg)
```
---------2018.5.3更新----------
关于https问题，可以购买或申请阿里云免费的“CA证书服务”。在阿里云OSS--bucket--域名管理先增加域名绑定，再点击“证书托管”，填写公钥与私钥。这样资源文件就支持http和https两种方式访问了。
```
![pic3](https://github.com/CHIKITCHONG/learning_record/blob/master/2018-11-23-3.jpg)

##### 知乎上的另一个关于此问题的回答
```
这两天被这个问题烦的头都大了。首先微信必须要求uploadfile的链接必须是https，而且必须是合法域名。而我们公司用的是阿里oss的服务，在添加uploadfile的时候，oss的域名被微信封禁（原因都懂）。所以当时我们想了两个方案，1.上传经过后端ng转一下，这样会有带宽问题 2.阿里云的oss支持配置自定义域名，但是仅支持http的。因为项目急着上线就临时方案1解决。后来晚上回去越想越不对劲，因为我们的服务是支持上传视频的，而且高峰期的qps很高，所以带宽隐患很大。第二天联系了一下阿里云的技术咨询了下，可以通过cdn加速，同时配置cdn的https证书。 最后落实到业务上：需要一个https的业务域名需要一个阿里cdn的cname需要cdn https安全加速需要给cdn配置源站
```
HTTPS安全加速设置_HTTPS安全加速_增值服务_用户指南_CDN-阿里云：https://help.aliyun.com/document_detail/27118.html

#### 2018.11.26更：只需在阿里云oss设置中的refer设置里添加：https://servicewechat.com即可


#### golang 包管理工具 --vendor、gomod
https://ieevee.com/tech/2018/08/28/go-modules.html

#### sql What does DESC do in SQL?
![pic5](https://github.com/CHIKITCHONG/learning_record/blob/master/2018-11-23-4.png)


#### sql limit and offset
![pic6](https://github.com/CHIKITCHONG/learning_record/blob/master/2018-11-23-5.png)



### 2018-11-25

#### linux command review
```
# history 查看在终端输入的命令

# date 查看时间

# env 系统设置的环境变量

# history | grep git 查看所有输入的历史中git 的命令

# mv one.py one.md 把one.py改名为one.md
```
#### gui 编程就前端用异步。。

#### vim cheat sheet(https://vim.rtorr.com/lang/zh_cn/)

##### 一般模式
```
  - 0 :跳到行头
  - shift + $ :跳到行尾
  - G: 移动到文件的最后一行
  - gg:移动到文件首行 
  - :set nu   # 设置行号
  - x:删除光标处的内容
  - dd:删除整行		  # 注：dd是剪切，内容会在剪切板
  - ndd:光标处往下删n行	# eg：3dd,往光标处下删三行
  - u:撤销
  - yy:复制当前行 p:黏贴当前行
```
##### 底部命令模式
```
:set hIsearch    # 搜索模式
	- 按/+key进行关键字向下搜索搜索、按?+key进行关键字向上搜索
	- 搜索到之后按n，N进行 上一个/下一个跳转
set nohIsearch	 # 取消搜索模式（取消高亮）

# 替换
e.g:

160 Languages
161 abc
162 iyoubon
163 century
164 home
# 将160-164的Language替换成mac/g
:160,164s/Languages/mac/g

e.g.2
:1,$s/old/new 	# 将第一行至之后一行的old替换成new
```
##### vim 之块状选择模式
```
crtl + v # 代码块状选择
- ~ 	# 所选块状全大写
- >	# 向右缩进
- < 	# 向左缩进
- d 	# 所选块状删除
```

#### vim 之打开多个文件
```
vim file1 file2   # 按:n\:N切换至上下两个文件
e file		  # 再打开一个文件（注：只显示当前打开的文件）
sp file 	  # 在当前打开的文件中，上下分屏再打开一个文件
vsp file 	  # 左右打开文件
crtl ww 	  # 在不同窗口中切换
:ls 	     	  # 列出来打开的文件
:syntax on 	  # 设置语法高亮
```



### 2018-11-26

#### pgorm 的一些相关东西
##### Select first 20 books:
```
var books []Book
err := db.Model(&books).Order("id ASC").Limit(20).Select()
// SELECT "book"."id", "book"."title", "book"."text"
// FROM "books"
// ORDER BY id ASC LIMIT 20
```

##### Select book by primary key:
```
book := &Book{Id: 1}
err := db.Select(book)
// SELECT "book"."id", "book"."title", "book"."text"
// FROM "books" WHERE id = 1
```
#### 同上，另一种写法
```
book := new(Book)
err := db.Model(book).Where("id = ?", 1).Select()
// SELECT "book"."id", "book"."title", "book"."text"
// FROM "books" WHERE id = 1
```


### 2018-11-27

#### Go命名返回值

```
Go 的返回值可被命名，它们会被视作定义在函数顶部的变量。

返回值的名称应当具有一定的意义，它可以作为文档使用。

没有参数的 return 语句返回已命名的返回值。也就是 直接 返回。

直接返回语句应当仅用在下面这样的短函数中。在长的函数中它们会影响代码的可读性。


package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}

# 输出7、10
```


#### Go defer 栈
`推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。`
```
package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

```

#### Golang cheatSheet!(https://devhints.io/go)



### 2018-11-29

#### golang orm.db 语句：
#### 选择除了软删除的项
`q := db.Model(&banks).Limit(p.Limit).Offset(p.Offset).Where(notDeleted)`

### 2018-12-04

#### golang 时间搓(int)格式化
```
func formatted (t int64) string{
	tm := time.Unix(t, 0)
	result := tm.Format("2006-01-02")
	return result
}
```

### 2018-12-10

#### python 中的枚举类
```
from enum import Enum, unique


@unique             # 保证枚举类的每个对象唯一
class Weekday(Enum):
    Sun = 0  # Sun的value被设定为0
    Mon = 1
    Tue = 2
    Wed = 3
    Thu = 4
    Fri = 5
    Sat = 6
    Set = 0b001000


u = Weekday.Set
print(u.value)


class Gender(Enum):
    Male = 0
    Female = 1


class Student(object):
    def __init__(self, name, gender):
        self.name = name
        self.gender = gender


# 测试:
bart = Student('Bart', Gender.Male)
if bart.gender == Gender.Male:
    print('测试通过!', Gender.Male.value)
else:
    print('测试失败!')


```
### 2018-12-17

#### Golang: How to read a text file?
```
You only need the io/ioutil package.

If you just want the content as string, then the simple solution is to use the ReadFile function from the io/ioutil package. This function returns a slice of bytes which you can easily convert to a string.


package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    b, err := ioutil.ReadFile("file.txt") // just pass the file name
    if err != nil {
        fmt.Print(err)
    }

    fmt.Println(b) // print the content as 'bytes'

    str := string(b) // convert content to a 'string'

    fmt.Println(str) // print the content as a 'string'
}
```
#### 另外一个问题，mac 上如果要清除使用痕迹的话，有什么便捷的方法么

```
recovery 模式直接清除
```


### 2018-12-28

#### PostgreSQL UPSERT 的功能与用法

##### PostgreSQL 9.5 引入了一项新功能，即 UPSERT（insert on conflict do）。当插入遇到约束错误时，直接返回或者改为执行 UPDATE。

###### Example (OnConflictDoUpdate)冲突时更新
```
db := modelDB()

var book *Book
for i := 0; i < 2; i++ {
    book = &Book{
        Id:    100,
        Title: fmt.Sprintf("title version #%d", i),
    }
    _, err := db.Model(book).
        OnConflict("(id) DO UPDATE").
        Set("title = EXCLUDED.title").
        Insert()
    if err != nil {
        panic(err)
    }

    err = db.Select(book)
    if err != nil {
        panic(err)
    }
    fmt.Println(book)
}

err := db.Delete(book)
if err != nil {
    panic(err)
}

```
#### Output
```
Book<Id=100 Title="title version #0">
Book<Id=100 Title="title version #1">
```


###### Example (OnConflictDoNothing)
```
db := modelDB()

book := Book{
    Id:    100,
    Title: "book 100",
}

for i := 0; i < 2; i++ {
    res, err := db.Model(&book).OnConflict("DO NOTHING").Insert()
    if err != nil {
        panic(err)
    }
    if res.RowsAffected() > 0 {
        fmt.Println("created")
    } else {
        fmt.Println("did nothing")
    }
}

err := db.Delete(&book)
if err != nil {
    panic(err)
}
```
##### Output
```
created
did nothing
```

### 2019-1-4
#### range over interface{} which stores a slice
```
Given the scenario where you have a function which accepts t interface{}. If it is determined that the t is a slice, how do I range over that slice? I will not know the incoming type, such as []string, []int or []MyType, at compile time.

func main() {
    data := []string{"one","two","three"}
    test(data)
    moredata := []int{1,2,3}
    test(data)
}

func test(t interface{}) {
    switch reflect.TypeOf(t).Kind() {
    case reflect.Slice:
        // how do I iterate here?
        for _,value := range t {
            fmt.Println(value)
        }
    }
}

## 运行结果是报错: prog.go:16:18: cannot range over t (type interface {})

### 正确的代码
Well I used reflect.ValueOf and then if it is a slice you can call Len() and Index() on the value to get the len of the slice and element at an index. I don't think you will be able to use the range operate to do this.

package main

import "fmt"
import "reflect"

func main() {
    data := []string{"one","two","three"}
    test(data)
    moredata := []int{1,2,3}
    test(moredata)
} 

func test(t interface{}) {
    switch reflect.TypeOf(t).Kind() {
    case reflect.Slice:
        s := reflect.ValueOf(t)

        for i := 0; i < s.Len(); i++ {
            fmt.Println(s.Index(i))
        }
    }
}
```
##### git 冲突
```
<<<<<<<标记冲突开始，后面跟的是当前分支中的内容。

HEAD指向当前分支末梢的提交。

=======之后，>>>>>>>之前是要merge过来的另一条分支上的代码。

>>>>>>>之后的dev是该分支的名字。

对于简单的合并，手工编辑，然后去掉这些标记，最后像往常的提交一样先add再commit即可。

```

##### continue
`continue只在for循环中出现 一旦遇到continue 则跳过进入下一项`
```
func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println("---debug---", i)
	}
}

```
### 2019-1-8

#### git checkout xxx
##### 检查分支是否存在, 没有就新建分支
##### git merge --no-off xx(合并xx分支到当前分支 保留 commit 记录)

#### redis
#### 在线 redis 演示：try.redis.io
```
# 对非空字符串进行 SETRANGE

redis> SET greeting "hello world"
OK

redis> SETRANGE greeting 6 "Redis"
(integer) 11

redis> GET greeting
"hello Redis"


# 对空字符串/不存在的 key 进行 SETRANGE

redis> EXISTS empty_string
(integer) 0

redis> SETRANGE empty_string 5 "Redis!"   # 对不存在的 key 使用 SETRANGE
(integer) 11

redis> GET empty_string                   # 空白处被"\x00"填充
"\x00\x00\x00\x00\x00Redis!"

```
#### 字符串操作(加减)
```
strien xxx redis     (获取key xx 的长度)
getrange xx 1 -1 	(获取key xx 从一到末的长度)
mset key:1 1 key:2 2 (设置多个key)
incr key:1 			(redis 把字符串和数字统称为字符串，可以直接进行加减操作)
append key:1 12
4					(状态成功)
get key:1
"1212"				(append是直接拼接字符串)
```

### 2019-1-11
#### 补充一些 pg 的资料
- 检查是否存在（exist）<br>
`exist,err:=db.Model(&Order{}).Where("out_trade_no = ?",orderNum).Exists()`
- 查询所有 <br>
```
// Select all users.
var users []User
err = db.Model(&users).Select()
if err != nil {
    panic(err)
}
fmt.Println(users)
```
- 通过主键查询
```
// Select user by primary key.
user := &User{Id: 1}
err := db.Select(user)
if err != nil {
    panic(err)
}
fmt.Println(user)
```
#### Select book user WHERE … AND (… OR …):（条件或）

- 关联查询(我们可以按条件关联查询)
```
type User struct {
	Id     int64
	Name   string
	Emails []string
	tableName struct{} `sql:"users"`
}
type Story struct {
	Id       int64
	Title    string
	AuthorId int64
	Author   *User
	tableName struct{} `sql:"storys"`
}

// Select story and associated author in one query
story := new(Story)
err = db.Model(story).
	Relation("Author").
	Where("story.id = ?", 1).
	Select()
if err != nil {
	panic(err)
}
fmt.Println(story)

# 输出
Story<1 Cool story User(1 admin [admin1@admin admin2@admin])>
```
- 关联查询所有
```
type User struct {
	Id     int64
	Name   string
	Emails []string
	tableName struct{} `sql:"users"`
}
type Story struct {
	Id       int64
	Title    string
	AuthorId int64
	Author   *User
	tableName struct{} `sql:"storys"`
}

func main() {
	db:=connet()
	var storys []Story

	err:=db.Model(&storys).Column("story.*").Relation("Author").Select()

	if err!=nil{
		panic(err)
	}
	fmt.Println(storys)
}

## 输出
go=# select * from storys;
 id |   title    | author_id 
----+------------+-----------
  1 | Cool story |         1
  2 | cool       |         1
(2 rows)

## 输出
[Story<1 Cool story User(1 admin [admin1@admin admin2@admin])> Story<2 cool User(1 admin [admin1@admin admin2@admin])>]
```
- pg.IN
```
# 文档中的描述
// In accepts a slice and returns a wrapper that can be used with PostgreSQL
// IN operator:
//
//    Where("id IN (?)", pg.In([]int{1, 2, 3, 4}))
//
// produces
//
//    WHERE id IN (1, 2, 3, 4)

E.G:
func (ShopCartDB) FindWithUidAndCartIds(db orm.DB, uid int, ids []int) ([]*model.UserShopCart, error) {
	carts := make([]*model.UserShopCart, 0)
	err := db.Model(&carts).Column("user_shop_cart.*", "Product").Where("user_id = ? ", uid).
		Where("\"user_shop_cart\".id IN (?)", pg.In(ids)).Select()
	return carts, err
}

## 多重 IN 
func (ShopCartDB) DelWithUidAndCartId(db orm.DB, uid int, ids []int) (bool, error) {
	result, err := db.Model((*model.UserShopCart)(nil)).
		Where("user_id = ? ", uid).
		Where("id IN (?)", pg.In(ids)).
		Delete()
	if err != nil {
		return false, err
	}
	return result.RowsAffected() > 0, nil
}
```

### sql UNION与UNION ALL 命令
```
UNION 操作符用于合并两个或多个 SELECT 语句的结果集。
请注意，UNION 内部的 SELECT 语句必须拥有相同数量的列。列也必须拥有相似的数据类型。同时，每条 SELECT 语句中的列的顺序必须相同。
# 注释：默认地，UNION 操作符选取不同的值。如果允许重复的值，请使用 UNION ALL。

使用 UNION ALL 命令
实例：
列出在中国和美国的所有的雇员：

SELECT E_Name FROM Employees_China
UNION ALL
SELECT E_Name FROM Employees_USA

结果
E_Name
Zhang, Hua
Wang, Wei
Carter, Thomas
Yang, Ming
Adams, John
Bush, George
Carter, Thomas
Gates, Bill
```
### 2019-1-16
#### golang 包管理
```
4. 包管理
通过上文，我们已经知道如何导入包、导入包的几种方式、第三方包的下载、下载包的存放路径等。

那自己创建的项目中如何管理第三方包呢? 即：将用到的第三方包移至个人项目工程下。

govendor 。

虽然存在多种包的管理方式。但我觉得现在的这种方式挺友好的。

下载
go get -u github.com/kardianos/govendor
govendor init
这一步触发的动作是创建：vendor 目录 和 vendor.json 文件

govendor add +external
这一步触发的动作是：将自己项目中使用到的第三方库，复制到vendor 目录下，并维护vendor.json文件

总结：通过本节，你会知道

go 的工作空间的文件系统组成
go 中包的导入方式
go 中第三方包的管理
下节我们写代码，看代码，分析代码。

希望对你有所启发。

再会。
```

#### golang 中channel 的简单例子
##### E.G. 并发中线程的通信机制
```
// 定义goroutine 1
func Echo(out chan<- string) {  					// 定义输出通道类型
	time.Sleep(1*time.Second)
	out <- "咖啡色的羊驼"
	close(out)
}

// 定义goroutine 2
func Receive(out chan<- string, in <-chan string) { // 定义输出通道类型和输入类型
	temp := <-in 				    // 阻塞等待 echo 的通道的返回
	out <- temp				    // 收到 echo 通道弹出的消息后才会执行这里
	close(out)			            // 一环扣一环 通过阻塞 达到进程间的互相联系
}

func main() {
	echo := make(chan string)
	receive := make(chan string)

	go Echo(echo)
	go Receive(receive, echo)

	getStr := <-receive   // 接收goroutine 2的返回

	fmt.Println(getStr)
}

```

#### golang 中 chnnel 死锁示例
```
func say(s string, c chan int) {
    fmt.Printf("%s say\n", s)
    //c <- 1 这里本来应该给c管道传值的，结果没传
}
func main() {
    c := make(chan int)
    go say("lisi", c)
    v1 := <-c //这里会一直阻塞，导致死锁
    fmt.Printf("lisi:%d\n", v1)  //前面死锁，这里无法输出
}
```

#### 接上,golang 中 Channel 的简单例子（调度携程使用函数）
```
# 在使用 gorounite 时，一些函数有返回值，单纯的 go 语句并不会捕获返回值
# 所以需要这么写
package main

import (
    "fmt"
)

//学生结构体(实体)
type Stu struct {
    Name string
    Age  int
}

func say(name string) Stu {
    fmt.Printf("%s say\n", name)
    stu := Stu{Name: name, Age: 18}
    return stu
}
func main() {
    c := make(chan int)
    go func() {
        stu := say("lisi") //返回一个学生实体
        fmt.Printf("我叫%s，年龄%d\n", stu.Name, stu.Age)
        c <- 1 //信号位表示调用完毕
    }()
    fmt.Println("go func")
    <-c
    fmt.Println("end")
}
```
##### linux 中显示 unix 时间：date +%s

### 2019=1-18
##### linux 查找两个文件里的相同以及不同处
```
# 查看两个文件中的不同
grep -v -f a b  从b中剔除a中有的

# 或者
comm -1 A B 不显示在A文件中独有内容(显示B文件独有内容+两个文件共有)
comm -2 A B 不显示在B文件中独有内容
comm -3 A B 不显示同时在两个文件中都存在的内容
comm -12 A B 显示A与B公共的部分
comm -23 A B 显示A独有的
comm -13 A B 显示B独有的

# 或者
diff A B  直接显示两个文件不同，并给出修改一致的建议(主要是对旧文件的建议)
diff -c A B  通过显示两个文件上下文，给出两个文件增减或删除信息，同时也会显示两个文件的修改时间
diff -y A B  模拟将屏幕分成两部分，显示B文件增加或删除的信息
```

### 2019-1-24
#### golang 可变参数
```
# 传递参数给带有可变参数的函数有两种形式，第一种与通常的参数传递没有什么区别，拿上一节的sum举个例子：
# 除了参数的个数是动态变化的之外和普通的函数调用是一致的。
# 第二种形式是使用...运算符以变量...的形式进行参数传递，这里的变量必须是与可变参数类型相同的slice，而不能是其他类型（没错，数组也不可以）：
numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
sum(numbers...) // 和sum(1, 2, 3, 4, 5, 6, 7, 8, 9. 10)等价
```
### 2019-1-25
#### golang 的 panic 与 recover()
```
# golang中没有try... catch...，所以当golang中遇到panic时，如果不进行recover，便会导致整个程序挂掉，具体例子如下：

package main
 
import (
	"fmt"
)
 
func main() {
	panic("fault")
	fmt.Println("panic")
}
运行结果：
panic: fault  
  goroutine 16 [running]:...

# 解决办理：利用defer延迟处理的recover进行恢复，具体例子如下：
package main
 
import (
    "fmt"
)
 
func main() {
    defer func() {
        fmt.Println("1")
    }()
    defer func() {
        if err := recover(); err != nil {
            fmt.Println(err)
        }
    }()
    panic("fault")
    fmt.Println("2")
}
 
运行结果：
  fault
  1
# 程序首先运行panic，出现故障，此时跳转到包含recover()的defer函数执行，recover捕获panic，此时panic就不继续传递．但是recover之后，程序并不会返# 回到panic那个点继续执行以后的动作，而是在recover这个点继续执行以后的动作，即执行上面的defer函数，输出１.

## 注意：利用recover处理panic指令，必须利用defer在panic之前声明，否则当panic时，recover无法捕获到panic，无法防止panic扩散．

# recover是一个内建的函数，可以让进入令人恐慌的流程中的goroutine恢复过来。recover仅在延迟函数中有效。在正常的执行过程中，调用recover会返回nil，并且没有其它任何效果。如果当前的goroutine陷入panic，调用recover可以捕获到panic的输入值，并且恢复正常的执行。
```
### 2019-1-30
#### 一个关于匿名函数的使用
```
# 设定休眠时间三秒
const SCHEDULE_RATE = time.Second * 3 // 调度暂停三秒,减轻压力

# 设置管道,后面用于阻塞消息
var ch = make(chan int)

# 定义一个接受 string, error 并返回 string， error 的函数
func Test(ss string, err error) (string, error) {
	fmt.Println("--- debug ", ss, err)
	return "hello world", nil
}

# 定义了一个 协程+匿名函数
# 这里的函数接受一个 匿名函数(接受string, error为参数)
func AysncTask(f func(string, error)) {
	
	# 这里开一个协程,接受一个匿名函数
	go func(fun func(string, error)) {
		for {
			job, err := Test("hello", nil)
			fmt.Println("--job--", job)
			if len(job) < 0 {
				fun(job, err)
				fmt.Println("****debug****")
			} else {
				time.Sleep(SCHEDULE_RATE)
			}
			//ch <- 1
		}
	}(f)# 匿名函数的参数, 这里的 f 是参数中 f 之后fun(string, error)
}

func main() {
	AysncTask(func(str string, e error) {
		fmt.Println("执行逻辑函数")
	})
	<- ch
}
```
### 2019-2-10
##### linux 中 tail 的用法
```
一、tail命令语法

tail [ -f ] [ -c Number | -n Number | -m Number | -b Number | -k Number ] [ File ]
参数解释：
-f 该参数用于监视File文件增长。
-c Number 从 Number 字节位置读取指定文件
-n Number 从 Number 行位置读取指定文件。
-m Number 从 Number 多字节字符位置读取指定文件，比方你的文件假设包括中文字，假设指定-c参数，可能导致截断，但使用-m则会避免该问题。
-b Number 从 Number 表示的512字节块位置读取指定文件。
-k Number 从 Number 表示的1KB块位置读取指定文件。
File 指定操作的目标文件名称
上述命令中，都涉及到number，假设不指定，默认显示10行。Number前面可使用正负号，表示该偏移从顶部还是从尾部開始计算。
tail可运行文件一般在/usr/bin/以下。

二、tail命令使用方法演示例子

1、tail -f filename
说明：监视filename文件的尾部内容（默认10行，相当于增加参数 -n 10），刷新显示在屏幕上。退出，按下CTRL+C。

2、tail -n 20 filename
说明：显示filename最后20行。

3、tail -r -n 10 filename
说明：逆序显示filename最后10行。
```

### 2019-2-11
#### 关于GOPROXY
```
关于 $GOPROXY

当我们使用go的时候，go默认会直接从代码库中去下载所需的相关依赖，GOPROXY 这个环境变量可以让我们控制自己从哪里去下载源代码，如果 GOPROXY 没有设置，go 会直接从代码库下载相关依赖代码。如果你像下面这样设置了这个环境变量，那么你就会通过 goproxy.io 下载所有的源代码。
# export GOPROXY=https://goproxy.io 

你可以通过置空这个环境变量来关闭，export GOPROXY=
如果要开机、全局生效的就要在/etc/profile中也加入即可。
查看当前系统中的变量,这里过滤GO的配置
[root@nginx-web ~]# export|grep "GO" 
declare -x GO111MODULE="on" 
declare -x GOPROXY="https://goproxy.io" 
declare -x GOROOT="/opt/go" 
```
#### 关于 linux 的"一切皆文件"是什么意思
```
一切皆是文件”是 Unix/Linux 的基本哲学之一，不仅普通的文件，目录、字符设备、块设备、 套接字等在 Unix/Linux 中都是以文件被对待；它们虽然类型不同，但是对其提供的却是同一套操作界面。(同一套操作界面是指对外统一用命令行操作？)
把一切都当成文件来看待，不论你是光驱，还是USB什么的，你都是一个文件。
其实都在于linux的 VFS Virtual File System 虚拟文件系统
```


### 2019-2-12

#### 简述 gRPC 和 protobuf
```
gRPC 一开始由 google 开发，是一款语言中立、平台中立、开源的远程过程调用(RPC)系统。
在 gRPC 里客户端应用可以像调用本地对象一样直接调用另一台不同的机器上服务端应用的方法，使得您能够更容易地创建分布式应用和服务。与许多 RPC 系统类似，gRPC 也是基于以下理念：定义一个服务，指定其能够被远程调用的方法（包含参数和返回类型）。在服务端实现这个接口，并运行一个 gRPC 服务器来处理客户端调用。在客户端拥有一个存根能够像服务端一样的方法。

特性：
基于HTTP/2 
HTTP/2 提供了连接多路复用、双向流、服务器推送、请求优先级、首部压缩等机制。可以节省带宽、降低TCP链接次数、节省CPU，帮助移动设备延长电池寿命等。gRPC 的协议设计上使用了HTTP2 现有的语义，请求和响应的数据使用HTTP Body 发送，其他的控制信息则用Header 表示。

IDL使用ProtoBuf  
IDL(交互式数据语言Interactive Data Language)
gRPC使用ProtoBuf来定义服务，ProtoBuf是由Google开发的一种数据序列化协议（类似于XML、JSON、hessian）。ProtoBuf能够将数据进行序列化，并广泛应用在数据存储、通信协议等方面。压缩和传输效率高，语法简单，表达力强。
多语言支持（C, C++, Python, PHP, Nodejs, C#, Objective-C、Golang、Java） 
```


#### go mod的使用步骤
```
1.首先将你的版本更新到最新的Go版本1.11，如何更新版本可以自行百度。
通过go命令行，进入到你当前的工程目录下，在命令行设置临时环境变量set GO111MODULE=on；

2.执行命令go mod init在当前目录下生成一个go.mod文件，执行这条命令时，当前目录不能存在go.mod文件。如果之前生成过，要先删除；
如果你工程中存在一些不能确定版本的包，那么生成的go.mod文件可能就不完整，因此继续执行下面的命令；

3.执行go mod tidy命令，它会添加缺失的模块以及移除不需要的模块。执行后会生成go.sum文件(模块下载条目)。添加参数-v，例如go mod tidy -v可以将执行的信息，即删除和添加的包打印到命令行；

4.执行命令go mod verify来检查当前模块的依赖是否全部下载下来，是否下载下来被修改过。如果所有的模块都没有被修改过，那么执行这条命令之后，会打印all modules verified。

5.执行命令go mod vendor生成vendor文件夹，该文件夹下将会放置你go.mod文件描述的依赖包，文件夹下同时还有一个文件modules.txt，它是你整个工程的所有模块。在执行这条命令之前，如果你工程之前有vendor目录，应该先进行删除。同理go mod vendor -v会将添加到vendor中的模块打印出来；


另外补充：
go 1.11 有了对模块的实验性支持，大部分的子命令都知道如何处理一个模块，比如 run build install get list mod 子命令，第三方工具可能会支持的晚一些。到 go 1.12 会删除对 GOPATH 的支持，go get 命令也会变成只能获取模块，不能像现在这样直接获取一个裸包。
可以用环境变量 GO111MODULE 开启或关闭模块支持，它有三个可选值：off、on、auto，默认值是 auto。


GO111MODULE=off 无模块支持，go 会从 GOPATH 和 vendor 文件夹寻找包。

GO111MODULE=on 模块支持，go 会忽略 GOPATH 和 vendor 文件夹，只根据 go.mod 下载依赖。   （注意！）

GO111MODULE=auto 在 $GOPATH/src 外面且根目录有 go.mod 文件时，开启模块支持。

在使用模块的时候，GOPATH 是无意义的，不过它还是会把下载的依赖储存在 $GOPATH/pkg/mod 中，也会把 go install 的结果放在 $GOPATH/bin 中。
# https://segmentfault.com/a/1190000016146377 go mod的概述
```


### 2019-2-13
#### 关于 grpc、twirp 两个 grpc 框架
```
grpc、twirp 都是两个grpc框架，其中 grpc 支持 http2,只接收 protobuf,而 twirp 可以接收 json 
序列化还有 protobuf 序列化, 写好 proto 生成文件时，生成了twirp 文件，但并没有使用twirp的网关
功能，而是使用了 twrip 的json序列化或protobuf序列化功能，再经过 grpc
```

### 2019-2-18
#### 关于 go 中 flag 包的使用
```
flag 包主要是实现命令行工具的包
例子：
func main () {
	// 测试使用 flag 包
	flag.String("p", "./config/file/config.dev.yaml", "Path to config file")
	flag.Parse()
}
# 之后再 go build main.go(./main.go)
# 在命令行中执行./main -p
# 可以看见输出：
CHIKITCHONGdeiMac:learnsql chikitchong$ ./main -p
flag needs an argument: -p
Usage of ./main:
  -p string
        Path to config file (default "./config/file/config.dev.yaml")


```

#### 关于 pg 的 pg_dump
```
pg_dump 是 pg 执行脚本的 bash
例子：
### 运行生成本地数据库的 ./cmd/migration/main.go 生成本地数据库
pg_dump -h 127.0.0.1 -p 5432 gcrpc-user > gcrpc_user.sql
```
