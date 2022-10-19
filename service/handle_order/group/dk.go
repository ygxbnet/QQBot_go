package group

import (
	"QQBot_go/internal/db"
	"QQBot_go/internal/db/model"
	"QQBot_go/internal/httpapi"
	"fmt"
	"github.com/tidwall/gjson"
	"strconv"
	"time"
)

var message_dk = "[CQ:at,qq=%s]" +
	"\n%s" +
	"\n%s"

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
