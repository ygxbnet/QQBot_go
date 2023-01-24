package service

import (
	"QQBot_go/internal/handler"
	"QQBot_go/service/handle_order"
)

func init() {
	handler.AddHandlerGroupMessageFunc(group)
}

// Services 信息打印
func Services() {
	Init()
}

func group(groupID string, userID string, message string) {
	handle_order.HandleGroupOrder(groupID, userID, message)
}
