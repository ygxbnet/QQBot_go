package refresh

import (
	"QQBot_go/internal/cqapi"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"strconv"
	"strings"
	"time"
)

var refreshStructs = map[refreshKey]*refresh{}

// RefreshHandle 刷屏处理
func RefreshHandle(groupID string, userID string, message string) {
	if refreshStructs[refreshKey{userID, groupID}] != nil {
		refreshStructs[refreshKey{userID, groupID}].Refresh(groupID, userID, message)
	}
}

// Refresh 刷屏
func Refresh(groupID string, userID string, message string) {
	refreshNumber := 2
	var messageID string

	if len(strings.Fields(message)) == 1 {
		// 刷屏
		var msg1 = fmt.Sprintf(
			"[CQ:at,qq=%s]"+
				"\n✅将把您的下一条消息作为刷屏消息"+
				"\n刷屏次数: 2次"+
				"\n/sp [刷屏次数](默认2次 最多为10次)", userID)
		messageID = gjson.Parse(cqapi.SendGroupMsg(groupID, msg1)).Get("data").Get("message_id").String()

	} else if len(strings.Fields(message)) == 2 {
		// 刷屏 指定刷屏次数
		num, err := strconv.Atoi(strings.Fields(message)[1])
		if err != nil {
			log.Error(err)
			cqapi.SendGroupMsg(groupID, fmt.Sprintf("[CQ:at,qq=%s]"+"\n❌指定刷屏次数错误", userID))
			return
		}
		if num <= 10 {
			refreshNumber = num
		} else if num == 110 {
			refreshNumber = 20
		} else {
			refreshNumber = 10
		}
		var msg2 = fmt.Sprintf(
			"[CQ:at,qq=%s]"+
				"\n✅将把您的下一条消息作为刷屏消息"+
				"\n刷屏次数: %d次", userID, refreshNumber)
		messageID = gjson.Parse(cqapi.SendGroupMsg(groupID, msg2)).Get("data").Get("message_id").String()

	} else if len(strings.Fields(message)) == 3 {
		// 刷屏 指定刷屏次数和刷屏内容
		num, err := strconv.Atoi(strings.Fields(message)[1])
		if err != nil {
			log.Error(err)
			cqapi.SendGroupMsg(groupID, fmt.Sprintf("[CQ:at,qq=%s]"+"\n❌指定刷屏次数错误", userID))
			return
		}
		if num <= 10 {
			refreshNumber = num
		} else if num == 110 {
			refreshNumber = 50
		} else {
			refreshNumber = 10
		}
		var msg2 = fmt.Sprintf(
			"[CQ:at,qq=%s]"+
				"\n✅将把您的下一条消息作为刷屏消息"+
				"\n刷屏次数: %d次", userID, refreshNumber)
		messageID = gjson.Parse(cqapi.SendGroupMsg(groupID, msg2)).Get("data").Get("message_id").String()

		if messageID != "" {
			time.AfterFunc(time.Minute, func() {
				cqapi.DeleteMsg(messageID)
			})
		}

		doRefreshWithMessage(groupID, userID, refreshNumber, strings.Fields(message)[2])
		return

	} else {
		// 参数错误
		cqapi.SendGroupMsg(groupID, "❌参数错误或多余")
		return
	}

	if messageID != "" {
		time.AfterFunc(time.Minute, func() {
			cqapi.DeleteMsg(messageID)
		})
	}
	doRefresh(groupID, userID, refreshNumber)
}

func doRefresh(groupID string, userID string, refreshNumber int) {
	// 刷屏实现
	if refreshStructs[refreshKey{userID, groupID}] == nil {
		r := &refresh{}
		r.SetNumber(refreshNumber)
		r.SetUserID(userID)
		r.SetGroupID(groupID)
		r.DelayDelete()

		refreshStructs[refreshKey{userID, groupID}] = r
	} else {
		refreshStructs[refreshKey{userID, groupID}].SetNumber(refreshNumber)
		refreshStructs[refreshKey{userID, groupID}].ResetTime()
	}
}

func doRefreshWithMessage(groupID string, userID string, refreshNumber int, message string) {
	// 刷屏实现
	if refreshStructs[refreshKey{userID, groupID}] == nil {
		r := &refresh{}
		r.SetNumber(refreshNumber)
		r.SetUserID(userID)
		r.SetGroupID(groupID)
		r.DelayDelete()

		refreshStructs[refreshKey{userID, groupID}] = r
	} else {
		refreshStructs[refreshKey{userID, groupID}].SetNumber(refreshNumber)
		refreshStructs[refreshKey{userID, groupID}].ResetTime()
	}
	refreshStructs[refreshKey{userID, groupID}].Refresh(groupID, userID, message)
}

// 索引架构体
type refreshKey struct {
	UserID  string
	GroupID string
}

// 刷屏结构体
type refresh struct {
	userID  string
	groupID string
	number  int
	time    int
}

func (receiver *refresh) SetNumber(number int) {
	receiver.number = number
}
func (receiver *refresh) SetUserID(userID string) {
	receiver.userID = userID
}
func (receiver *refresh) SetGroupID(groupID string) {
	receiver.groupID = groupID
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
		delete(refreshStructs, refreshKey{receiver.userID, receiver.groupID})
	}()
}
func (receiver *refresh) Refresh(groupID string, userID string, message string) {
	for i := 1; i <= receiver.number; i++ {
		cqapi.SendGroupMsg(groupID, message)
	}
	delete(refreshStructs, refreshKey{userID, groupID})
}
