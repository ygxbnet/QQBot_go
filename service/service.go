package service

import (
	"QQBot_go/internal/handler"
	"QQBot_go/service/handleorder"
)

// Services 信息打印
func Services() {
	handler.AddHandlerGroupMessageFunc(group)
	Init()
}

func group(groupID string, userID string, message string, messageID string) {
	handleorder.HandleGroupOrder(groupID, userID, message, messageID)
}
