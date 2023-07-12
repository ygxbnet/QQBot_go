package httpapi

import (
	"QQBot_go/internal/config"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strings"
)

func sendHTTP(path string, body map[string]string) []byte {
	httpURL := config.Get().Server.HTTPAPI.URL

	jsonData, err := json.Marshal(body)
	if err != nil {
		log.Error(err)
		return nil
	}
	postBody := strings.NewReader(string(jsonData))

	req, err := http.NewRequest("POST", httpURL+path, postBody)
	if err != nil {
		log.Error(err)
		return nil
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error("请求[" + httpURL + path + "]错误")
		log.Infof("body: %s\n", strings.Replace(fmt.Sprintf("%s", body), "\n", "\\n", -1))
		log.Error(err)
		return nil
	}
	defer res.Body.Close()

	RespBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error(err)
		return nil
	}

	log.Info("请求[" + httpURL + path + "]成功")
	log.Infof("body: %s\n", strings.Replace(fmt.Sprintf("%s", body), "\n", "\\n", -1))
	http.DefaultClient.CloseIdleConnections() // 关闭空闲连接，防止内存泄漏

	return RespBody
}
