package services_api

import (
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"strings"
)

// GetBingPictureURL 获取Bing图片地址
func GetBingPictureURL() (urlImg string, nameImg string) {

	resp, err := http.Post("https://cn.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1&mkt=zh-CN", "", nil)
	if err != nil {
		log.Error("请求[https://cn.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1&mkt=zh-CN]错误")
		log.Error(err)
		return "", ""
	}
	defer resp.Body.Close()

	RespBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return "", ""
	}
	dalyPictureURL := gjson.Parse(string(RespBody)).Get("images").Array()[0].Get("url").String()
	url := "https://cn.bing.com" + dalyPictureURL

	return strings.Split(url, "&")[0], strings.Split(strings.Split(url, "&")[0], "=")[1]
}
