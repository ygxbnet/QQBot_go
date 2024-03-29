package group

import (
	"QQBot_go/internal/base"
	"QQBot_go/internal/config"
	"QQBot_go/internal/cqapi"
	"QQBot_go/service/group/chatgpt"
	"QQBot_go/service/group/dk"
	"QQBot_go/service/group/get_random_picture"
	"QQBot_go/service/group/refresh"
	"crypto/md5"
	"encoding/hex"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strings"
	"time"
)

var HELP_MESSAGE = " ====>  帮助信息  <==== " +
	"\n" +
	"\n● /help 获取帮助" +
	"\n● /info 获取机器人信息" +
	"\n" +
	"\n ====>  功能指令  <==== " +
	"\n" +
	"\n● /dk  进行打卡" +
	"\n● /q   开始一个新的对话" +
	"\n● /p   获取随机风景图" +
	"\n● /sp  进行刷屏"

var INFO_MESSAGE = "本机器人由 YGXB_net 开发" +
	"\nQQ: " + config.Get().Account.AdminID +
	"\n\n当前版本：" + base.Version +
	"\n构建时间：" + base.BuildTime

// HandleGroupOrder 处理Group命令
func HandleGroupOrder(groupID string, userID string, message string, messageID string) {
	switch strings.Fields(message)[0] {
	case "/", "／":
		// 指令为空时
		cqapi.SendGroupMsg(groupID, "指令不能为空")

	case "/help", "／help":
		// 帮助指令
		cqapi.SendGroupMsg(groupID, HELP_MESSAGE)

	case "/info", "／info":
		// 机器人信息
		cqapi.SendGroupMsg(groupID, INFO_MESSAGE)

	case "/dk", "／dk", "打卡", "&#91;冒泡&#93;", "&#91;打卡&#93;":
		// 打卡指令
		dk.Dk(groupID, userID, messageID)

	case "/sp", "／sp", "刷屏":
		// 刷屏指令
		refresh.Refresh(groupID, userID, message)

	case "/date", "／date", "时间":
		// 发送服务器当前时间
		cqapi.SendGroupMsg(groupID, time.Now().Format("2006-01-02 15:04:05"))

	case "/p", "／p", "图片":
		get_random_picture.GetRandomPicture(groupID, userID, message)

	case "/q", "/question", "问个问题", "问一个问题":
		// 问问题
		go chatgpt.AskQuestion(groupID, userID, message, messageID)

	case "/test", "／test":
		// 测试指令
		cqapi.SendGroupMsg(groupID, "This is test")

	default:
		refresh.RefreshHandle(groupID, userID, message)
		handleEmojiOrder(groupID, userID, message, messageID)

		// 因为切片会出现长度不足，所以会抛出异常
		defer func() { recover() }()
		if message[0:1] == "/" || message[0:3] == "／" {
			cqapi.SendGroupMsg(groupID, "❌命令输入错误或没有此命令\n请输入 /help 查看帮助")
			return
		}
		if strings.Index(message, "[CQ:at,qq="+config.Get().Account.BotID+"]") != -1 {
			// @机器人 交给问问题处理方法处理
			chatgpt.AskQuestion(groupID, userID, message, messageID)
			return
		}
	}
}

func handleEmojiOrder(groupID string, userID string, message string, messageID string) {
	// 判断是否为图片消息
	if strings.Index(message, "CQ:image") == -1 {
		return
	}
	// 提取图片消息的URL
	// 原始数据举例: [CQ:image,file=d3ab70d3f8b6b4eb2c7878d5177dc051.image,subType=1,url=https://gchat.qpic.cn/gchatpic_new/3040809965/2058987946-2446050292-D3AB70D3F8B6B4EB2C7878D5177DC051/0?term=3]
	indexURL := strings.Index(message, "url")
	URL := message[indexURL+4:][:len(message[indexURL+4:])-1]
	// 请求URL，获取数据
	resp, err := http.Get(URL)
	if err != nil {
		log.Error("请求URL错误 ", err)
		return
	}
	data, _ := io.ReadAll(resp.Body)
	// 对请求到的数据求MD5
	FileMd5 := md5.New()
	FileMd5.Write(data)
	MD5Str := hex.EncodeToString(FileMd5.Sum(nil))

	switch MD5Str {
	case "d3ab70d3f8b6b4eb2c7878d5177dc051", // https://gchat.qpic.cn/gchatpic_new/3040809965/2058987946-2282106232-D3AB70D3F8B6B4EB2C7878D5177DC051/0?term=3
		"0833ab984df318f53c07466160859ca6", // https://gchat.qpic.cn/gchatpic_new/3040809965/818848626-2357317952-0833AB984DF318F53C07466160859CA6/0?term=3
		"a3caf31ff742d543a0645ad6710e077c", // https://gchat.qpic.cn/gchatpic_new/3040809965/818848626-3205803506-A3CAF31FF742D543A0645AD6710E077C/0?term=3
		"00fb5731dcaff37dd940ddaabcd20f10": // https://gchat.qpic.cn/gchatpic_new/3040809965/818848626-2682086032-00FB5731DCAFF37DD940DDAABCD20F10/0?term=3

		dk.Dk(groupID, userID, messageID)
	}
}
