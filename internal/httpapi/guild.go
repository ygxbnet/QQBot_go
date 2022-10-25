package httpapi

import (
	log "github.com/sirupsen/logrus"
)

func Send_guild_channel_msg(guild_id string, channel_id string, message string) {
	data := make(map[string]string)
	data["guild_id"] = guild_id
	data["channel_id"] = channel_id
	data["message"] = message

	body := sendHTTP("/send_guild_channel_msg", data)
	log.Infof("Guild消息发送结果: %s", string(body))
}
