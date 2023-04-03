package service

import (
	"QQBot_go/internal/handler"
	"QQBot_go/service/handle_order"
)

// Services 信息打印
func Services() {
	handler.AddHandlerGroupMessageFunc(group)
	Init()
}

func group(groupID string, userID string, message string, messageID string) {
	handle_order.HandleGroupOrder(groupID, userID, message, messageID)
}
