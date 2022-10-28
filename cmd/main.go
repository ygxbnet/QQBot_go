package cmd

import (
	"QQBot_go/internal/base"
	log "github.com/sirupsen/logrus"
	"os"
)

func Init() {
	log.Info("QQBot_go正在运行......")
	log.Infof("当前版本: %s", base.Version)
	Path, _ := os.Getwd()
	log.Info("位置:", Path)
}
