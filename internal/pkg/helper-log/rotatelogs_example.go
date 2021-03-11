package helper_log

import (
	"io"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

func newRotateLogs(cfg *LogOptionConfig) (io.Writer, error) {
	err := mkLogDir(cfg.OutputPath)
	if err != nil {
		return nil, err
	}

	writer, err := rotatelogs.New(
		cfg.OutputPath+".%Y-%m-%d-%H",
		rotatelogs.WithLinkName(cfg.OutputPath),                    // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(time.Duration(cfg.MaxAge)*time.Hour), // 文件最大保存时间
		rotatelogs.WithRotationTime(time.Hour),                     // 日志切割时间间隔
	)

	return writer, err
}
