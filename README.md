
# go-logger

## 关于
对于zap框架的封装使用，我们结合了lumberjack来实现日志分割，你可以更轻便的使用zap库来开发自己的系统

## 快速开始

你可以通过初始化的配置来自定义自己的日志模式，我们提供了默认值`DefaultLevel=zapcore.InfoLevel`和`DefaultTimeLayout=time.RFC3339`，我们建议您可以在开发模式下使用debug等级同时配合控制台打印日志，在生产环境下使用info同时关闭日志打印
```go
package logger  
  
import (  
   "errors"  
   "go.uber.org/zap"   "testing")  
// json日志配置
func TestNewJSONLogger(t *testing.T) {  
   logger, err := NewJSONLogger(  
      // 日志等级  
      WithDebugLevel(),  
      // 写出的文件  
      WithFileRotationP("/log/test.log"),  
      // 不在控制台打印  
      WithDisableConsole(),  
   )  
   if err != nil {  
      t.Fatal(err)  
   }  
  
   defer logger.Sync()  
  
   err = errors.New("pkg error")  
   logger.Error("err occurs", WrapMeta(nil, NewMeta("para1", "value1"), NewMeta("para2", "value2"))...)  
   logger.Error("err occurs", WrapMeta(err, NewMeta("para1", "value1"), NewMeta("para2", "value2"))...)  
  
   logger.Info("err occurs", zap.String("key", "value"))  
   logger.Warn("err occurs", zap.Any("key", "value"))  
   logger.Error("err occurs", zap.Int("key", 1))  
   logger.Debug("err occurs", zap.Int("key", 1))  
}
```
打印出来的日志
```json
{"level":"error","time":"2022-12-14T08:50:24+08:00","caller":"logger/logger_test.go:25","msg":"err occurs","meta":{"para1":"value1","para2":"value2"}}
{"level":"error","time":"2022-12-14T08:50:24+08:00","caller":"logger/logger_test.go:26","msg":"err occurs","error":"pkg error","meta":{"para1":"value1","para2":"value2"}}
{"level":"info","time":"2022-12-14T08:50:24+08:00","caller":"logger/logger_test.go:28","msg":"err occurs","key":"value"}
{"level":"warn","time":"2022-12-14T08:50:24+08:00","caller":"logger/logger_test.go:29","msg":"err occurs","key":"value"}
{"level":"error","time":"2022-12-14T08:50:24+08:00","caller":"logger/logger_test.go:30","msg":"err occurs","key":1}
{"level":"debug","time":"2022-12-14T08:50:24+08:00","caller":"logger/logger_test.go:31","msg":"err occurs","key":1}
```
## API

我们还提供了对外使用的API，您可以直接使用API来打印日志。我们提供了四个对外开发的API`Info` `Debug`  `Warn` `Error`，即使没有进行NewJSONLogger的初始化操作，你也可以默认进行控制台的打印操作
```go
Debug("err occurs", zap.String("key", "err.Error()"))  
Info("err occurs", zap.String("key", "err.Error()"))  
Warn("err occurs", zap.String("key", "err.Error()"))  
Error("err occurs", zap.String("key", "err.Error()"))
```
上面这段代码将打印出以下数据，默认不打印debug级别的日志
```json
{"level":"info","time":"2022-12-14T13:25:50+08:00","caller":"logger/logger.go:262","msg":"err occurs","key":"err.Error()"}
{"level":"warn","time":"2022-12-14T13:25:50+08:00","caller":"logger/logger.go:272","msg":"err occurs","key":"err.Error()"}
{"level":"error","time":"2022-12-14T13:25:50+08:00","caller":"logger/logger.go:277","msg":"err occurs","key":"err.Error()"}
```
如何您进行了NewJSONLogger的初始化，我们将使用您提供的logger对象替换掉内部的全局logger对象，实现对外API的改变，使日志操作变得更简单，下面为您提供使用示例
```go
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
  
   Debug("err occurs", zap.String("key", "err.Error()"))  
   Info("err occurs", zap.String("key", "err.Error()"))  
   Warn("err occurs", zap.String("key", "err.Error()"))  
   Error("err occurs", zap.String("key", "err.Error()"))  
}
```