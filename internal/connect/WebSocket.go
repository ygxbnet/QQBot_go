package connect

import (
	"QQBot_go/internal/config"
	"QQBot_go/internal/handler"
	"time"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

// Connect 连接WebSocket
func Connect() {
	webSocketURL := config.Parse().Server.Websocket.URL

	log.Info("正在连接: ", webSocketURL)
	c, _, err := websocket.DefaultDialer.Dial(webSocketURL, nil)
	if err != nil {
		log.Error("连接错误: ", err)
		time.Sleep(time.Second)
		log.Info("正在重连......")
		time.Sleep(2 * time.Second)

		if c != nil {
			err := c.Close()
			if err != nil {
				log.Error(err)
			}
		}
		Connect()
	} else {
		log.Info("连接成功")
	}
	defer c.Close()
	// done := make(chan struct{})
	// 接受消息
	go func() {
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Error("消息接受错误: ", err)

				if c != nil {
					err := c.Close()
					if err != nil {
						log.Error(err)
					}
				}
				log.Info("正在重连")
				Connect()

				return
			}
			handler.EventHandler(string(message)) // 处理消息
		}
	}()

	for true {
		time.Sleep(time.Second)
	}
}
