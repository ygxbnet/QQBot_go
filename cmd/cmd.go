package cmd

import (
	"QQBot_go/connect"
	"QQBot_go/db"
	"QQBot_go/service"
	log "github.com/sirupsen/logrus"
	"os"

	_ "QQBot_go/log"
)

func Main() {
	log.Infoln("程序正在运行......")
	Path, _ := os.Getwd()
	log.Infoln("位置:", Path)

	service.Services()

	db.CreateDBFile()

	connect.Connect()
}
