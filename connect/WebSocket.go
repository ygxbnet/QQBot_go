package connect

import (
	"QQBot_go/config"
	"QQBot_go/handler"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"time"
)

func Connect() {
	log.Info("正在连接:", config.WebSocket_url)
	c, _, err := websocket.DefaultDialer.Dial(config.WebSocket_url, nil)
	if err != nil {
		log.Error("连接错误", err)
		time.Sleep(time.Second)
		log.Info("正在重连......")
		time.Sleep(2 * time.Second)
		Connect()
	} else {
		log.Info("连接成功")
	}
	defer c.Close()
	//done := make(chan struct{})
	//接受消息
	go func() {
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Error("消息接受错误", err)
				return
			}
			handler.EventHandler(string(message)) //处理消息
		}
	}()

	for true {
		time.Sleep(time.Second)
	}
}
