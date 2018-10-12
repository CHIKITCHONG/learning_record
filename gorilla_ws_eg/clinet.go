// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	//设置一个channel来发送信号,长度为1
	interrupt := make(chan os.Signal, 1)
	// 一直运行一直到收到一个信号
	signal.Notify(interrupt, os.Interrupt)

	// url.URL是一个结构体，需要初始化,套路写法
	// Scheme:方案 Host：本机,path：请求处理websocket消息的路径
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/echo"}

	log.Printf("connecting to %s", u.String())

	// 创建链接
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			// TextMessage 文本消息
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			// 发送关闭消息后 服务器会收到read: websocket: close 1000 (normal)
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			//case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}