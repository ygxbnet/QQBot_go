package api

import (
	"QQBot_go/config"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strings"
)

func sendHTTP(Sndpoint string, body []string) []byte {

	var PostBody string

	for _, v := range body {
		PostBody += v + "&"
	}

	resp, err := http.Post(config.Http_url+Sndpoint,
		"application/x-www-form-urlencoded",
		strings.NewReader(PostBody))
	if err != nil {
		log.Infoln("请求[" + config.Http_url + Sndpoint + "]错误")

		log.Infof("body: %s\n", strings.Replace(fmt.Sprintf("%s", body), "\n", "\\n", -1))

		log.Infoln(err)
		return nil
	}

	defer resp.Body.Close()
	RespBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Infoln(err)
		return nil
	}

	log.Infoln("请求[" + config.Http_url + Sndpoint + "]成功")

	log.Infof("body: %s\n", strings.Replace(fmt.Sprintf("%s", body), "\n", "\\n", -1))
	return RespBody
}
