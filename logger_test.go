package logger

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"strings"
	"testing"
	"time"
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
		// 高亮
		WithEnableHighlighting(),
	)
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		logger.Sync()
		Sync()
	}()

	err = errors.New("pkg error")
	logger.Error("err occurs", WrapMeta(nil, NewMeta("para1", "value1"), NewMeta("para2", "value2"))...)
	logger.Error("err occurs", WrapMeta(err, NewMeta("para1", "value1"), NewMeta("para2", "value2"))...)

	logger.Info("err occurs", zap.String("key", "value"))
	logger.Warn("err occurs", zap.Any("key", "value"))
	logger.Error("err occurs", zap.Int("key", 1))
	logger.Debug("err occurs", zap.Int("key", 1))

	Debug("err occurs", zap.String("key", "err.Error()"))
	Info("err occurs", zap.String("key", "err.Error()"))
	Warn("err occurs", zap.String("key", "err.Error()"))
	Error("err occurs", zap.String("key", "err.Error()"))
}

func TestAtomicLeve(t *testing.T) {
	_, err := NewJSONLogger()
	assert.NoError(t, err)
	logger.Info("This is a info message")   // This should now be logged
	logger.Debug("This is a debug message") // This should now be logged

	http.HandleFunc("/loglevel", atomicLevel.ServeHTTP)
	go http.ListenAndServe(":8080", nil)

	time.Sleep(1 * time.Second) // Wait for the server to start

	// Test initial log level
	assert.Equal(t, zapcore.InfoLevel, atomicLevel.Level())

	// Change log level to debug
	reader := strings.NewReader(`{"level": "debug"}`)
	req, err := http.NewRequest("PUT", "http://localhost:8080/loglevel", reader)
	assert.NoError(t, err)
	do, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, do.StatusCode)

	// Verify log level is changed
	assert.Equal(t, zapcore.DebugLevel, atomicLevel.Level())
	logger.Info("This is a info message")   // This should now be logged
	logger.Debug("This is a debug message") // This should now be logged
}

func TestChangeLevelHandlerFunc(t *testing.T) {
	logger, err := NewJSONLogger()
	assert.NoError(t, err)
	logger.Info("This is a info message")   // This should now be logged
	logger.Debug("This is a debug message") // This should now be logged
	http.HandleFunc("/loglevel", ChangeLevelHandlerFunc())
	go http.ListenAndServe(":8080", nil)

	time.Sleep(1 * time.Second) // Wait for the server to start

	// Test initial log level
	assert.Equal(t, zapcore.InfoLevel, atomicLevel.Level())
	// Test initial log level
	assert.Equal(t, zapcore.InfoLevel, atomicLevel.Level())

	// Change log level to debug
	reader := strings.NewReader(`{"level": "debug"}`)
	req, err := http.NewRequest("PUT", "http://localhost:8080/loglevel", reader)
	assert.NoError(t, err)
	do, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, do.StatusCode)

	// Verify log level is changed
	assert.Equal(t, zapcore.DebugLevel, atomicLevel.Level())
	logger.Info("This is a info message")   // This should now be logged
	logger.Debug("This is a debug message") // This should now be logged
}
