package service

import (
	"QQBot_go/internal/handler"
	"QQBot_go/service/handle_order"
	log "github.com/sirupsen/logrus"
)

func init() {
	handler.AddHandlerGroupMessageFunc(group)
}

// Services 信息打印
func Services() {
	log.Info("功能模块将以插件模式运行")
	Init()
}

func group(groupID string, userID string, message string) {
	handle_order.HandleGroupOrder(groupID, userID, message)
}
