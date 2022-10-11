package service

import (
	"QQBot_go/handler"
	"QQBot_go/service/handle_order"
	log "github.com/sirupsen/logrus"
)

func init() {
	handler.AddHandlerGroupMessageFunc(Group)
}

func Services() {
	log.Info("功能模块将以插件模式运行")
}

func Group(group_id string, user_id string, message string) {
	if message[0:1] == "/" {
		handle_order.HandleOrder_Group(group_id, user_id, message)
	} else {
	}
}
