package db

import (
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"io/ioutil"
	"os"
)

var DBFileName = "data.json"

func CreateDBFile() {
	path, _ := os.Getwd()
	_, err := os.Stat(path + "/" + DBFileName)
	if err != nil {
		fmt.Println("未找到:", DBFileName)
		file, _ := os.Create(DBFileName)
		fmt.Println("已创建:", file.Name())
		ioutil.WriteFile(DBFileName, []byte("{\"guild\":{},\"group\":{}}"), 0644)
	} else {
		fmt.Println("文件:", DBFileName, "已存在")
	}
}
func ReadDBFile(option string, user_id string) string {
	path, _ := os.Getwd()
	file, _ := ioutil.ReadFile(path + "/" + DBFileName)
	return gjson.Parse(string(file)).Get(option).Get(user_id).String()
}
func WriteDBFile(option string, user_id string, data interface{}) {
	path, _ := os.Getwd()
	file, _ := ioutil.ReadFile(path + "/" + DBFileName)
	json, _ := sjson.Set(string(file), option+"."+user_id, data)
	_ = ioutil.WriteFile(path+"/"+DBFileName, []byte(json), 0644)
}
