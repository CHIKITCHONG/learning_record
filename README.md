# learning_record
### 学习记录

 

#### golang websocket develop
- websocket js实现
- golang写服务器响应<br>
## [websocket question eaiser handler](https://github.com/gorilla/websocket/issues)

### [微信websocket实现点餐](https://blog.csdn.net/qqxyy99/article/details/80059888)
### [微信小程序导航：免费视频+精品教程+DEMO集合（长期更新）](https://www.cnblogs.com/wxapp/p/5962642.html)
### [golang+websocket实现实时推送](https://juejin.im/post/5ac0e11f518825555e5dfd1c)
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


<br>
```
# 传和返回结构体都是以*结构体的方式
eg

type Account struct {
	svc *account.Service
}

func NewAccount(svc *account.Service, ar *echo.Group)
```
