package api

import (
	log "github.com/sirupsen/logrus"
	"strconv"
)

func Send_group_msg(group_id string, message string) string {
	body := sendHTTP("/send_group_msg", []string{"group_id=" + group_id, "message=" + message})

	log.Infof("Group消息发送结果: %s", string(body))
	return string(body)
}
func Set_group_ban(group_id string, user_id string, duration int) string {
	body := sendHTTP("/set_group_ban", []string{"group_id=" + group_id, "user_id=" + user_id, "duration=" + strconv.Itoa(duration)})

	log.Infof("Group禁言结果: %s", string(body))
	return string(body)
}
