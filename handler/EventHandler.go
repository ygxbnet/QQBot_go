package handler

import (
	"fmt"
	"github.com/tidwall/gjson"
)

func EventHandler(message string) {
	if gjson.Parse(message).Get("post_type").String() == "message" {
		fmt.Println("接受到消息:", string(message))

		switch gjson.Parse(message).Get("message_type").String() {
		case "guild":

			if gjson.Parse(message).Get("user_id").String() != "144115218684493716" {
				fmt.Println("频道消息:", message)

				var guild_id = gjson.Parse(message).Get("guild_id").String()
				var channel_id = gjson.Parse(message).Get("channel_id").String()
				var user_id = gjson.Parse(message).Get("user_id").String()
				var msg = gjson.Parse(message).Get("message").String()
				GuildMessage(guild_id, channel_id, user_id, msg)
			}
		case "group":

			fmt.Println("群组消息:", message)

			var group_id = gjson.Parse(message).Get("group_id").String()
			var user_id = gjson.Parse(message).Get("user_id").String()
			var msg = gjson.Parse(message).Get("message").String()

			GroupMessage(group_id, user_id, msg)
		case "private":

			fmt.Println("私聊消息:", message)
			PrivateMessage(message)
		}
	} else {
		return
	}
}
