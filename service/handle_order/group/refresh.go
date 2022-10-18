package group

import (
	"QQBot_go/internal/httpapi"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

// 刷屏
func groupRefresh(group_id string, user_id string, message string) {

	refreshNumber := 5

	if len(strings.Fields(message)) == 1 {
		//刷屏
		var msg1 = fmt.Sprintf(
			"[CQ:at,qq=%s]"+
				"\n✅将把您的下一条消息作为刷屏消息"+
				"\n/sp [刷屏次数](默认5次)", user_id)
		httpapi.Send_group_msg(group_id, msg1)

	} else if len(strings.Fields(message)) == 2 {
		//刷屏 指定刷屏次数
		num, err := strconv.Atoi(strings.Fields(message)[1])
		if err != nil {
			log.Error(err)
			httpapi.Send_group_msg(group_id, fmt.Sprintf("[CQ:at,qq=%s]"+"\n❌指定刷屏次数错误", user_id))
			return
		}

		if num <= 20 {
			refreshNumber = num
		} else {
			refreshNumber = 20
		}
		var msg2 = fmt.Sprintf(
			"[CQ:at,qq=%s]"+
				"\n✅将把您的下一条消息作为刷屏消息"+
				"\n刷屏次数: %d", user_id, refreshNumber)
		httpapi.Send_group_msg(group_id, msg2)

	} else {
		//参数错误
		httpapi.Send_group_msg(group_id, "❌参数错误或多余")
		return
	}

}
