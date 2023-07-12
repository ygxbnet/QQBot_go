package timeEvent

import (
	"QQBot_go/internal/config"
	"QQBot_go/internal/httpapi"
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"testing"
)

func TestDailyJoke(t *testing.T) {
	var msg []string

	for i := 0; i < 2; i++ {
		response, _ := http.Get("https://v.api.aa1.cn/api/api-wenan-gaoxiao/index.php?aa1=json")
		if response.StatusCode == 200 {
			bytes, _ := io.ReadAll(response.Body)
			msg = append(msg, gjson.Parse(string(bytes)).Get("0.gaoxiao").String())
		}
	}
	httpapi.SendGroupMsg(config.Get().Group.MainID, fmt.Sprintf("每日笑话二则：\n1. %s\n2. %s", msg[0], msg[1]))
}
