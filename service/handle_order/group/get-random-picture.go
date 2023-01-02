package group

import (
	"QQBot_go/internal/httpapi"
	"fmt"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"strconv"
	"strings"
)

func GetRandomPicture(groupID string, userID string, message string) {
	if len(strings.Fields(message)) == 1 {
		httpapi.SendGroupMsg(groupID, "[CQ:image,file=https://api.ixiaowai.cn/gqapi/gqapi.php?a="+strconv.Itoa(rand.Int())+"]")
	} else if len(strings.Fields(message)) == 2 {
		num, err := strconv.Atoi(strings.Fields(message)[1])
		if err != nil {
			log.Error(err)
			httpapi.SendGroupMsg(groupID, fmt.Sprintf("[CQ:at,qq=%s]"+"\n❌指定图片数量错误", userID))
			return
		}
		for i := 0; i < num && i < 10; i++ {
			httpapi.SendGroupMsg(groupID, "[CQ:image,file=https://api.ixiaowai.cn/gqapi/gqapi.php?a="+strconv.Itoa(rand.Int())+"]")
		}
	} else {
		httpapi.SendGroupMsg(groupID, "❌参数错误或多余")
	}
}
