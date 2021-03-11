package helper_log

import (
	"github.com/sirupsen/logrus"
)

/*

原文链接：

[Golang logrus 日志包及日志切割的实现 https://www.jb51.net/article/180448.htm](https://www.jb51.net/article/180448.htm)
[Logrus基本用法 https://www.jianshu.com/p/2d90b32acade](https://www.jianshu.com/p/2d90b32acade)
[bilibili代码基础日志组件 https://github.com/bilibili/sniper/blob/master/util/log/log.go#L61](https://github.com/bilibili/sniper/blob/master/util/log/log.go#L61)

GitHub：[https://github.com/sirupsen/logrus](https://github.com/sirupsen/logrus)
doc：[https://godoc.org/github.com/sirupsen/logrus](https://godoc.org/github.com/sirupsen/logrus)

*/

func NewLogrusLogger(options ...LogOption) (*logrus.Logger, error) {
	cfg := newLogOptionConfig(options...)

	writer, err := newLumberjackLogs(cfg)
	if err != nil {
		return nil, err
	}

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(writer)
	logger.SetLevel(logrus.DebugLevel)

	return logger, nil
}
