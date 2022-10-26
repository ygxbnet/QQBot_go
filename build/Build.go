package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

var version = ""
var returnVersion = ""
var command = "cd ../" +
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
	data, err := os.ReadFile("./build/version")
	if err != nil {
		fmt.Println(err)
		return
	}
	version = string(data)
	fmt.Println("当前版本号: ", version)

	fmt.Println("请输入要构建的版本号(默认为: ", version, ")")
	fmt.Scanln(&returnVersion)
	if returnVersion == "" {
		returnVersion = version
	}

	command := fmt.Sprintf(command, returnVersion)
	cmd := exec.Command("cmd", command)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Println(out)

	os.WriteFile("./build/version", []byte(returnVersion), 0666)
}
