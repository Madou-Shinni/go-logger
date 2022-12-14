package logger

import (
	"go.uber.org/zap"
	"testing"
)

func TestNewJSONLogger(t *testing.T) {
	logger, err := NewJSONLogger(
		// 日志等级
		WithDebugLevel(),
		// 写出的文件
		WithFileRotationP("/log/test.log"),
		// 不在控制台打印
		WithDisableConsole(),
		// 时间格式化
		WithTimeLayout("2006-01-02 15:04:05"),
	)
	if err != nil {
		t.Fatal(err)
	}

	defer logger.Sync()
	//
	//err = errors.New("pkg error")
	//logger.Error("err occurs", WrapMeta(nil, NewMeta("para1", "value1"), NewMeta("para2", "value2"))...)
	//logger.Error("err occurs", WrapMeta(err, NewMeta("para1", "value1"), NewMeta("para2", "value2"))...)
	//
	//logger.Info("err occurs", zap.String("key", "value"))
	//logger.Warn("err occurs", zap.Any("key", "value"))
	//logger.Error("err occurs", zap.Int("key", 1))
	//logger.Debug("err occurs", zap.Int("key", 1))

	Debug("err occurs", zap.String("key", "err.Error()"))
	Info("err occurs", zap.String("key", "err.Error()"))
	Warn("err occurs", zap.String("key", "err.Error()"))
	Error("err occurs", zap.String("key", "err.Error()"))
}
