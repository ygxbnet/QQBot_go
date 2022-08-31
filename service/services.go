package service

import (
	"QQBot_go/handler"
	"QQBot_go/service/handle_order"
	"fmt"
)

func init() {
	handler.AddHandlerGroupMessageFunc(Group)
	handler.AddHandlerGuildMessageFunc(Guild)
}

func Services() {
	fmt.Println("功能模块将以插件模式运行")
}

func Guild(guild_id string, channel_id string, user_id string, message string) {
	if channel_id == "2513644" && guild_id == "40005401641382238" {
		if message[0:1] == "/" {
			handle_order.HandleOrder_Guild(guild_id, channel_id, user_id, message)
		} else {
		}
	}
}

func Group(group_id string, user_id string, message string) {
	if message[0:1] == "/" {
		handle_order.HandleOrder_Group(group_id, user_id, message)
	} else {
	}
}
