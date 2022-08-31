package handle_order

import (
	"QQBot_go/api"
	"QQBot_go/db"
	"QQBot_go/db/model"
	"fmt"
	"github.com/tidwall/gjson"
	"strconv"
	"strings"
	"time"
)

func HandleOrder_Guild(guild_id string, channel_id string, user_id string, message string) {
	switch strings.Fields(message)[0][1:] {
	case "":
		api.Send_guild_channel_msg(guild_id, channel_id, "指令不能为空")
	case "help":
		var help_info = "[	帮助信息	]\n/test 让Bot发送Hello\n/sp 让Bot发送5条Hello信息进行刷屏"
		api.Send_guild_channel_msg(guild_id, channel_id, help_info)

	case "test":
		api.Send_guild_channel_msg(guild_id, channel_id, "Hello")
	case "sp":
		for i := 1; i <= 5; i++ {
			api.Send_guild_channel_msg(guild_id, channel_id, "刷屏"+strconv.Itoa(i))
		}
	case "谁是sb", "sb", "whoissb":
		api.Send_guild_channel_msg(guild_id, channel_id, "[CQ:at,qq="+user_id+"]是")
	case "t":
		api.Send_guild_channel_msg(guild_id, channel_id, "/t")
	case "dk", "打卡":
		Guild_dk(guild_id, channel_id, user_id)
		api.Send_guild_channel_msg(guild_id, channel_id, "此功能正在编写中......(YGXB_net正在掉头发中......[CQ:face,id=1])")
	default:
		api.Send_guild_channel_msg(guild_id, channel_id, "命令输入错误或没有此命令\n请输入/help查看帮助")
	}
}

func Guild_dk(guild_id string, channel_id string, user_id string) {
	UserData := db.ReadDBFile("guild", user_id)
	message := ""
	var dk_data model.User_Data
	time_now, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	time_data, _ := time.Parse("2006-01-02", gjson.Parse(UserData).Get("dk_last_time").String())
	time_difference := (time_now.Unix() - time_data.Unix()) / 86400

	if UserData == "" { //没有打卡记录
		dk_data.DK_Last_Time = time_now.Format("2006-01-02")
		dk_data.DK_Times = 1
		message = fmt.Sprintf(message_dk, user_id, "成功", "这是你的第一次打卡！")
	} else { //有打卡记录
		if time_difference == 0 { //当天打卡（打卡失败）
			dk_data.DK_Last_Time = time_now.Format("2006-01-02")
			dk_data.DK_Times = 1
			message = fmt.Sprintf(message_dk, user_id, "失败", "今天你已经打卡了！")
		} else if time_difference == 1 { //昨天打卡
			dk_data.DK_Last_Time = time_now.Format("2006-01-02")
			dk_data.DK_Times = int(gjson.Parse(UserData).Get("dk_times").Int()) + 1
			message = fmt.Sprintf(message_dk, user_id, "成功", "你已经连续打卡了"+gjson.Parse(UserData).Get("dk_times").String()+"次了！")
		} else if time_difference > 1 { //间隔两天以上打卡
			dk_data.DK_Last_Time = time_now.Format("2006-01-02")
			dk_data.DK_Times = int(gjson.Parse(UserData).Get("dk_times").Int())
			message = fmt.Sprintf(message_dk, user_id, "成功", "上次时间打卡为："+time_data.Format("2006-01-02"))
		}
	}
	db.WriteDBFile("guild", user_id, dk_data)
	api.Send_guild_channel_msg(guild_id, channel_id, message)
}
