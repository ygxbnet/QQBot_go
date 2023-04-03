package main

import (
	_ "QQBot_go/internal/log"

	_ "QQBot_go/internal/config"
)
import (
	"QQBot_go/internal/base"
	"QQBot_go/internal/connect"
	"QQBot_go/internal/db"
	"QQBot_go/service"
	log "github.com/sirupsen/logrus"
	"os"
	"sync"
)

func main() {
	log.Info("QQBot_go正在运行......")
	log.Infof("当前版本: %s", base.VERSION)
	log.Infof("构建时间: %s", base.BUILD_TIME)
	Path, _ := os.Getwd()
	log.Info("运行位置: ", Path)

	db.CreateDBFile()
	connect.Connect()
	go service.Services()

	// 阻塞主进程
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
