package handle_order

import (
	"QQBot_go/db"
	"QQBot_go/db/model"
	"QQBot_go/internal/httpapi"
	"fmt"
	"github.com/tidwall/gjson"
	"strconv"
	"strings"
	"time"
)

var help_info = "----------帮助信息----------" +
	"\n\n/help 获取帮助" +
	"\n/info 获取机器人信息" +
	"\n\n/dk 进行打卡" +
	"\n/sp 进行刷屏"

func HandleOrder_Group(group_id string, user_id string, message string) {
	switch strings.Fields(message)[0][0:] {
	case "/", "／": //指令为空时
		httpapi.Send_group_msg(group_id, "指令不能为空")

	case "/help", "／help":
		httpapi.Send_group_msg(group_id, help_info)

	case "/info", "／info": //机器人信息
		httpapi.Send_group_msg(group_id, info)

	case "/dk", "／dk", "打卡", "&#91;冒泡&#93;":
		Group_dk(group_id, user_id)

	case "/sp", "／sp", "刷屏":
		groupRefresh(group_id, user_id, message)

	case "/test", "／test":
		httpapi.Send_group_msg(group_id, "[CQ:share,url=https://gitee.com/YGXB-net/QQBot_go/blob/develop/CHANGELOG.md#更新日志]")

	default:
		if message[0:1] == "/" || message[0:3] == "／" {
			httpapi.Send_group_msg(group_id, "命令输入错误或没有此命令\n请输入 /help 查看帮助")
		}
	}
}

// 刷屏
func groupRefresh(group_id string, user_id string, message string) {

	bannedNumber := 5
	var msg1 = fmt.Sprintf(
		"[CQ:at,qq=%s]"+
			"\n将把您的下一条消息作为刷屏消息"+
			"\n/sp [刷屏次数](默认5次)", user_id)
	var msg2 = fmt.Sprintf(
		"[CQ:at,qq=%s]"+
			"\n将把您的下一条消息作为刷屏消息"+
			"\n刷屏次数: %d", user_id, bannedNumber)
	var msg_error = "参数错误或多余"

	if len(strings.Fields(message)) == 1 {

		httpapi.Send_group_msg(group_id, msg1)
	} else if len(strings.Fields(message)) == 2 {

		httpapi.Send_group_msg(group_id, msg2)
	} else {
		httpapi.Send_group_msg(group_id, msg_error)
		return
	}
}

// 打卡
func Group_dk(group_id string, user_id string) {
	UserData := db.ReadDBFile("group", user_id)
	message := ""
	var dk_data model.User_Data
	time_now, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	time_data, _ := time.Parse("2006-01-02", gjson.Parse(UserData).Get("dk_last_time").String())
	time_difference := (time_now.Unix() - time_data.Unix()) / 86400

	if UserData == "" { //没有打卡记录
		dk_data.DK_Last_Time = time_now.Format("2006-01-02")
		dk_data.DK_Times = 1
		message = fmt.Sprintf(message_dk, user_id, "✅打卡成功", "这是你的第一次打卡！")

	} else { //有打卡记录
		if time_difference == 0 { //当天打卡（打卡失败）
			dk_data.DK_Last_Time = time_now.Format("2006-01-02")
			dk_data.DK_Times = int(gjson.Parse(UserData).Get("dk_times").Int())
			message = fmt.Sprintf(message_dk, user_id, "❌打卡失败", "今天你已经打卡了！")

		} else if time_difference == 1 { //昨天打卡
			dk_data.DK_Last_Time = time_now.Format("2006-01-02")
			dk_data.DK_Times = int(gjson.Parse(UserData).Get("dk_times").Int()) + 1
			message = fmt.Sprintf(message_dk, user_id, "✅打卡成功",
				"你已经连续打卡了"+strconv.Itoa(dk_data.DK_Times)+"次了！"+
					"\n[CQ:face,id=144][CQ:face,id=144][CQ:face,id=144][CQ:face,id=144][CQ:face,id=144]")

		} else if time_difference > 1 { //间隔两天以上打卡
			dk_data.DK_Last_Time = time_now.Format("2006-01-02")
			now := time.Now()
			now_unix := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).Unix()

			dk_data.DK_Times = int(gjson.Parse(UserData).Get("dk_times").Int()) + 1 //int(gjson.Parse(UserData).Get("dk_times").Int())

			message = fmt.Sprintf(message_dk, user_id, "✅打卡成功",
				"你已经打卡了"+strconv.Itoa(dk_data.DK_Times)+"次了"+
					"\n上次打卡时间为: "+
					"\n"+time_data.Format("2006-01-02")+
					"（"+strconv.FormatInt((now_unix-time_data.Unix())/(60*60*24), 10)+"天前）")
		}
	}

	db.WriteDBFile("group", user_id, dk_data)
	httpapi.Send_group_msg(group_id, message)
}
