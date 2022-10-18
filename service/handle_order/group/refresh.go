package group

import (
	"QQBot_go/internal/httpapi"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

var refreshStructs = map[string]*refresh{}

// 刷屏
func GroupRefresh(group_id string, user_id string, message string) {
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
	doRefresh(group_id, user_id, refreshNumber)
}

func doRefresh(group_id string, user_id string, refreshNumber int) {
	//刷屏实现
	if _, ok := refreshStructs[user_id]; !ok {
		r := &refresh{}
		r.setUserId(user_id)
		r.setNumber(refreshNumber)
		r.do()

		refreshStructs[user_id] = r
	} else {
		refreshStructs[user_id].do()
	}
}

type refresh struct {
	userID string
	number int

	times int
	isdo  bool
}

func (receiver *refresh) setUserId(userID string) {
	receiver.userID = userID
}

func (receiver *refresh) setNumber(number int) {
	receiver.number = number
}

func (receiver *refresh) do() {
	receiver.isdo = true
	go func() {
		for receiver.isdo {
			time.Sleep(time.Second)
			receiver.times += 1
			if receiver.times >= 5 {
				receiver.stop()
			}

		}
	}()
}

func (receiver *refresh) stop() {
	receiver.isdo = false
	receiver.times = 0
}
