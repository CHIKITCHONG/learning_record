package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"net/http"
	"sort"
	"strings"
	"time"
)

const (
	ServerID = "lhtzbj12@126.com"
)

/*

需要一个 请求、发送信息类别 判断状态
//MesssageReq 发送消息请求
type MesssageReq struct {
	Type string  `json:"type"` //请求的类别 login表示登录  msg表消息 users用户列表
	Data Message `json:"data"` //用户发送的信息
}


断连---需要持续连接

发送函数


任务调度{

	-- 收信息
	-- 发信息

}
 */


//LoginReq 登录请求
type LoginReq struct {
	Type string   `json:"type"` //请求的类别 login表示登录  msg表消息 users用户列表
	Data UserInfo `json:"data"` //用户信息
}

//LoginRespData 登录返回
type LoginRespData struct {
	Result int    `json:"result"`
	Msg    string `json:"msg"`
}

//MesssageReq 发送消息请求
type MesssageReq struct {
	Type string  `json:"type"` //请求的类别 login表示登录  msg表消息 users用户列表
	Data Message `json:"data"` //用户发送的信息
}

//TransferData 数据传输结构体
type TransferData struct {
	Type string      `json:"type"` //数据类别 login表示登录  msg表消息 users用户列表
	Data interface{} `json:"data"` //数据
}

//Message 消息的结构体
type Message struct {
	From string `json:"from"` //来自谁 email
	To   string `json:"to"`   //发给谁 email 空表示表示所有人
	Time string `json:"time"` //消息发出的时间
	Cont string `json:"cont"` //消息内容
}

//UserInfo 用户基本信息结构体
type UserInfo struct {
	NickName    string            `json:"nickname"` //昵称
	Email       string            `json:"email"`    //用户邮箱
	Send2Client chan TransferData `json:"-"`        //发给客户端的数据
}

//RoomInfo 房间信息
type RoomInfo struct {
	RoomName    string     `json:"roomname"`    //房间名称
	OnlineNum   int        `json:"onlinenum"`   //在线人数
	OnlineUsers []UserInfo `json:"onlineusers"` //在线用户列表
	ServerID    string     `json:"serverid"`    //服务器标识
}

var (
	users         = make(map[string]UserInfo)   //用户列表
	entering      = make(chan UserInfo)         //用户进入
	leaving       = make(chan UserInfo)         //用户离开
	transferDatas = make(chan TransferData, 10) //广播消息队列
)

func checkerr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

//解析数据包里的type
func getType(str string) string {
	key := "\"type\":"
	index := strings.Index(str, key)
	str = str[index+len(key)+1:]
	index = strings.Index(str, "\"")
	return str[0:index]
}


//任务调度 用户进入、离开、消息分发
func taskSchedule() {
	//广播聊天室信息
	bcroominfo := make(chan int, 1)
	for {
		select {

		case u := <-transferDatas:
			fmt.Println("****<{transData   uuu}>*****", u)
			handleTransData(u) //有请u求需要处理

		case u := <-entering:
			fmt.Println("****<{1u}>*****", u)
			users[u.Email] = u //有新用户进入
			//广播聊天室信息
			bcroominfo <- 1

		case u := <-leaving:
			fmt.Println("****<{2u}>*****", u)
			delete(users, u.Email) //有用户离开
			//广播聊天室信息
			bcroominfo <- 1

		case <-bcroominfo:
			var data = TransferData{
				Type: "roominfo",
				Data: RoomInfo{
					RoomName:    "简易测试聊天室",
					OnlineNum:   len(users),
					OnlineUsers: usersMap2Slice(users),
					ServerID:    ServerID,
				},
			}
			transferDatas <- data
		}
	}
}


func handleTransData(data TransferData) {
	//当为msg时，发给部分用户，否则发给所有人
	if data.Type == "msg" {
		msg := data.Data.(Message)
		to := strings.TrimSpace(msg.To)
		if to != "" {
			//发送给发送者
			if u, ok := users[msg.From]; ok {
				u.Send2Client <- data
			}
			//如果接收者不为空，则发送给指定接收者
			tos := strings.Split(to, ";")
			for _, t := range tos {
				//如果是自己的，则跳过
				if t == msg.From {
					continue
				}
				if u, ok := users[t]; ok {
					//修改to为单个的目标
					msg.To = t
					data.Data = msg
					u.Send2Client <- data
				}
			}
		} else {
			//发给所有人
			for key := range users {
				users[key].Send2Client <- data
			}
		}
	} else {
		for _, v := range users {
			//发给所有人
			v.Send2Client <- data
		}
	}
}


//Index 聊天室页面
func Index(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "index.html")
}

func main() {
	fmt.Println("服务启动...")
	//启用任务调度gorountine
	go taskSchedule()
	//聊天室页面
	http.HandleFunc("/", Index)
	//聊天消息WebSocket
	http.Handle("/chat", websocket.Handler(chatRoom))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("监听端口失败")
	}
}


// websocket请求处理
func chatRoom(ws *websocket.Conn) {
	defer ws.Close()
	//新建用户
	user := UserInfo{
		Send2Client: make(chan TransferData, 10),
	}

	go send2client(ws, user)

	for {

		var datastr string
		//收到消息
		err := websocket.Message.Receive(ws, &datastr)
		// 若换成echo,则是receive(ws, []bytes)
		if err != nil {
			//客户端断开连接
			if err == io.EOF {
				//用户离开
				leaving <- user
				fmt.Println("client left")
				break
			}
		}

		//解析数据包里的type
		dtype := getType(datastr)
		switch dtype {
		case "login": //用户登录
			var logReq LoginReq
			err = json.Unmarshal([]byte(datastr), &logReq)
			if err == nil {
				//设置用户信息，并将用户存入列表
				user.Email = logReq.Data.Email
				user.NickName = logReq.Data.NickName
				//有用户进入聊天室
				entering <- user
				//发送欢迎语
				sendMessage(user, ServerID, "欢迎进入聊天室")
				///发送登录成功
				sendLoginResult(user, 1, "登录成功")
			}

		case "msg": //收到用户发送消息
			var msgReq MesssageReq
			err = json.Unmarshal([]byte(datastr), &msgReq)
			if err == nil {
				msgReq.Data.Time = time.Now().Format("2006-01-02 15:04:05")
				//将消息存入全局消息队列
				var transData = TransferData{
					Type: "msg",
					Data: msgReq.Data,
				}
				transferDatas <- transData
			}
		}
	}
}


//{"type": "roominfo", "data": {"form":"lhhtzj12@126.com", "to":"python@126.com"}}
//将用户消息chan里数据发送到客户端
// {"type":"roominfo","data":{"roomname":"简易测试聊天室",
// 							  "onlinenum":1,
// 							  "onlineusers":[{"nickname":"python","email":"py@126.com"}],
// 							  "serverid":"lhtzbj12@126.com"}
// 								}
func send2client(ws *websocket.Conn, user UserInfo) {
	for tdata := range user.Send2Client {
		b, _ := json.Marshal(tdata)
		fmt.Println("指定发送给用户？", string(b))
		websocket.Message.Send(ws, string(b))
	}
}


func sendLoginResult(user UserInfo, result int, msg string) {
	//返回登录结果
	var retunData = TransferData{
		Type: "login",
		Data: LoginRespData{
			Result: result,
			Msg:    msg,
		},
	}
	// user.Send2Client 是通道,往这里发消息就可以在for 捕获,然后case <-- transdata 在转成transdata之前，就已经TranferData{}格式化
	user.Send2Client <- retunData
}


//给用户发送消息
func sendMessage(user UserInfo, Form, Cont string) {
	//返回消息
	dataout := TransferData{
		Type: "msg",
		Data: Message{
			From: Form,
			To:   user.Email,
			Time: time.Now().Format("2006-01-02 15:04:05"),
			Cont: Cont,
		},
	}
	fmt.Println("*****<{--给用户发消息--}>*****", user.Send2Client)
	user.Send2Client <- dataout
}


//将用户列表从map转成[]
func usersMap2Slice(users map[string]UserInfo) []UserInfo {
	length := len(users)
	keys := make([]string, 0, length)
	result := make([]UserInfo, 0, length)
	for key := range users {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		result = append(result, users[key])
	}
	return result
}