package service

import (
	"QQBot_go/internal/config"
	"QQBot_go/internal/httpapi"
	"QQBot_go/service/services_api"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

// Init 初始化
func Init() {
	var count = 0

	// 每10min定时向Test群发送消息
	go func() {
		for true {
			count = count + 1
			httpapi.SendGroupMsg(config.Parse().Group.InfoID, "每10min定时发送\n次数: "+strconv.Itoa(count))
			// time.Sleep(time.Second * 10)
			time.Sleep(time.Minute * 10)
		}
	}()

	// 每天8点定时发送Bing的每日壁纸
	go func() {
		for {
			log.Info("每日获取Bing壁纸开启")
			// 开启计时
			t := time.NewTimer(getTimeDifference(6, 0, 0))
			<-t.C

			urlImg, nameImg := services_api.GetBingPictureURL()
			httpapi.SendGroupMsg(config.Parse().Group.MainID, "[CQ:image,file="+nameImg+",subType=0,url="+urlImg+"]")

			timeMessage := fmt.Sprintf("今天是: %d年%d月%d日 星期%s", time.Now().Year(), time.Now().Month(), time.Now().Day(), conversionWeek(time.Now().Weekday().String()))
			httpapi.SendGroupMsg(config.Parse().Group.MainID, timeMessage)
		}
	}()

	// 每天晚上9点提醒未打卡的打卡
	// go func() {
	//	for {
	//		t := time.NewTimer(getTimeDifference(20, 0, 0))
	//		<-t.C
	//
	//		now := time.Now()
	//		August31 := time.Date(2022, time.August, 31, 0, 0, 0, 0, time.Local)
	//		if now.Sub(August31) < 0 {
	//			httpapi.SendGroupMsg("1038122549", "[CQ:at,qq=all] 快点来打卡呀！")
	//		} else if now.Sub(August31) < time.Hour*24 {
	//			httpapi.SendGroupMsg("1038122549", "[CQ:at,qq=all]\n现在是假期的最后一天了，也是假期最后一天打卡了，我破例一回，就当作每一天都打卡了。")
	//			httpapi.SendGroupMsg("1038122549", "快点来打卡吧！")
	//		}
	//	}
	// }()
}

// 获取时间差
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
