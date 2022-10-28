package main

import (
	"QQBot_go/cmd"
	"QQBot_go/internal/config"
	"QQBot_go/internal/connect"
	"QQBot_go/internal/db"
	"QQBot_go/internal/log"
	"QQBot_go/service"
)

func main() {
	log.Init()
	config.Init()
	cmd.Init()

	service.Services()
	db.CreateDBFile()
	connect.Connect()
}
