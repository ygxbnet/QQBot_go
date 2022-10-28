package config

import (
	_ "embed"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
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
		log.Warn("请修改配置文件后再重新启动")
		os.Exit(0)
	}
}

// Parse 从默认配置文件路径中获取
func Parse() Config {
	file, err := os.ReadFile("config.yml")
	if err != nil {
		log.Error("读取配置文件错误", err)
		os.Exit(1)
	}

	config := Config{}
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Error("解析配置文件错误", err)
		os.Exit(1)
	}
	return config
}
