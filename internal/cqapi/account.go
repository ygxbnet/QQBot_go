package cqapi

import log "github.com/sirupsen/logrus"

// GetLoginInfo 获取登录号信息
func GetLoginInfo() (userId string, nickName string) {
	body := sendHTTP("/get_login_info", nil)
	log.Infof("获取登录号信息: %s", string(body))
	return "", ""
}
