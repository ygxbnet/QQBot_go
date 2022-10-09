package log

import (
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

func init() {
	//创建logs文件夹
	err := createDir("logs")
	if err != nil {
		log.Error(err)
	}

	fileName := getTime() + ".log"
	logFile, err := os.OpenFile("./logs/"+fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Error(err)
	}
	logFile.Write([]byte("\n-----------------------------------------------------------------------------------------------------\n"))
	out := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(out)

	log.SetFormatter(&log.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15.04.05",
	})
	log.SetLevel(log.InfoLevel)
}

func getTime() string {
	return time.Now().Format("2006-01-02")
}

func createDir(path string) error {
	err := os.Mkdir(path, 0755)
	if err != nil {
		return err
	}
	return nil
}
