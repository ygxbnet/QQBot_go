package log

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strings"
	"time"
)

var logFile *os.File

// init 初始化log
func init() {
	// 创建logs文件夹
	createDir("logs")

	log.AddHook(&MyHook{})
	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		ForceColors:      true,
		FullTimestamp:    true,
		TimestampFormat:  "2006-01-02 15:04:05",
		CallerPrettyfier: func(*runtime.Frame) (function string, file string) { return "", "" },
	})
	log.SetLevel(log.InfoLevel)

	timeEvent()
}

func getTime() string {
	return time.Now().Format("2006-01-02")
}

func createDir(path string) {
	_, err := os.Stat(path)
	if err != nil {
		err = os.Mkdir(path, 0755)
		if err != nil {
			log.Error(err)
		}
	}
}

func setOutput() *os.File {
	fileName := getTime() + ".log"
	logFile, err := os.OpenFile("./logs/"+fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Error(err)
	}
	return logFile
}

func timeEvent() {
	go func() {
		logFile = setOutput()
		_, err := logFile.Write([]byte("\n-----------------------------------------------------------------------------------------------------\n"))
		if err != nil {
			log.Error(err)
		}

		for true {
			timing(0, 0, 0)
			err := logFile.Close()
			if err != nil {
				log.Error(err)
			}
			logFile = setOutput()

			log.Info("已切换日志文件为: ", time.Now().Format("2006-01-02"), ".log")
		}
	}()
}

func timing(Hour int, Min int, Sec int) {
	now := time.Now()
	var next time.Time

	setTime := time.Date(now.Year(), now.Month(), now.Day(), Hour, Min, Sec, 0, now.Location())

	if setTime.After(now) {
		next = setTime
	} else {
		next = setTime.Add(24 * time.Hour)
		next = time.Date(next.Year(), next.Month(), next.Day(), Hour, Min, Sec, 0, now.Location())
	}

	t := time.NewTimer(next.Sub(now))
	<-t.C
}

// MyHook ...
type MyHook struct {
}

// Levels 定义所有等级触发hook
func (h *MyHook) Levels() []log.Level {
	return log.AllLevels
}

// Fire 将日志写入到指定日志文件中 并且将异常日志写入到指定日志文件中
func (h *MyHook) Fire(entry *log.Entry) error {

	_, err := logFile.Write([]byte(fmt.Sprintf("[%s] [%s] %v\n",
		entry.Time.Format("2006-01-02 15:04:05"),
		strings.ToUpper(entry.Level.String()),
		strings.Replace(entry.Message, "\n", "\\n", -1))))
	if err != nil {
		return err
	}

	if entry.Level == log.ErrorLevel || entry.Level == log.FatalLevel || entry.Level == log.PanicLevel {

		errorFile, err := os.OpenFile("./logs/error-"+getTime()+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		defer errorFile.Close()

		if err != nil {
			return err
		}

		_, err = errorFile.Write([]byte(fmt.Sprintf("[%s] [%s] [%v:%v] %v\n",
			entry.Time.Format("2006-01-02 15:04:05"),
			strings.ToUpper(entry.Level.String()),
			entry.Caller.File,
			entry.Caller.Line,
			strings.Replace(entry.Message, "\n", "\\n", -1))))
		if err != nil {
			return err
		}
	}
	return nil
}
