package connect

import (
	"QQBot_go/config"
	"QQBot_go/handler"
	"fmt"
	"github.com/gorilla/websocket"
	"time"
)

func Connect() {
	fmt.Println("正在连接:", config.WebSocket_url)
	c, _, err := websocket.DefaultDialer.Dial(config.WebSocket_url, nil)
	if err != nil {
		fmt.Println("连接错误", err)
		time.Sleep(time.Second)
		fmt.Println("正在重连......")
		time.Sleep(2 * time.Second)
		Connect()
	} else {
		fmt.Println("连接成功")
	}
	defer c.Close()
	//done := make(chan struct{})
	//接受消息
	go func() {
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				fmt.Println("消息接受错误", err)
				return
			}
			handler.EventHandler(string(message)) //处理消息
		}
	}()

	for true {
		time.Sleep(time.Second)
	}
}
