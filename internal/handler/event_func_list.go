package handler

import (
	"container/list"
)

var handlerGuildMessageFunc = list.New()
var handlerPrivateMessageFunc = list.New()
var handlerGroupMessageFunc = list.New()

// AddHandlerGuildMessageFunc 添加处理Guild消息方法
func AddHandlerGuildMessageFunc(HandlerFunc func(guildID string, channelID string, userID string, message string)) {
	handlerGuildMessageFunc.PushBack(HandlerFunc)
}

// AddHandlerPrivateMessageFunc 添加处理Private消息方法
func AddHandlerPrivateMessageFunc(HandlerFunc func(message string)) {
	handlerPrivateMessageFunc.PushBack(HandlerFunc)
}

// AddHandlerGroupMessageFunc 添加处理Group消息方法
func AddHandlerGroupMessageFunc(HandlerFunc func(groupID string, userID string, message string, messageID string)) {
	handlerGroupMessageFunc.PushBack(HandlerFunc)
}

// GuildMessage Guild消息
func GuildMessage(guildID string, channelID string, userID string, message string) {
	for e := handlerGuildMessageFunc.Front(); e != nil; e = e.Next() {
		go e.Value.(func(guildID string, channelID string, userID string, message string))(guildID, channelID, userID, message)
	}
}

// PrivateMessage Private消息
func PrivateMessage(message string) {
	for e := handlerPrivateMessageFunc.Front(); e != nil; e = e.Next() {
		go e.Value.(func(message string))(message)
	}
}

// GroupMessage Group消息
func GroupMessage(groupID string, userID string, message string, messageID string) {
	for e := handlerGroupMessageFunc.Front(); e != nil; e = e.Next() {
		go e.Value.(func(groupID string, userID string, message string, messageID string))(groupID, userID, message, messageID)
	}
}
