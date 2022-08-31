package services_api

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func Get_Bing_Picture_URL() (url_img string, name_img string) {

	resp, err := http.Post("https://cn.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1&mkt=zh-CN", "", nil)
	if err != nil {
		fmt.Println("请求[https://cn.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1&mkt=zh-CN]错误")
		fmt.Println(err)
		return "", ""
	}
	defer resp.Body.Close()

	RespBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "", ""
	}
	daly_picture_url := gjson.Parse(string(RespBody)).Get("images").Array()[0].Get("url").String()
	url := "https://cn.bing.com" + daly_picture_url

	return strings.Split(url, "&")[0], strings.Split(strings.Split(url, "&")[0], "=")[1]
}

func Get_Random_Picture_URL() (url_img string, name_img string) {

	resp, err := http.Post("https://api.ixiaowai.cn/api/api.php?return=json", "", nil)
	if err != nil {
		fmt.Println("请求[https://api.ixiaowai.cn/api/api.php?return=json]错误")
		fmt.Println(err)
		return "", ""
	}
	defer resp.Body.Close()

	RespBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("解析[https://api.ixiaowai.cn/api/api.php?return=json]请求内容错误")
		fmt.Println(err)
		return "", ""
	}
	fmt.Println(string(RespBody))
	random_picture_url := gjson.Parse(string(RespBody)).Get("imgurl").String()
	random_picture_name := strings.Split(random_picture_url, "/")[len(strings.Split(random_picture_url, "/"))-1]

	return random_picture_url, random_picture_name
}
