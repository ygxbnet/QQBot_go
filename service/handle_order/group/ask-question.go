package group

import (
	"QQBot_go/internal/httpapi"
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"strings"
)

func AskQuestion(groupID string, userID string, message string, messageID string) {

	// apiurl := "https://v1.apigpt.cn/?apitype=sql&q=%s"
	apiurl := "https://v1.apigpt.cn/?q=%s"

	if len(strings.Fields(message)) == 1 {
		httpapi.SendGroupMsg(groupID,
			"你好，我是ChatGPT的替身，你可以问我各种问题，我都可以帮你回答（但不一定正确 ←_←）"+
				"\n注意：请不要问政治相关问题"+
				"\n\n使用方法：/q [问题内容]"+
				"\n例如：/q 你是谁")
	} else {
		httpapi.SendGroupMsg(groupID, "AI正在努力生成中，请稍后......")

		response, err := http.Get(fmt.Sprintf(apiurl, strings.Fields(message)[1]))
		if err != nil {
			httpapi.SendGroupMsg(groupID, "请求发生错误：\n"+err.Error())
			return
		}
		returnMessage, _ := io.ReadAll(response.Body)
		httpapi.SendGroupMsg(groupID, fmt.Sprintf("[CQ:reply,id=%s]%s", messageID, gjson.Parse(string(returnMessage)).Get("ChatGPT_Answer").String()))
	}
}
