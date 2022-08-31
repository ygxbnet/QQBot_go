package cmd

import (
	"QQBot_go/connect"
	"QQBot_go/db"
	"QQBot_go/service"
	"fmt"
	"os"
)

func Main() {
	fmt.Println("程序正在运行......")
	Path, _ := os.Getwd()
	fmt.Println("位置:", Path)

	service.Services()

	db.CreateDBFile()

	connect.Connect()
}
