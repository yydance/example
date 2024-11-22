package logger

import (
	"demo-base/internal/conf"
	"os"

	"github.com/gofiber/contrib/fiberzap/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	AccessLog = iota
	ErrorLog
)

func NewCustomLogger(logType int) *zap.Logger {
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	writeSyncer := fileWriter(logType)
	logLevel := getLogLevel()

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encodeConfig),
		writeSyncer,
		logLevel,
	)
	return zap.New(core)
}

func getLogLevel() zapcore.LevelEnabler {
	level := zapcore.WarnLevel
	switch conf.LogLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	case "fatal":
		level = zapcore.FatalLevel
	}

	return level

}

func fileWriter(logType int) zapcore.WriteSyncer {
	logPath := conf.ErrorLog
	if logType == AccessLog {
		logPath = conf.AccessLog
	}
	if logPath == "/dev/stdout" {
		return zapcore.Lock(os.Stdout)
	}
	if logPath == "/dev/stderr" {
		return zapcore.Lock(os.Stderr)
	}

	writer, _, err := zap.Open(logPath)
	if err != nil {
		panic(err)
	}
	return writer
}

func newLoggerConfig() *fiberzap.LoggerConfig {
	return fiberzap.NewLogger(fiberzap.LoggerConfig{
		ExtraKeys: []string{"requestid"},
		SetLogger: NewCustomLogger(ErrorLog),
	})
}

func Trace(v ...interface{}) {
	newLoggerConfig().Trace(v...)
}

func Debug(v ...interface{}) {
	newLoggerConfig().Debug(v...)
}

func Info(v ...interface{}) {
	newLoggerConfig().Info(v...)
}

func Warn(v ...interface{}) {
	newLoggerConfig().Warn(v...)
}

func Error(v ...interface{}) {
	newLoggerConfig().Error(v...)
}

func Fatal(v ...interface{}) {
	newLoggerConfig().Fatal(v...)
}

func Panic(v ...interface{}) {
	newLoggerConfig().Panic(v...)
}

func Tracef(format string, v ...interface{}) {
	newLoggerConfig().Tracef(format, v...)
}

func Debugf(format string, v ...interface{}) {
	newLoggerConfig().Debugf(format, v...)
}

func Infof(format string, v ...interface{}) {
	newLoggerConfig().Infof(format, v...)
}

func Warnf(format string, v ...interface{}) {
	newLoggerConfig().Warnf(format, v...)
}

func Errorf(format string, v ...interface{}) {
	newLoggerConfig().Errorf(format, v...)
}

func Fatalf(format string, v ...interface{}) {
	newLoggerConfig().Fatalf(format, v...)
}

func Panicf(format string, v ...interface{}) {
	newLoggerConfig().Panicf(format, v...)
}

func Tracew(msg string, keysAndValues ...interface{}) {
	newLoggerConfig().Tracew(msg, keysAndValues...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	newLoggerConfig().Debugw(msg, keysAndValues...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	newLoggerConfig().Infow(msg, keysAndValues...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	newLoggerConfig().Warnw(msg, keysAndValues...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	newLoggerConfig().Errorw(msg, keysAndValues...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	newLoggerConfig().Fatalw(msg, keysAndValues...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	newLoggerConfig().Panicw(msg, keysAndValues...)
}
