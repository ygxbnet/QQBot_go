package handle_order

import (
	"QQBot_go/api"
	"QQBot_go/db"
	"QQBot_go/db/model"
	"fmt"
	"github.com/tidwall/gjson"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var help_info = "----------帮助信息----------" +
	"\n\n/help 获取帮助" +
	"\n/info 获取机器人信息" +
	"\n/dk 或 /打卡 进行打卡" +
	"\n/jy 或 /禁言 对指定的人禁言指定时长" +
	//"\n/p 或 /图片 获取Bing每日的壁纸" +
	"\n\n----------注意----------" +
	"\n\n\"/\"为英文输入法的\"/\" 而非中文输入法的\"／\""

func HandleOrder_Group(group_id string, user_id string, message string) {
	switch strings.Fields(message)[0][1:] {
	case "": //指令为空时
		api.Send_group_msg(group_id, "指令不能为空")

	case "info": //机器人信息
		api.Send_group_msg(group_id, info)

	case "help":
		api.Send_group_msg(group_id, help_info)

	case "dk", "打卡":
		Group_dk(group_id, user_id)

	case "jy", "禁言":
		Group_jy(group_id, user_id, message)

	//case "p", "图片":
	//	api.Send_group_msg(group_id, "此功能正在开发（头发都要没了！）")
	//url_img, name_img := services_api.Get_Random_Picture_URL()
	//fmt.Println(url_img, name_img)
	//api.Send_group_msg(group_id, "[CQ:image,file="+name_img+",subType=0,url="+url_img+"]\nBing每日壁纸")

	case "test":
		msg := fmt.Sprintf(message_dk, user_id, "成功", "你已经打卡了次了\n上次时间打卡为：2006-01-02")
		api.Send_group_msg(group_id, msg)

	default:
		api.Send_group_msg(group_id, "命令输入错误或没有此命令\n请输入 /help 查看帮助")
	}
}

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
		message = fmt.Sprintf(message_dk, user_id, "成功", "这是你的第一次打卡！")
	} else { //有打卡记录
		if time_difference == 0 { //当天打卡（打卡失败）
			dk_data.DK_Last_Time = time_now.Format("2006-01-02")
			dk_data.DK_Times = int(gjson.Parse(UserData).Get("dk_times").Int())
			message = fmt.Sprintf(message_dk, user_id, "失败", "今天你已经打卡了！")
		} else if time_difference == 1 { //昨天打卡
			dk_data.DK_Last_Time = time_now.Format("2006-01-02")
			dk_data.DK_Times = int(gjson.Parse(UserData).Get("dk_times").Int()) + 1
			message = fmt.Sprintf(message_dk, user_id, "成功", "你已经连续打卡了"+strconv.Itoa(dk_data.DK_Times)+"次了！\n假期打卡已全部完成！[CQ:face,id=144][CQ:face,id=144][CQ:face,id=144]")
		} else if time_difference > 1 { //间隔两天以上打卡
			dk_data.DK_Last_Time = time_now.Format("2006-01-02")
			dk_data.DK_Times = int(gjson.Parse(UserData).Get("dk_times").Int()) + 1 //int(gjson.Parse(UserData).Get("dk_times").Int())
			message = fmt.Sprintf(message_dk, user_id, "成功", "你已经打卡了"+strconv.Itoa(dk_data.DK_Times)+"次了\n上次时间打卡为："+time_data.Format("2006-01-02"))
		}
	}
	db.WriteDBFile("group", user_id, dk_data)
	api.Send_group_msg(group_id, message)
}

func Group_jy(group_id string, user_id string, message string) {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Group处理禁言时发生错误：")
			fmt.Println(err)
			api.Send_group_msg(group_id, "Group处理禁言时发生错误")

			str := fmt.Sprintf("%v", err)
			api.Send_group_msg("115987946", str)
		}
	}()

	var duration int
	if len(strings.Fields(message)) == 2 {
		duration = 10
	} else if len(strings.Fields(message)) == 3 {
		duration, _ = strconv.Atoi(strings.Fields(message)[2])
	} else {
		api.Send_group_msg(group_id,
			"缺少指令参数"+
				"\n\n/jy [@的人] [时间(秒)](可选,默认10秒)"+
				"\n\n例如："+
				"\n/jy @YGXB_net 60"+
				"\n/jy @YGXB_net")
		return
	}

	reg := regexp.MustCompile("\\d+")
	silence_user_id := reg.FindAllString(strings.Fields(message)[1], -1)[0]
	result := api.Set_group_ban(group_id, silence_user_id, duration)

	status := gjson.Parse(result).Get("status")
	if status.String() == "ok" {
		api.Send_group_msg(group_id, "执行成功")
	} else {
		api.Send_group_msg(group_id, "执行失败："+gjson.Parse(result).Get("wording").String())
	}
}
