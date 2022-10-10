package log

import (
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

func init() {
	//创建logs文件夹
	createDir("logs")

	logFile := setOutput()
	out := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(out)
	log.SetFormatter(&log.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15.04.05",
	})
	log.SetLevel(log.InfoLevel)

	timeEvent(logFile)
}

func getTime() string {
	return time.Now().Format("2006-01-02")
}

func createDir(path string) {
	err := os.Mkdir(path, 0755)
	if err != nil {
		log.Error(err)
	}
}

func setOutput() *os.File {
	fileName := getTime() + ".log"
	logFile, err := os.OpenFile("./logs/"+fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Error(err)
	}
	logFile.Write([]byte("\n-----------------------------------------------------------------------------------------------------\n"))
	return logFile
}

func timeEvent(writer *os.File) {
	go func() {
		var logFile = writer
		for true {
			timing(0, 0, 5)
			logFile.Close()

			logFile = setOutput()
			out := io.MultiWriter(os.Stdout, logFile)
			log.SetOutput(out)
			log.Info("已切换日志文件为: ", time.Now().Format("2006-01-02"), ".log")
		}
	}()
}

func timing(Hour int, Min int, Sec int) {
	now := time.Now()
	var next time.Time

	set_time := time.Date(now.Year(), now.Month(), now.Day(), Hour, Min, Sec, 0, now.Location())

	if set_time.After(now) {
		next = set_time
	} else {
		next = set_time.Add(24 * time.Hour)
		next = time.Date(next.Year(), next.Month(), next.Day(), Hour, Min, Sec, 0, now.Location())
	}

	t := time.NewTimer(next.Sub(now))
	<-t.C
}
