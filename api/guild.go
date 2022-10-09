package api

import (
	"fmt"
)

func Send_guild_channel_msg(guild_id string, channel_id string, message string) {
	body := sendHTTP("/send_guild_channel_msg", []string{"guild_id=" + guild_id, "channel_id=" + channel_id, "message=" + message})

	fmt.Printf("Guild消息发送结果: %s", string(body))
}
