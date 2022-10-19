package handler

import (
	"container/list"
)

var HandlerGuildMessageFuncs = list.New()
var HandlerPrivateMessageFuncs = list.New()
var HandlerGroupMessageFuncs = list.New()

func AddHandlerGuildMessageFunc(HandlerFunc func(guild_id string, channel_id string, user_id string, message string)) {
	HandlerGuildMessageFuncs.PushBack(HandlerFunc)
}
func AddHandlerPrivateMessageFunc(HandlerFunc func(message string)) {
	HandlerPrivateMessageFuncs.PushBack(HandlerFunc)
}
func AddHandlerGroupMessageFunc(HandlerFunc func(group_id string, user_id string, message string)) {
	HandlerGroupMessageFuncs.PushBack(HandlerFunc)
}

func GuildMessage(guild_id string, channel_id string, user_id string, message string) {
	for e := HandlerGuildMessageFuncs.Front(); e != nil; e = e.Next() {
		go e.Value.(func(guild_id string, channel_id string, user_id string, message string))(guild_id, channel_id, user_id, message)
	}
}
func PrivateMessage(message string) {
	for e := HandlerPrivateMessageFuncs.Front(); e != nil; e = e.Next() {
		go e.Value.(func(message string))(message)
	}
}
func GroupMessage(group_id string, user_id string, message string) {
	for e := HandlerGroupMessageFuncs.Front(); e != nil; e = e.Next() {
		go e.Value.(func(group_id string, user_id string, message string))(group_id, user_id, message)
	}
}
