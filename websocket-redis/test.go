package main

import (
    "fmt"
    "github.com/gorilla/websocket"

    //"log"
    "github.com/garyburd/redigo/redis"
    //"net/http"
)

var (
    gRedisConn  = func() (redis.Conn, error) {
        return redis.Dial("tcp", ":6379")
    }
)

const(
    testMessage = "Hello"
    testID = "1234568910"
)


func connectWS() *websocket.Conn {
   // Returning 404
   u := "ws://127.0.0.1:5342?/ws?id=1"

   // Debug
   conn, r, err := websocket.DefaultDialer.Dial(u, nil)
   if err != nil {
       fmt.Println(r)
       panic(err)
   }

   return conn
}

func main() {
    go func() {
        if c, err := gRedisConn(); err != nil {
            fmt.Println("Error establishing Redis connection")
        } else {
            c.Do("PUBLISH", testID, testMessage)
        }
    }()
    fmt.Println("-----programing start here------")
    ws := connectWS()
    if ws != nil {}
    //for {
    //  _, resp, err := ws.ReadMessage()
    //  if err != nil {
    //      fmt.Println("read:", err)
    //      return
    //  }
    //
    //  fmt.Println("Got value", resp)
    //}
}

