package config

import (
	_ "embed"
	log "github.com/sirupsen/logrus"
	"os"
)

//go:embed config-default.yml
var defaultConfig string

func init() {
	_, err := os.Stat("config.yml")
	if os.IsNotExist(err) {
		file, err := os.Create("config.yml")
		if err != nil {
			log.Error(err)
			return
		}
		defer file.Close()

		_, err = file.Write([]byte(defaultConfig))
		if err != nil {
			log.Error(err)
			return
		}
		log.Info("未发现配置文件，已创建 config.yml")
	}

	//TODO
}

// WebSocketURL WebSocket地址
var WebSocketURL = "ws://192.168.3.45:8080"

// HTTPURL Http地址
var HTTPURL = "http://192.168.3.45:5700"
