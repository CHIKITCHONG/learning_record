package main
//package main
//
//import (
//	"github.com/gorilla/websocket"
//	"github.com/labstack/echo"
//	"net/http"
//	"working_project/tourismcreater/app/middlewares"
//)
//
//
//// 允许跨域
//var (
//	upgrader = websocket.Upgrader {
//		CheckOrigin: func(r *http.Request) bool {
//			return true
//		},
//	}
//)
//
//func webHandler(c echo.Context) error {
//	return nil
//}
//
//func main() {
//	e := echo.New()
//	e.Use(middlewares.Logger())
//	e.GET("ws", webHandler)
//	e.Logger.Fatal(e.Start(":8000"))
//}