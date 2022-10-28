package main

import (
	_ "QQBot_go/internal/log"

	_ "QQBot_go/internal/config"

	"QQBot_go/cmd"
	"QQBot_go/internal/connect"
	"QQBot_go/internal/db"
	"QQBot_go/service"
)

func main() {
	cmd.Init()

	service.Services()
	db.CreateDBFile()
	connect.Connect()
}
