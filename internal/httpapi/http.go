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

	jsonData, err := json.Marshal(body)
	if err != nil {
		log.Error(err)
		return nil
	}
	postBody := strings.NewReader(string(jsonData))

	req, err := http.NewRequest("POST", config.Http_url+path, postBody)
	if err != nil {
		log.Error(err)
		return nil
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error("请求[" + config.Http_url + path + "]错误")
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

	log.Info("请求[" + config.Http_url + path + "]成功")

	log.Infof("body: %s\n", strings.Replace(fmt.Sprintf("%s", body), "\n", "\\n", -1))
	return RespBody
}
