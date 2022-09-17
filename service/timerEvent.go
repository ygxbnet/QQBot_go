package service

import (
	"QQBot_go/api"
	"QQBot_go/service/services_api"
	"fmt"
	"strconv"
	"time"
)

func init() {
	var conut = 0

	//每10min定时向Test群发送消息
	go func() {
		for true {
			conut = conut + 1
			api.Send_group_msg("115987946", "每10min定时发送\n次数："+strconv.Itoa(conut))
			//time.Sleep(time.Second * 10)
			time.Sleep(time.Minute * 10)
		}
	}()

	//每天8点定时发送Bing的每日壁纸
	go func() {
		for {
			fmt.Println("每日获取Bing壁纸开启")
			//开启计时
			t := time.NewTimer(get_time_difference(6, 0, 0))
			<-t.C

			url_img, name_img := services_api.Get_Bing_Picture_URL()
			api.Send_group_msg("1038122549", "[CQ:image,file="+name_img+",subType=0,url="+url_img+"]")
			//api.Send_group_msg("1038122549", "每日Bing壁纸")
		}
	}()

	//每天晚上9点提醒未打卡的打卡
	//go func() {
	//	for {
	//		t := time.NewTimer(get_time_difference(20, 0, 0))
	//		<-t.C
	//
	//		now := time.Now()
	//		August31 := time.Date(2022, time.August, 31, 0, 0, 0, 0, time.Local)
	//		if now.Sub(August31) < 0 {
	//			api.Send_group_msg("1038122549", "[CQ:at,qq=all] 快点来打卡呀！")
	//		} else if now.Sub(August31) < time.Hour*24 {
	//			api.Send_group_msg("1038122549", "[CQ:at,qq=all]\n现在是假期的最后一天了，也是假期最后一天打卡了，我破例一回，就当作每一天都打卡了。")
	//			api.Send_group_msg("1038122549", "快点来打卡吧！")
	//		}
	//	}
	//}()
}

// 获取时间差
func get_time_difference(Hour int, Min int, Sec int) time.Duration {
	now := time.Now()
	var next time.Time

	set_time := time.Date(now.Year(), now.Month(), now.Day(), Hour, Min, Sec, 0, now.Location())

	if set_time.After(now) {
		next = set_time
	} else {
		next = set_time.Add(24 * time.Hour)
		next = time.Date(next.Year(), next.Month(), next.Day(), Hour, Min, Sec, 0, now.Location())
	}

	return next.Sub(now)
}
