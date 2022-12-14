
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