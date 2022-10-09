package api

import (
	"QQBot_go/config"
	"fmt"
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
		fmt.Println("请求[" + config.Http_url + Sndpoint + "]错误")

		fmt.Printf("body: %s\n", strings.Replace(fmt.Sprintf("%s", body), "\n", "\\n", -1))

		fmt.Println(err)
		return nil
	}

	defer resp.Body.Close()
	RespBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Println("请求[" + config.Http_url + Sndpoint + "]成功")

	fmt.Printf("body: %s\n", strings.Replace(fmt.Sprintf("%s", body), "\n", "\\n", -1))

	return RespBody
}
