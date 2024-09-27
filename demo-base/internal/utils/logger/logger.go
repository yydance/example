package logger

import (
	"demo-base/internal/conf"

	"github.com/gofiber/contrib/fiberzap/v2"
	fiberlog "github.com/gofiber/fiber/v2/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var LogLevels = map[string]fiberlog.Level{
	"debug": fiberlog.LevelDebug,
	"info":  fiberlog.LevelInfo,
	"warn":  fiberlog.LevelWarn,
	"error": fiberlog.LevelError,
	"fatal": fiberlog.LevelFatal,
	"panic": fiberlog.LevelPanic,
}

var LogLevel []zapcore.Level
var level = conf.LogConfig.Level

switch level {
case "debug","info":
	LogLevel = []zapcore.Level{zapcore.InfoLevel, zapcore.WarnLevel, zapcore.ErrorLevel}
case "warn":
	LogLevel = []zapcore.Level{zapcore.InfoLevel,zapcore.WarnLevel}
case "error","fatal","panic":
	LogLevel = []zapcore.Level{zapcore.ErrorLevel}
default:
	LogLevel = []zapcore.Level{zapcore.InfoLevel, zapcore.WarnLevel, zapcore.ErrorLevel}
}

var Logger = fiberzap.NewLogger(fiberzap.LoggerConfig{
	ExtraKeys: []string{"requestid"},
})

func Trace(v ...interface{}) {
	Logger.Trace(v...)
}

func Debug(v ...interface{}) {
	Logger.Debug(v...)
}

func Info(v ...interface{}) {
	Logger.Info(v...)
}

func Warn(v ...interface{}) {
	Logger.Warn(v...)
}

func Error(v ...interface{}) {
	Logger.Error(v...)
}

func Fatal(v ...interface{}) {
	Logger.Fatal(v...)
}

func Panic(v ...interface{}) {
	Logger.Panic(v...)
}

func Tracef(format string, v ...interface{}) {
	Logger.Tracef(format, v...)
}

func Debugf(format string, v ...interface{}) {
	Logger.Debugf(format, v...)
}

func Infof(format string, v ...interface{}) {
	Logger.Infof(format, v...)
}

func Warnf(format string, v ...interface{}) {
	Logger.Warnf(format, v...)
}

func Errorf(format string, v ...interface{}) {
	Logger.Errorf(format, v...)
}

func Fatalf(format string, v ...interface{}) {
	Logger.Fatalf(format, v...)
}

func Panicf(format string, v ...interface{}) {
	Logger.Panicf(format, v...)
}

func Tracew(msg string, keysAndValues ...interface{}) {
	Logger.Tracew(msg, keysAndValues...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	Logger.Debugw(msg, keysAndValues...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	Logger.Infow(msg, keysAndValues...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	Logger.Warnw(msg, keysAndValues...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	Logger.Errorw(msg, keysAndValues...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	Logger.Fatalw(msg, keysAndValues...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	Logger.Panicw(msg, keysAndValues...)
}

