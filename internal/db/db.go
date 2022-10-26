package db

import (
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"os"
)

var dbFileName = "data.json"

// CreateDBFile 创建DB文件
func CreateDBFile() {
	path, _ := os.Getwd()
	_, err := os.Stat(path + "/" + dbFileName)
	if err != nil {
		log.Info("未找到:", dbFileName)
		file, _ := os.Create(dbFileName)
		log.Info("已创建:", file.Name())
		os.WriteFile(dbFileName, []byte("{\"guild\":{},\"group\":{}}"), 0644)
	} else {
		log.Info("文件:", dbFileName, "已存在")
	}
}

// ReadDBFile 读取DB文件
func ReadDBFile(option string, userID string) string {
	path, _ := os.Getwd()
	file, _ := os.ReadFile(path + "/" + dbFileName)
	return gjson.Parse(string(file)).Get(option).Get(userID).String()
}

// WriteDBFile 写入DB文件
func WriteDBFile(option string, userID string, data interface{}) {
	path, _ := os.Getwd()
	file, _ := os.ReadFile(path + "/" + dbFileName)
	json, _ := sjson.Set(string(file), option+"."+userID, data)
	_ = os.WriteFile(path+"/"+dbFileName, []byte(json), 0644)
}
