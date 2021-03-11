package helper_log

import (
	"io"

	"github.com/natefinch/lumberjack"
)

func newLumberjackLogs(cfg *LogOptionConfig) (io.Writer, error) {
	err := mkLogDir(cfg.OutputPath)
	if err != nil {
		return nil, err
	}

	writer := &lumberjack.Logger{
		Filename:   cfg.OutputPath, //日志文件的位置
		MaxSize:    100,            //在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: 10,             //保留旧文件的最大个数
		MaxAge:     cfg.MaxAge,     //保留旧文件的最大天数
		Compress:   false,          //是否压缩/归档旧文件
	}

	return writer, nil
}
