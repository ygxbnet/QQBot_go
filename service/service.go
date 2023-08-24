package service

import (
	"QQBot_go/internal/handler"
	group2 "QQBot_go/service/group"
)

// Services 信息打印
func Services() {
	handler.AddHandlerGroupMessageFunc(group)
	group2.Init()
}

func group(groupID string, userID string, message string, messageID string) {
	group2.HandleGroupOrder(groupID, userID, message, messageID)
}
