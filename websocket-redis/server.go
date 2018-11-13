package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var serverAddress = ":5342"

// required : 线程安全的websocket 对象,安全关闭
// 刷新两次也会卡
var (
	cache     *Cache
	pubSub    *redis.PubSubConn  	// redis 订阅服务
	redisConn = func() (redis.Conn, error) {
		return redis.Dial("tcp", ":6379")
	}
)

func init() {
	cache = &Cache{
		Users: make([]*User, 0, 1),
	}
}

type User struct {
	ID   string
	conn *websocket.Conn
}


/*type cache struct{
	user[
		ID:		string
		conn: 	*websocket.Conn
		]
	sync.Mutex
}
*/
type Cache struct {
	Users []*User
	sync.Mutex
}

type Message struct {
	DeliveryID string `json:"id"`
	Content    string `json:"content"`
}


// 新建一个user对象,开启订阅
func (c *Cache) newUser(conn *websocket.Conn, id string) *User {
	u := &User{
		ID:   id,
		conn: conn,
	}
	// 订阅(收取消息)
	if err := pubSub.Subscribe(u.ID); err != nil {
		panic(err)
	}
	c.Lock()
	defer c.Unlock()
	c.Users = append(c.Users, u)
	return u
}


func main() {
	// 开启redis服务
	redisConn, err := redisConn()
	if err != nil {
		panic(err)
	}
	// 最后关闭redis服务
	defer redisConn.Close()


	// 创建redis订阅实例？
	pubSub = &redis.PubSubConn{Conn: redisConn}
	// 最后关掉
	defer pubSub.Close()

	// 开启一个发送信息的服务
	go deliverMessages()

	http.HandleFunc("/ws", wsHandler)
	http.HandleFunc("/ws/push", wsPusher)
	log.Printf("server started at %s\n", serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, nil))
}


func wsHandler(w http.ResponseWriter, r *http.Request) {
	// 升级协议头
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("upgrader error %s\n" + err.Error())
		return
	}

	// 创建一个根据user_id的独立通道
	u := cache.newUser(conn, r.FormValue("id"))
	log.Printf("user %s joined\n", u.ID)
	//heart_id, err := strconv.Atoi(r.FormValue(r.FormValue("id")))
	//go func() {
	//	var (
	//		err error
	//	)
	//	for {
	//		if err = u.conn.WriteMessage(heart_id, []byte("heartbeat")); err != nil {
	//			u.conn.Close()
	//		}
	//		time.Sleep(30 * time.Second)
	//	}
	//}()
	for {
		var m Message
		if err := u.conn.ReadJSON(&m); err != nil {
			u.conn.Close()
			log.Printf("error on ws. message %s\n", err)
			break
		}

		if c, err := redisConn(); err != nil {
			log.Printf("error on redis conn. %s\n", err)
		} else {
			// 这是发布消息(发送信息)
			c.Do("PUBLISH", m.DeliveryID, string(m.Content))
		}
	}
	defer conn.Close()
	// 心跳机制,启动线程,不断发消息
	//go func() {
	//	var (
	//		err error
	//	)
	//	for {
	//		resultID, errors := strconv.Atoi(r.FormValue("id"))
	//		if errors != nil {
	//			return
	//		}
	//		if err = u.conn.WriteMessage(resultID, []byte(" ")); err != nil {
	//			u.conn.Close()
	//			break
	//		}
	//		time.Sleep(120 * time.Second)
	//	}
	//}()
}

//func (conn *Cache)Close(){
//	// 线程安全，可多次调用
//	conn.Close()
//	// 利用标记，让closeChan只关闭一次
//	conn.Lock()
//	if !conn.isClosed {
//		close(conn.closeChan)
//		conn.isClosed = true
//	}
//	// 释放线程锁
//	conn.mutex.Unlock()
//
//}
// 写消息api
func wsPusher(w http.ResponseWriter, r *http.Request) {
	if c, err := redisConn(); err != nil {
		log.Println("redis 链接错误")
	} else {
		// 这是发布消息(发送信息)
		fmt.Println("id= ", r.FormValue("id"))
		//fmt.Println("content= ", r.URL.Query().Get("content"))
		fmt.Println("content= ", r.FormValue("content"))

		c.Do("PUBLISH", r.FormValue("id"), r.FormValue("content"))

	}
}

func deliverMessages() {
	for {
		switch v := pubSub.Receive().(type) {
		case redis.Message:
			cache.findAndDeliver(v.Channel, string(v.Data))
		case redis.Subscription:
			log.Printf("subscription message: %s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			log.Println("error pub/sub on connection, delivery has stopped")
			return
		}
	}
}

func (c *Cache) findAndDeliver(userID string, content string) {
	m := Message{
		Content: content,
	}
	for _, u := range c.Users {
		if u.ID == userID {
			if err := u.conn.WriteJSON(m); err != nil {
				u.conn.Close()
				log.Printf("error on message delivery through ws. e: %s\n", err)
			} else {
				log.Printf("user %s found at our store, message sent\n", userID)
			}
			return
		}
	}
	log.Printf("user %s not found at our store\n", userID)
}
