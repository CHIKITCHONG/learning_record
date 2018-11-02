package main

import (
	"fmt"
	"net/http"
		"github.com/gorilla/websocket"
		"./impl"
		"time"
		)
var(
	upgrader = websocket.Upgrader{
		// 允许跨域
		CheckOrigin:func(r *http.Request) bool{
			return true
		},
	}
)
func wsHandler(w http.ResponseWriter , r *http.Request){
	//	w.Write([]byte("hello"))
	var(
		wsConn *websocket.Conn
		err error
		conn *impl.Connection
		data []byte
	)
	// 完成ws协议的握手操作
	// Upgrade:websocket
	//wsConn, err = websocket.Upgrade(w, r, http.Header{"Sec-Websocket-Protocol": nil}, 1024, 1024);
	if err != nil {
		return
	}
	if wsConn , err = upgrader.Upgrade(w,r,nil); err != nil{
		return
	}
	request2 := r.Host
	fmt.Println("----request here----", request2)
	if conn , err = impl.InitConnection(wsConn); err != nil{
		goto ERR
	}


	// 启动线程，不断发消息
	go func(){
		var (err error)
		for{
			if err = conn.WriteMessage([]byte("heartbeat"));err != nil{
				return 
			}
			time.Sleep(2*time.Second)
		}
	}()

	for {
		if data , err = conn.ReadMessage();err != nil{
			goto ERR
		}
		if err = conn.WriteMessage(data);err !=nil{
			goto ERR
		}
	}

	ERR:
		conn.Close()

}





func main(){
	http.HandleFunc("/ws",wsHandler)
	http.ListenAndServe("0.0.0.0:5000",nil)
}