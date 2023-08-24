package cqapi

import (
	log "github.com/sirupsen/logrus"
)

// SendGuildChannelMsg 发送频道消息
func SendGuildChannelMsg(guildID string, channelID string, message string) {
	data := make(map[string]string)
	data["guild_id"] = guildID
	data["channel_id"] = channelID
	data["message"] = message

	body := sendHTTP("/send_guild_channel_msg", data)
	log.Infof("Guild消息发送结果: %s", string(body))
}
