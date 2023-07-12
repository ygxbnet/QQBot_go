package group

import (
	"QQBot_go/internal/config"
	"QQBot_go/internal/httpapi"
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"strings"
)

func AskQuestion(groupID string, userID string, message string, messageID string) {

	helpMessage := "你好，我是 ChatGPT 的替身。" +
		"\n你可以问我各种问题，我都可以帮你回答（但不一定正确 (←_←)）" +
		"\n" +
		"\n使用方法：/q [提问内容]" +
		"\n或者直接 @我 [提问内容]" +
		"\n" +
		"\n例如：/q 请介绍一下你自己" +
		"\n注：该功能还没有接入连续上下文对话"

	// 拼接请求 URL
	var apiURL string
	if config.Get().OpenAI.BaseURL == "" {
		apiURL = "https://api.openai.com" + "/v1/chat/completions"
	} else {
		apiURL = config.Get().OpenAI.BaseURL + "/v1/chat/completions"
	}

	if len(strings.Fields(message)) == 1 {
		// 获取插件帮助
		httpapi.SendGroupMsg(groupID, helpMessage)

	} else {
		// 拼接请求，使用 gpt-3.5-turbo 模型
		jsonByte, _ := json.Marshal(
			Body{
				Model: "gpt-3.5-turbo",
				Messages: []Message{
					{Role: "user", Content: strings.Fields(message)[1]},
				},
			})
		client := &http.Client{}
		req, _ := http.NewRequest("POST", apiURL, bytes.NewReader(jsonByte))
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer "+config.Get().OpenAI.APIKey)

		// 发送请求
		res, err := client.Do(req)
		if err != nil {
			log.Error(err)
			httpapi.SendGroupMsg(groupID, "请求发生错误：\n"+err.Error())
			return
		}
		defer res.Body.Close()

		// 处理请求并给用户回应
		returnMessage, _ := io.ReadAll(res.Body)
		if gjson.Parse(string(returnMessage)).Get("choices.0.message.content").String() == "" {
			httpapi.SendGroupMsg(groupID, fmt.Sprintf("[CQ:reply,id=%s]获取失败，请重试或换一个问题", messageID))
		} else {
			httpapi.SendGroupMsg(groupID, fmt.Sprintf("[CQ:reply,id=%s]ChatGPT：%s", messageID, gjson.Parse(string(returnMessage)).Get("choices.0.message.content").String()))
		}
	}
}

type Body struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
