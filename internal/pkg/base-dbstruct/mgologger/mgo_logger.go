package mgologger

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"time"
)

type MgoLogger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})

	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
}

var MgoLog MgoLogger = NewMgoLogger()

func NewMgoLogger() *logrus.Logger {

	output := "./log/mgoextractor_runtime.log"
	writer, err := rotatelogs.New(
		output+".%Y-%m-%d-%H",
		rotatelogs.WithLinkName(output),        // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(4*time.Hour),     // 文件最大保存时间
		rotatelogs.WithRotationTime(time.Hour), // 日志切割时间间隔
	)
	if err != nil {
		return nil
	}

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(writer)
	logger.SetLevel(logrus.InfoLevel)
	return logger
}
