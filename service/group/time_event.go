package group

import (
	"QQBot_go/internal/base"
	"QQBot_go/internal/config"
	"QQBot_go/internal/cqapi"
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"strconv"
	"time"
)

// Init 初始化
func Init() {
	// 发送基本信息
	msg := fmt.Sprintf("[CQ:at,qq=%s]\nBot 已启动\n当前程序版本：%s\n构建时间：%s", config.Get().Account.AdminID, base.Version, base.BuildTime)
	cqapi.SendGroupMsg(config.Get().Group.InfoID, msg)

	// 每 12h 定时向 Test 群发送消息
	go func() {
		count := 0
		for {
			time.Sleep(time.Hour * 12)
			count = count + 1
			cqapi.SendGroupMsg(config.Get().Group.InfoID, "每 12h 定时发送消息\n程序发送次数："+strconv.Itoa(count))
		}
	}()

	// 0点定时发送相关消息
	go func() {
		for {
			<-time.NewTimer(getTimeDifference(0, 0, 0)).C // 定时器

			// 发送消息
			timeMessage := fmt.Sprintf("现在是：%d年%d月%d日 星期%s\n又是新的一天，希望大家过得开心", time.Now().Year(), time.Now().Month(), time.Now().Day(), conversionWeek(time.Now().Weekday().String()))
			cqapi.SendGroupMsg(config.Get().Group.MainID, timeMessage)

			// urlImg, nameImg := services_api.GetBingPictureURL()
			// httpapi.SendGroupMsg(config.Get().Group.MainID, "[CQ:image,file="+nameImg+",subType=0,url="+urlImg+"]")
		}
	}()
}

func Init_Backup() {

	// 6点定时发送相关消息
	go func() {
		for {
			<-time.NewTimer(getTimeDifference(6, 0, 0)).C

			{ // 每日一句
				response, _ := http.Get("https://v.api.aa1.cn/api/pyq/index.php?aa1=json")
				bytes, _ := io.ReadAll(response.Body)
				msg := gjson.Parse(string(bytes))
				if response.StatusCode == 200 {
					message := fmt.Sprintf("%s", msg.Get("pyq").String())
					cqapi.SendGroupMsg(config.Get().Group.MainID, message)
				} else {
					cqapi.SendGroupMsg(config.Get().Group.MainID, "早上好！！！")
				}
			}
			// { // 早上好
			// 	response, _ := http.Get("https://v.api.aa1.cn/api/zaoanyulu/index.php")
			// 	bytes, _ := io.ReadAll(response.Body)
			// 	msg := gjson.Get(string(bytes))
			// 	if response.StatusCode == 200 && msg.Get("code").Int() == 1 {
			// 		message := fmt.Sprintf("%s", msg.Get("text").String())
			// 		httpapi.SendGroupMsg(config.Get().Group.MainID, message)
			// 	} else {
			// 		httpapi.SendGroupMsg(config.Get().Group.MainID, "早上好！！！")
			// 	}
			// }
			// { // 每日笑话
			// 	var msg []string
			//
			// 	for i := 0; i < 3; i++ {
			// 		response, _ := http.Get("https://v.api.aa1.cn/api/api-wenan-gaoxiao/index.php?aa1=json")
			// 		if response.StatusCode == 200 {
			// 			bytes, _ := io.ReadAll(response.Body)
			// 			msg = append(msg, gjson.Get(string(bytes)).Get("0.gaoxiao").String())
			// 		}
			// 	}
			// 	httpapi.SendGroupMsg(config.Get().Group.MainID, fmt.Sprintf("每日笑话三则：\n1. %s\n2. %s\n3. %s", msg[0], msg[1], msg[2]))
			// }
		}
	}()
}

// getTimeDifference 获取时间差
func getTimeDifference(Hour int, Min int, Sec int) time.Duration {
	now := time.Now()
	var next time.Time

	setTime := time.Date(now.Year(), now.Month(), now.Day(), Hour, Min, Sec, 0, now.Location())

	if setTime.After(now) {
		next = setTime
	} else {
		next = setTime.Add(24 * time.Hour)
		next = time.Date(next.Year(), next.Month(), next.Day(), Hour, Min, Sec, 0, now.Location())
	}

	return next.Sub(now)
}

// conversionWeek 星期转换文字
func conversionWeek(Weekday string) string {
	switch Weekday {
	case "Sunday":
		return "天"
	case "Monday":
		return "一"
	case "Tuesday":
		return "二"
	case "Wednesday":
		return "三"
	case "Thursday":
		return "四"
	case "Friday":
		return "五"
	case "Saturday":
		return "六"
	}
	return "Error"
}
