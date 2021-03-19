package helper_log

import "path/filepath"

type LogOption func(cfg *LogOptionConfig)

type LogOptionConfig struct {
	OutputPath string `json:"outputpath" yaml:"outputpath"`

	MaxSize    int  `json:"maxsize" yaml:"maxsize"`
	MaxAge     int  `json:"maxage" yaml:"maxage"`
	MaxBackups int  `json:"maxbackups" yaml:"maxbackups"`
	LocalTime  bool `json:"localtime" yaml:"localtime"`
	Compress   bool `json:"compress" yaml:"compress"`
}

func newLogOptionConfig(options ...LogOption) *LogOptionConfig {
	cfg := defaultLogOptionConfig()

	for i := range options {
		options[i](cfg)
	}

	cfg.OutputPath, _ = filepath.Abs(cfg.OutputPath)

	return cfg
}

func defaultLogOptionConfig() *LogOptionConfig {

	return &LogOptionConfig{
		OutputPath: "./log/runtime.log",
		MaxAge:     24 * 3, //保留旧文件的最大小时数
		MaxSize:    1024,   //在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: 3,      //保留旧文件的最大个数
		LocalTime:  true,   //
		Compress:   true,   //是否压缩/归档旧文件
	}
}
