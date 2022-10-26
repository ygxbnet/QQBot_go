package main

import (
	_ "QQBot_go/internal/log"

	"QQBot_go/internal/base"
	"QQBot_go/internal/connect"
	"QQBot_go/internal/db"
	"QQBot_go/service"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	log.Info("QQBot_go正在运行......")
	log.Infof("当前版本: %s", base.Version)
	Path, _ := os.Getwd()
	log.Info("位置:", Path)

	service.Services()
	db.CreateDBFile()
	connect.Connect()
}
