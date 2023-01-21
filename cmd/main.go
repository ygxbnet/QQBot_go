package cmd

import (
	"QQBot_go/internal/base"
	log "github.com/sirupsen/logrus"
	"os"
)

func Init() {
	log.Info("QQBot_go正在运行......")
	log.Infof("当前版本: %s", base.VERSION)
	log.Infof("构建时间: %s", base.BUILD_TIME)
	Path, _ := os.Getwd()
	log.Info("位置:", Path)
}
