package group

import (
	"QQBot_go/internal/httpapi"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

func GetRandomPicture(groupID string, userID string, message string) {
	if len(strings.Fields(message)) == 1 {
		httpapi.SendGroupMsg(groupID, "[CQ:image,file=https://v.api.aa1.cn/api/api-fj-1/index.php?aa1=yuantu,cache=0]")
	} else if len(strings.Fields(message)) == 2 {
		num, err := strconv.Atoi(strings.Fields(message)[1])
		if err != nil {
			log.Error(err)
			httpapi.SendGroupMsg(groupID, fmt.Sprintf("[CQ:at,qq=%s]"+"\n❌指定图片数量错误", userID))
			return
		}
		for i := 0; i < num && i < 5; i++ {
			httpapi.SendGroupMsg(groupID, "[CQ:image,file=https://v.api.aa1.cn/api/api-fj-1/index.php?aa1=yuantu,cache=0]")
		}
	} else {
		httpapi.SendGroupMsg(groupID, "❌参数错误或多余")
	}
}
