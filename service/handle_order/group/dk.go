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

var messageDk = "[CQ:at,qq=%s]" +
	"\n%s" +
	"\n%s"

// Dk 打卡
func Dk(groupID string, userID string) {
	UserData := db.ReadDBFile("group", userID)
	message := ""
	var dkData model.UserData
	timeNow, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	timeData, _ := time.Parse("2006-01-02", gjson.Parse(UserData).Get("dk_last_time").String())
	timeDifference := (timeNow.Unix() - timeData.Unix()) / 86400

	if UserData == "" { //没有打卡记录
		dkData.DkLastTime = timeNow.Format("2006-01-02")
		dkData.DkTimes = 1
		message = fmt.Sprintf(messageDk, userID, "✅打卡成功", "这是你的第一次打卡！")

	} else { //有打卡记录
		if timeDifference == 0 { //当天打卡（打卡失败）
			dkData.DkLastTime = timeNow.Format("2006-01-02")
			dkData.DkTimes = int(gjson.Parse(UserData).Get("dk_times").Int())
			message = fmt.Sprintf(messageDk, userID, "❌打卡失败", "今天你已经打卡了！")

		} else if timeDifference == 1 { //昨天打卡
			dkData.DkLastTime = timeNow.Format("2006-01-02")
			dkData.DkTimes = int(gjson.Parse(UserData).Get("dk_times").Int()) + 1
			message = fmt.Sprintf(messageDk, userID, "✅打卡成功",
				"你已经连续打卡了"+strconv.Itoa(dkData.DkTimes)+"次了！"+
					"\n[CQ:face,id=144][CQ:face,id=144][CQ:face,id=144][CQ:face,id=144][CQ:face,id=144]")

		} else if timeDifference > 1 { //间隔两天以上打卡
			dkData.DkLastTime = timeNow.Format("2006-01-02")
			now := time.Now()
			nowUnix := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).Unix()

			dkData.DkTimes = int(gjson.Parse(UserData).Get("dk_times").Int()) + 1 //int(gjson.Parse(UserData).Get("dk_times").Int())

			message = fmt.Sprintf(messageDk, userID, "✅打卡成功",
				"你已经打卡了"+strconv.Itoa(dkData.DkTimes)+"次了"+
					"\n上次打卡时间为: "+
					"\n"+timeData.Format("2006-01-02")+
					"（"+strconv.FormatInt((nowUnix-timeData.Unix())/(60*60*24), 10)+"天前）")
		}
	}

	db.WriteDBFile("group", userID, dkData)
	httpapi.SendGroupMsg(groupID, message)
}
