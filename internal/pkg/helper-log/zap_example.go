package helper_log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/*

原文链接：

[如何开发高性能 go 组件？https://zhuanlan.zhihu.com/p/41991119](https://zhuanlan.zhihu.com/p/41991119)
[在Go语言项目中使用Zap日志库 https://zhuanlan.zhihu.com/p/88856378](https://zhuanlan.zhihu.com/p/88856378)
[Go 每日一库之 zap https://zhuanlan.zhihu.com/p/136093026](https://zhuanlan.zhihu.com/p/136093026)

GitHub：[https://github.com/uber-go/zap](https://github.com/uber-go/zap)
doc：[https://pkg.go.dev/go.uber.org/zap?tab=doc](https://pkg.go.dev/go.uber.org/zap?tab=doc)

*/

func NewZapLogger(options ...LogOption) (*zap.SugaredLogger, error) {
	cfg := newLogOptionConfig(options...)

	writer, err := newRotateLogs(cfg)
	if err != nil {
		return nil, err
	}
	writeSyncer := zapcore.AddSync(writer)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	//zap提供了便捷的方法SugarLogger，可以使用printf格式符的方式。
	//SugaredLogger的使用比Logger简单，只是性能比Logger低 50% 左右，可以用在非热点函数中。
	sugarLogger := logger.Sugar()

	return sugarLogger, nil
}

//zap提供了几个快速创建logger的方法，zap.NewExample()、zap.NewDevelopment()、zap.NewProduction()，还有高度定制化的创建方法zap.New()。
//创建前 3 个logger时，zap会使用一些预定义的设置，它们的使用场景也有所不同。
//Example适合用在测试代码中，Development在开发环境中使用，Production用在生成环境。

func NewZapExample() *zap.Logger {
	return zap.NewExample()
}

func NewZapDevelopment() (*zap.Logger, error) {
	return zap.NewDevelopment()
}

func NewZapProduction() (*zap.Logger, error) {
	return zap.NewProduction()
}
