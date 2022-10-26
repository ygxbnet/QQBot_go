package handler

import (
	"container/list"
)

var handlerGuildMessageFuncs = list.New()
var handlerPrivateMessageFuncs = list.New()
var handlerGroupMessageFuncs = list.New()

// AddHandlerGuildMessageFunc 添加处理Guild消息方法
func AddHandlerGuildMessageFunc(HandlerFunc func(guildID string, channelID string, userID string, message string)) {
	handlerGuildMessageFuncs.PushBack(HandlerFunc)
}

// AddHandlerPrivateMessageFunc 添加处理Private消息方法
func AddHandlerPrivateMessageFunc(HandlerFunc func(message string)) {
	handlerPrivateMessageFuncs.PushBack(HandlerFunc)
}

// AddHandlerGroupMessageFunc 添加处理Group消息方法
func AddHandlerGroupMessageFunc(HandlerFunc func(groupID string, userID string, message string)) {
	handlerGroupMessageFuncs.PushBack(HandlerFunc)
}

// GuildMessage Guild消息
func GuildMessage(guildID string, channelID string, userID string, message string) {
	for e := handlerGuildMessageFuncs.Front(); e != nil; e = e.Next() {
		go e.Value.(func(guildID string, channelID string, userID string, message string))(guildID, channelID, userID, message)
	}
}

// PrivateMessage Private消息
func PrivateMessage(message string) {
	for e := handlerPrivateMessageFuncs.Front(); e != nil; e = e.Next() {
		go e.Value.(func(message string))(message)
	}
}

// GroupMessage Group消息
func GroupMessage(groupID string, userID string, message string) {
	for e := handlerGroupMessageFuncs.Front(); e != nil; e = e.Next() {
		go e.Value.(func(groupID string, userID string, message string))(groupID, userID, message)
	}
}
