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
	// 连接WebSocket
	c, _, err := websocket.DefaultDialer.Dial(webSocketURL, nil)
	if err != nil {
		log.Error("连接错误: ", err)
		time.Sleep(time.Second * 4)
		log.Info("正在重连......")
		time.Sleep(time.Second * 4)

		if c != nil {
			c.Close()
		}
		Connect()
	} else {
		log.Info(webSocketURL + " 连接成功")
		// 接受消息
		go func() {
			for {
				_, message, err := c.ReadMessage()
				if err != nil {
					log.Error("消息接受错误: ", err)

					if c != nil {
						c.Close()
					}
					log.Info("正在重连")
					break
				} else {
					handler.EventHandler(string(message)) // 处理消息
				}
			}
			Connect()
		}()
	}
}
