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

var historyMessage = make(map[string][]Message)

func AskQuestion(groupID string, userID string, message string, messageID string) {

	helpMessage := "你好，我是 ChatGPT 的替身。" +
		"\n你可以与我对话，可以和我聊天，可以问我问题，我都会一直陪着你 ٩( 'ω' )و " +
		"\n" +
		"\n使用方法：/q [对话内容]" +
		"\n或者直接 @我 [对话内容]" +
		"\n例如：/q 请介绍一下你自己" +
		"\n" +
		"\n命令：/q restart 重新开启一个对话"

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

	} else if strings.Fields(message)[1] == "restart" {
		// 重新开启一个新的对话
		_, status := historyMessage[groupID]
		if status {
			delete(historyMessage, groupID)
		}
		httpapi.SendGroupMsg(groupID, "已经重新开启一个新的对话")

	} else {
		// 拼接请求，用于支持连续对话，使用 gpt-3.5-turbo 模型
		historyMessage[groupID] = append(historyMessage[groupID], Message{
			Role:    "user",
			Content: strings.Fields(message)[1],
		})
		jsonByte, _ := json.Marshal(Body{
			Model:    "gpt-3.5-turbo",
			Messages: historyMessage[groupID],
		})

		client := &http.Client{}
		req, _ := http.NewRequest("POST", apiURL, bytes.NewReader(jsonByte))
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer "+config.Get().OpenAI.APIKey[0])

		// 发送请求
		log.Info("正在发送请求：", apiURL, string(jsonByte))
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
			httpapi.SendGroupMsg(groupID, fmt.Sprintf("[CQ:reply,id=%s]获取失败，请重试或换一个问题\n%s", messageID, returnMessage))
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
