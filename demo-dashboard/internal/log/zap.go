package log

import (
	"os"

	fiberlog "github.com/gofiber/fiber/v2/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"demo-dashboard/internal/conf"
)

const (
	AccessLevel fiberlog.Level = iota - 1
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	PanicLevel
	FatalLevel
)

func getLogLevel() fiberlog.Level {
	level := WarnLevel
	switch conf.ErrorLogLevel {
	case "debug":
		level = DebugLevel
	case "info":
		level = InfoLevel
	case "warn":
		level = WarnLevel
	case "error":
		level = ErrorLevel
	case "panic":
		level = PanicLevel
	case "fatal":
		level = FatalLevel
	}
	return level
}

func fileWriter(logType fiberlog.Level) zapcore.WriteSyncer {
	logPath := conf.ErrorLogPath
	if logType == AccessLevel {
		logPath = conf.AccessLogPath
	}
	//standard output
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
