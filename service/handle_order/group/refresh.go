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

// RefreshHandle 刷屏处理
func RefreshHandle(groupID string, userID string, message string) {
	if refreshStructs[userID] != nil {
		refreshStructs[userID].Refresh(groupID, userID, message)
	}
}

// Refresh 刷屏
func Refresh(groupID string, userID string, message string) {
	refreshNumber := 2

	if len(strings.Fields(message)) == 1 {
		//刷屏
		var msg1 = fmt.Sprintf(
			"[CQ:at,qq=%s]"+
				"\n✅将把您的下一条消息作为刷屏消息"+
				"\n刷屏次数: 2次"+
				"\n/sp [刷屏次数](默认2次 最多为10次)", userID)
		httpapi.SendGroupMsg(groupID, msg1)

	} else if len(strings.Fields(message)) == 2 {
		//刷屏 指定刷屏次数
		num, err := strconv.Atoi(strings.Fields(message)[1])
		if err != nil {
			log.Error(err)
			httpapi.SendGroupMsg(groupID, fmt.Sprintf("[CQ:at,qq=%s]"+"\n❌指定刷屏次数错误", userID))
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
				"\n刷屏次数: %d次", userID, refreshNumber)
		httpapi.SendGroupMsg(groupID, msg2)
	} else {
		//参数错误
		httpapi.SendGroupMsg(groupID, "❌参数错误或多余")
		return
	}
	doRefresh(groupID, userID, refreshNumber)
}

func doRefresh(groupID string, userID string, refreshNumber int) {
	//刷屏实现
	if refreshStructs[userID] == nil {
		r := &refresh{}
		r.SetNumber(refreshNumber)
		r.SetUserID(userID)
		r.DelayDelete()

		refreshStructs[userID] = r
	} else {
		refreshStructs[userID].SetNumber(refreshNumber)
		refreshStructs[userID].ResetTime()
	}
}

// 刷屏结构体
type refresh struct {
	userID string
	number int
	time   int
}

func (receiver *refresh) SetNumber(number int) {
	receiver.number = number
}
func (receiver *refresh) SetUserID(userID string) {
	receiver.userID = userID
}
func (receiver *refresh) ResetTime() {
	receiver.time = 300
}
func (receiver *refresh) DelayDelete() {
	receiver.time = 300
	go func() {
		for receiver.time > 0 {
			time.Sleep(time.Second)
			receiver.time--
		}
		delete(refreshStructs, receiver.userID)
	}()
}
func (receiver *refresh) Refresh(groupID string, userID string, message string) {
	for i := 1; i <= receiver.number; i++ {
		httpapi.SendGroupMsg(groupID, message)
	}
	delete(refreshStructs, userID)
}
