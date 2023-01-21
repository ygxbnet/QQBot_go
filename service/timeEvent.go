package service

import (
	"QQBot_go/internal/base"
	"QQBot_go/internal/config"
	"QQBot_go/internal/httpapi"
	"QQBot_go/service/services_api"
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
	httpapi.SendGroupMsg(config.Parse().Group.InfoID, fmt.Sprintf("Bot已启动\n当前版本: %s\n构建时间: %s", base.VERSION, base.BUILD_TIME))

	// 每10min定时向 Test 群发送消息
	go func() {
		var count = 0
		for true {
			count = count + 1
			httpapi.SendGroupMsg(config.Parse().Group.InfoID, "每10min定时发送\n次数: "+strconv.Itoa(count))
			time.Sleep(time.Minute * 10)
		}
	}()

	// 0点定时发送相关消息
	go func() {
		for {
			t := time.NewTimer(getTimeDifference(0, 0, 0))
			<-t.C

			// 发送消息
			timeMessage := fmt.Sprintf("今天是: %d年%d月%d日 星期%s", time.Now().Year(), time.Now().Month(), time.Now().Day(), conversionWeek(time.Now().Weekday().String()))
			httpapi.SendGroupMsg(config.Parse().Group.MainID, timeMessage)

			urlImg, nameImg := services_api.GetBingPictureURL()
			httpapi.SendGroupMsg(config.Parse().Group.MainID, "[CQ:image,file="+nameImg+",subType=0,url="+urlImg+"]")
		}
	}()

	// 6点定时发送问好
	go func() {
		for {
			t := time.NewTimer(getTimeDifference(6, 0, 0))
			<-t.C

			{ // 早上好
				response, _ := http.Get("https://v.api.aa1.cn/api/yiyan/index.php?type=json")
				bytes, _ := io.ReadAll(response.Body)
				msg := gjson.Parse(string(bytes))
				if response.StatusCode == 200 {
					message := fmt.Sprintf("%s", msg.Get("yiyan").String())
					httpapi.SendGroupMsg(config.Parse().Group.MainID, message)
				} else {
					httpapi.SendGroupMsg(config.Parse().Group.MainID, "早上好！！！")
				}
			}
			// { // 早上好
			// 	response, _ := http.Get("https://v.api.aa1.cn/api/zaoanyulu/index.php")
			// 	bytes, _ := io.ReadAll(response.Body)
			// 	msg := gjson.Parse(string(bytes))
			// 	if response.StatusCode == 200 && msg.Get("code").Int() == 1 {
			// 		message := fmt.Sprintf("%s", msg.Get("text").String())
			// 		httpapi.SendGroupMsg(config.Parse().Group.MainID, message)
			// 	} else {
			// 		httpapi.SendGroupMsg(config.Parse().Group.MainID, "早上好！！！")
			// 	}
			// }
			{ // 每日笑话
				var msg []string

				for i := 0; i < 3; i++ {
					response, _ := http.Get("https://v.api.aa1.cn/api/api-wenan-gaoxiao/index.php?aa1=json")
					if response.StatusCode == 200 {
						bytes, _ := io.ReadAll(response.Body)
						msg = append(msg, gjson.Parse(string(bytes)).Get("0.gaoxiao").String())
					}
				}
				httpapi.SendGroupMsg(config.Parse().Group.MainID, fmt.Sprintf("每日笑话三则：\n1. %s\n2. %s\n3. %s", msg[0], msg[1], msg[2]))
			}
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
