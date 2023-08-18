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
var apiKey []string
var apiURL string

func init() {
	// 拼接请求 URL
	if config.Get().OpenAI.BaseURL == "" {
		apiURL = "https://api.openai.com" + "/v1/chat/completions"
	} else {
		apiURL = config.Get().OpenAI.BaseURL + "/v1/chat/completions"
	}
	// 验证配置文件中所有的 OpenAI Key
	log.Info("正在验证所有 OpenAI Key")
	verifyOpenAIKey()
}

func AskQuestion(groupID string, userID string, message string, messageID string) {

	helpMessage := "你好，我是 ChatGPT 的替身。" +
		"\n你可以与我对话，可以和我聊天，可以问我问题，我都会一直陪着你 ٩( 'ω' )و " +
		"\n" +
		"\n使用方法：/q [对话内容]" +
		"\n或者直接 @我 [对话内容]" +
		"\n例如：/q 请介绍一下你自己" +
		"\n" +
		"\n命令：" +
		"\n/q restart 重新开启一个对话" +
		"\n/q check 检查 OpenAI Key"

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

	} else if strings.Fields(message)[1] == "check" {
		// 检查 OpenAI Key是否可用
		var key []string
		var rMessage string

		for _, value := range config.Get().OpenAI.APIKey {
			client := &http.Client{}
			req, _ := http.NewRequest("GET", config.Get().OpenAI.BaseURL+"/v1/models", nil)
			req.Header.Add("Authorization", "Bearer "+value)
			res, err := client.Do(req)
			if err != nil {
				log.Error(err)
				return
			}

			if res.StatusCode == 200 {
				rMessage += fmt.Sprintf("✅ 有效 %s......\n", value[:15])
				key = append(key, value)
			} else {
				rMessage += fmt.Sprintf("❌ 无效 %s......\n", value[:15])
			}
		}
		apiKey = key
		httpapi.SendGroupMsg(groupID, rMessage[:len(rMessage)-2])

	} else {
		// 拼接请求，用于支持连续对话，使用 gpt-3.5-turbo 模型
		historyMessage[groupID] = append(historyMessage[groupID], Message{
			Role:    "user",
			Content: fmt.Sprintf("%s", strings.Fields(message)[1:]),
		})
		jsonByte, _ := json.Marshal(Body{
			Model:    "gpt-3.5-turbo",
			Messages: historyMessage[groupID],
		})
		log.Info(string(jsonByte))
		getResponseMessage(groupID, messageID, jsonByte)
	}
}

// getResponseMessage 发送请求，获取 OpenAI 回复
func getResponseMessage(groupID string, messageID string, jsonByte []byte) {
	if len(apiKey) == 0 {
		verifyOpenAIKey()
		if len(apiKey) == 0 {
			httpapi.SendGroupMsg(groupID, "当前配置文件中已没有可用 OpenAI Key，请重新添加")
			return
		}
	}

	client := &http.Client{}
	req, _ := http.NewRequest("POST", apiURL, bytes.NewReader(jsonByte))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+apiKey[0])

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
	if gjson.Parse(string(returnMessage)).Get("choices.0.message.content").String() != "" {
		// 成功获取到回答
		httpapi.SendGroupMsg(groupID, fmt.Sprintf("[CQ:reply,id=%s]ChatGPT：\n%s",
			messageID,
			gjson.Parse(string(returnMessage)).Get("choices.0.message.content").String(),
		))

	} else if res.StatusCode == 401 {
		// OpenAI Key 无效
		log.Info("请求失败：", returnMessage)
		if len(apiKey) <= 0 {
			httpapi.SendGroupMsg(groupID, "当前已没有可用 OpenAI Key。正在从配置文件中重新检索")
			verifyOpenAIKey()
			getResponseMessage(groupID, messageID, jsonByte)
		} else {
			httpapi.SendGroupMsg(groupID, "当前 OpenAI Key 不可用，正在切换 Key 并重新获取回复，请稍后...")
			apiKey = apiKey[1:]
			getResponseMessage(groupID, messageID, jsonByte)
		}
	} else {
		httpapi.SendGroupMsg(groupID, fmt.Sprintf("[CQ:reply,id=%s]获取失败，请重试或换一个问题\n%s",
			messageID,
			returnMessage,
		))
	}
}

// verifyOpenAIKey 验证配置文件中的 OpenAI Key
func verifyOpenAIKey() {
	var key []string
	for index, value := range config.Get().OpenAI.APIKey {
		client := &http.Client{}
		req, _ := http.NewRequest("GET", config.Get().OpenAI.BaseURL+"/v1/models", nil)
		req.Header.Add("Authorization", "Bearer "+value)
		res, err := client.Do(req)
		if err != nil {
			log.Error(err)
			return
		}

		if res.StatusCode == 200 {
			log.Infof("✅ %d Key %s 有效", index, value)
			key = append(key, value)
		} else {
			log.Infof("❌ %d Key %s 无效", index, value)
		}
	}
	apiKey = key
}

// Body 数据结构
type Body struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
