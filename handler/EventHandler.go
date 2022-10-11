package handler

import (
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

func EventHandler(message string) {
	if gjson.Parse(message).Get("post_type").String() == "message" {

		var messageType = ""
		switch gjson.Parse(message).Get("message_type").String() {
		case "guild":

			if gjson.Parse(message).Get("user_id").String() != "144115218684493716" {
				messageType = "Guild"

				var guild_id = gjson.Parse(message).Get("guild_id").String()
				var channel_id = gjson.Parse(message).Get("channel_id").String()
				var user_id = gjson.Parse(message).Get("user_id").String()
				var msg = gjson.Parse(message).Get("message").String()
				GuildMessage(guild_id, channel_id, user_id, msg)
			}
		case "group":
			messageType = "Group"

			var group_id = gjson.Parse(message).Get("group_id").String()
			var user_id = gjson.Parse(message).Get("user_id").String()
			var msg = gjson.Parse(message).Get("message").String()

			GroupMessage(group_id, user_id, msg)
		case "private":
			messageType = "Private"

			PrivateMessage(message)
		}
		log.Info("接受到["+messageType+"]消息: ", string(message))
	} else {
		return
	}
}
