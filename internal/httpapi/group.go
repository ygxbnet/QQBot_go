package httpapi

import (
	log "github.com/sirupsen/logrus"
	"strconv"
)

// SendGroupMsg 发送Group消息
func SendGroupMsg(groupID string, message string) string {
	data := make(map[string]string)
	data["group_id"] = groupID
	data["message"] = message

	body := sendHTTP("/send_group_msg", data)
	log.Infof("Group消息发送结果: %s", string(body))
	return string(body)
}

// SetGroupBan Group禁言
func SetGroupBan(groupID string, userID string, duration int) string {
	data := make(map[string]string)
	data["group_id"] = groupID
	data["user_id"] = userID
	data["duration"] = strconv.Itoa(duration)

	body := sendHTTP("/set_group_ban", data)
	log.Infof("Group禁言结果: %s", string(body))
	return string(body)
}
