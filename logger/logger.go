package logger

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

var logrusInstance *logrus.Logger
var logrusInitOnce sync.Once

func getLogrus() *logrus.Logger {
	logrusInitOnce.Do(func() {
		logrusInstance = logrus.New()
		logrusInstance.SetFormatter(&nested.Formatter{
			HideKeys:        true,
			TimestampFormat: time.RFC3339,
			FieldsOrder:     []string{"component", "category"},
		})
	})
	return logrusInstance
}

func Debug(args ...interface{}) {
	//logger(msg, 0)
	getLogrus().Debug(args)
}

func Warn(args ...interface{}) {
	getLogrus().Warn(args)
	//logger(msg, 1)
}

func Error(msg string, err error) {
	if err == nil {
		getLogrus().Errorf(msg)
	} else {
		getLogrus().Errorf(msg, err)
	}
	//logger(msg, 2)
	//if err != nil {
	//logger(err.Error(), 2)
	//}
}

func Info(args ...interface{}) {
	getLogrus().Info(args)
	//logger(msg, 3)
}

//var logTypeList = []string{"DEBG", "WARN", "ERRO", "INFO"}
//var logColorList = []int{0, 33, 31, 36}
//
//func log(msg string, logType int) {
//	fmt.Printf("%c[1;0;%dm[%s %s]%c[0m %s\n", 0x1B, logColorList[logType], logTypeList[logType], getCurrentTimeFormatStr(), 0x1B, msg)
//}
//
//func getCurrentTimeFormatStr() string {
//	return time.Now().Format("2006-01-02 15:04:05")
//}
