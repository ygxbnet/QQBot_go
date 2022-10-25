package httpapi

import (
	log "github.com/sirupsen/logrus"
	"strconv"
)

func Send_group_msg(group_id string, message string) string {
	data := make(map[string]string)
	data["group_id"] = group_id
	data["message"] = message

	body := sendHTTP("/send_group_msg", data)
	log.Infof("Group消息发送结果: %s", string(body))
	return string(body)
}
func Set_group_ban(group_id string, user_id string, duration int) string {
	data := make(map[string]string)
	data["group_id"] = group_id
	data["user_id"] = user_id
	data["duration"] = strconv.Itoa(duration)

	body := sendHTTP("/set_group_ban", data)
	log.Infof("Group禁言结果: %s", string(body))
	return string(body)
}
