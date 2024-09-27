package logger

import (
	"demo-base/internal/conf"

	"github.com/gofiber/contrib/fiberzap/v2"
	"go.uber.org/zap/zapcore"
)

func LogLevel() []zapcore.Level {
	switch conf.LogConfig.Level {
	case "debug", "info":
		return []zapcore.Level{zapcore.ErrorLevel, zapcore.WarnLevel, zapcore.InfoLevel}
	case "warn":
		return []zapcore.Level{zapcore.WarnLevel, zapcore.InfoLevel}
	case "error", "fatal", "panic":
		return []zapcore.Level{zapcore.ErrorLevel}
	default:
		return []zapcore.Level{zapcore.InfoLevel, zapcore.WarnLevel, zapcore.ErrorLevel}
	}
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
