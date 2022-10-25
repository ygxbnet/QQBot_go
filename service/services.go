package service

import (
	"QQBot_go/internal/handler"
	"QQBot_go/service/handle_order"
	log "github.com/sirupsen/logrus"
)

func init() {
	handler.AddHandlerGroupMessageFunc(Group)
}

func Services() {
	log.Info("功能模块将以插件模式运行")
	Init()
}

func Group(group_id string, user_id string, message string) {
	handle_order.HandleOrder_Group(group_id, user_id, message)
}
