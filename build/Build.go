package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

var Version = ""
var ReturnVersion = ""
var Command = "cd ../" +
	"\ngo env -w CGO_ENABLED=0" +
	"\ngo env -w GOOS=linux" +
	"\ngo env -w GOARCH=arm" +
	"\ngo env -w GOARM=7" +
	"\ngo build -ldflags -X QQBot_go/data.version=%s" +
	"\ngo env -w CGO_ENABLED=1" +
	"\ngo env -w GOOS=windows" +
	"\ngo env -w GOARCH=amd64"

func main() {
	fmt.Println("构建开始")

	fmt.Println("正在读取版本")
	data, err := os.ReadFile("./build/Version")
	if err != nil {
		fmt.Println(err)
		return
	}
	Version = string(data)
	fmt.Println("当前版本号：", Version)

	fmt.Println("请输入要构建的版本号(默认为：", Version, ")")
	fmt.Scanln(&ReturnVersion)
	if ReturnVersion == "" {
		ReturnVersion = Version
	}

	command := fmt.Sprintf(Command, ReturnVersion)
	cmd := exec.Command("cmd", command)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Println(out)

	os.WriteFile("./build/Version", []byte(ReturnVersion), 0666)
}
