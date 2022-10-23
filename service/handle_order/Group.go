package handle_order

import (
	"QQBot_go/internal/base"
	"QQBot_go/internal/httpapi"
	"QQBot_go/service/handle_order/group"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strings"
)

var help_info = "----------帮助信息----------" +
	"\n\n/help 获取帮助" +
	"\n/info 获取机器人信息" +
	"\n\n/dk 进行打卡"

var info = "本机器人由YGXB_net开发" +
	"\nQQ:3040809965" +
	"\n\n当前版本: " + base.Version +
	"\n更新日志: https://gitee.com/YGXB-net/QQBot_go/blob/master/CHANGELOG.md"

func HandleOrder_Group(group_id string, user_id string, message string) {
	switch strings.Fields(message)[0] {
	case "/", "／":
		//指令为空时
		httpapi.Send_group_msg(group_id, "指令不能为空")
	case "/help", "／help":
		//帮助指令
		httpapi.Send_group_msg(group_id, help_info)
	case "/info", "／info":
		//机器人信息
		httpapi.Send_group_msg(group_id, info)
	case "/dk", "／dk", "打卡", "&#91;冒泡&#93;":
		//打卡指令
		group.Group_dk(group_id, user_id)
	case "/sp", "／sp", "刷屏":
		//刷屏指令
		group.GroupRefresh(group_id, user_id, message)
	case "/test", "／test":
		httpapi.Send_group_msg(group_id, "This is test")
	default:
		handleEmojisOrder(group_id, user_id, message)
		//因为切片会出现长度不足，所以会抛出异常
		defer func() { recover() }()
		if message[0:1] == "/" || message[0:3] == "／" {
			httpapi.Send_group_msg(group_id, "命令输入错误或没有此命令\n请输入 /help 查看帮助")
		} else if strings.Index(message, "[CQ:at,qq=2700154874]") != -1 {
			httpapi.Send_group_msg(group_id, fmt.Sprintf("[CQ:at,qq=%s] 叫你爸爸干嘛？", user_id))
		}
	}
}

func handleEmojisOrder(group_id string, user_id string, message string) {
	//判断是否为图片消息
	if strings.Index(message, "CQ:image") == -1 {
		return
	}
	//提取图片消息的URL
	//原始数据举例: [CQ:image,file=d3ab70d3f8b6b4eb2c7878d5177dc051.image,subType=1,url=https://gchat.qpic.cn/gchatpic_new/3040809965/2058987946-2446050292-D3AB70D3F8B6B4EB2C7878D5177DC051/0?term=3]
	indexURL := strings.Index(message, "url")
	URL := message[indexURL+4:][:len(message[indexURL+4:])-1]
	//请求URL，获取数据
	resp, err := http.Get(URL)
	if err != nil {
		log.Error("请求URL错误 ", err)
		return
	}
	data, _ := io.ReadAll(resp.Body)
	//对请求到的数据求MD5
	md5 := md5.New()
	md5.Write(data)
	MD5Str := hex.EncodeToString(md5.Sum(nil))

	switch MD5Str {
	case "d3ab70d3f8b6b4eb2c7878d5177dc051":
		//此MD5值对应的文件为:
		//https://gchat.qpic.cn/gchatpic_new/3040809965/2058987946-2282106232-D3AB70D3F8B6B4EB2C7878D5177DC051/0?term=3
		group.Group_dk(group_id, user_id)
	}
}
