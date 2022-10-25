package group

import (
	"QQBot_go/internal/httpapi"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

var refreshStructs = map[string]*refresh{}

func RefreshHandle(group_id string, user_id string, message string) {
	if refreshStructs[user_id] != nil {
		refreshStructs[user_id].refresh(group_id, user_id, message)
	}
}

// 刷屏
func GroupRefresh(group_id string, user_id string, message string) {
	refreshNumber := 2

	if len(strings.Fields(message)) == 1 {
		//刷屏
		var msg1 = fmt.Sprintf(
			"[CQ:at,qq=%s]"+
				"\n✅将把您的下一条消息作为刷屏消息"+
				"\n/sp [刷屏次数](默认2次 最多为10次)", user_id)
		httpapi.Send_group_msg(group_id, msg1)

	} else if len(strings.Fields(message)) == 2 {
		//刷屏 指定刷屏次数
		num, err := strconv.Atoi(strings.Fields(message)[1])
		if err != nil {
			log.Error(err)
			httpapi.Send_group_msg(group_id, fmt.Sprintf("[CQ:at,qq=%s]"+"\n❌指定刷屏次数错误", user_id))
			return
		}
		if num <= 10 {
			refreshNumber = num
		} else {
			refreshNumber = 10
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
	doRefresh(group_id, user_id, refreshNumber)
}

func doRefresh(group_id string, user_id string, refreshNumber int) {
	//刷屏实现
	if _, ok := refreshStructs[user_id]; !ok {
		r := &refresh{}
		r.setNumber(refreshNumber)

		refreshStructs[user_id] = r
	} else {
		delete(refreshStructs, user_id)
	}
}

// 刷屏结构体
type refresh struct {
	number int
}

func (receiver *refresh) setNumber(number int) {
	receiver.number = number
}
func (receiver *refresh) refresh(group_id string, user_id string, message string) {
	for i := 1; i <= receiver.number; i++ {
		httpapi.Send_group_msg(group_id, message)
	}
	delete(refreshStructs, user_id)
}
