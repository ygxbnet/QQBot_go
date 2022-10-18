package handle_order

import (
	"QQBot_go/internal/base"
	"QQBot_go/internal/httpapi"
	"QQBot_go/service/handle_order/group"
	"strings"
)

var help_info = "----------帮助信息----------" +
	"\n\n/help 获取帮助" +
	"\n/info 获取机器人信息" +
	"\n\n/dk 进行打卡" +
	"\n/sp 进行刷屏"

var info = "本机器人由YGXB_net开发" +
	"\nQQ:3040809965" +
	"\n\n当前版本: " + base.Version +
	"\n更新日志: https://gitee.com/YGXB-net/QQBot_go/blob/master/CHANGELOG.md"

func HandleOrder_Group(group_id string, user_id string, message string) {
	switch strings.Fields(message)[0][0:] {
	case "/", "／":
		//指令为空时
		httpapi.Send_group_msg(group_id, "指令不能为空")

	case "/help", "／help":
		httpapi.Send_group_msg(group_id, help_info)

	case "/info", "／info":
		//机器人信息
		httpapi.Send_group_msg(group_id, info)

	case "/dk", "／dk", "打卡", "&#91;冒泡&#93;":
		group.Group_dk(group_id, user_id)

	case "/sp", "／sp", "刷屏":
		group.GroupRefresh(group_id, user_id, message)

	case "/test", "／test":
		httpapi.Send_group_msg(group_id, "[CQ:share,url=https://gitee.com/YGXB-net/QQBot_go/blob/develop/CHANGELOG.md#更新日志]")

	default:
		if message[0:1] == "/" || message[0:3] == "／" {
			httpapi.Send_group_msg(group_id, "命令输入错误或没有此命令\n请输入 /help 查看帮助")
		}
	}
}
