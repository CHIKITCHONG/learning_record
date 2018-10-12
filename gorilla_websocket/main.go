//package main
//
//import (
//	"fmt"
//	"net/http"
//
//	"github.com/gorilla/websocket"
//)
//
//var upgrader = websocket.Upgrader{
//	ReadBufferSize:  1024,
//	WriteBufferSize: 1024,
//}
//
//func main() {
//	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
//		conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
//
//		for {
//			// Read message from browser
//			// 从客户端读取消息
//			msgType, msg, err := conn.ReadMessage()
//			if err != nil {
//				return
//			}
//
//			// Print the message to the console
//			// 把消息打印到终端
//			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))
//
//			// Write message back to browser
//			// 把客户端的消息发回去
//			if err = conn.WriteMessage(msgType, msg); err != nil {
//				return
//			}
//		}
//	})
//
//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		http.ServeFile(w, r, "websockets.html")
//	})
//
//	http.ListenAndServe(":8080", nil)
//}

package main

/*
error：
 */
import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/middleware"
)

//
var upgrader = websocket.Upgrader{
	EnableCompression: true,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}


func hello(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	// 创建完后需要关闭ws
	defer ws.Close()

	for {
		// Write []byte类型
		err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
		if err != nil {
			log.Fatal(err)
		}

		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", msg)
	}
}

//func sending_message(message string) {
//
//	handle_message := []byte(message)
//	err := ws.WriteMessage(websocket.TextMessage, handle_message)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("%s\n", handle_message)
//}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//e.Static("/", "websockets.html")
	e.File("/" , "websockets.html")
	e.GET("/echo", hello)
	e.Logger.Fatal(e.Start(":8000"))
}