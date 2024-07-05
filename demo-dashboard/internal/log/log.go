package log

import (
	"demo-dashboard/internal/conf"
	"io"
	"os"

	"github.com/gofiber/contrib/fiberzap/v2"
	fiberlog "github.com/gofiber/fiber/v2/log"
)

var Logger = fiberzap.NewLogger(fiberzap.LoggerConfig{
	ExtraKeys: []string{"requestId", "referer", "protocol", "port", "host", "path", "ua", "body", "queryParams", "bytesSent", "reqHeaders"},
})

var LogLevel = map[string]fiberlog.Level{
	"trace": fiberlog.LevelTrace,
	"debug": fiberlog.LevelDebug,
	"info":  fiberlog.LevelInfo,
	"warn":  fiberlog.LevelWarn,
	"error": fiberlog.LevelError,
	"fatal": fiberlog.LevelFatal,
	"panic": fiberlog.LevelPanic,
}

func init() {
	Logger.SetOutput(fileWriter(fiberlog.LevelError))
	Logger.SetLevel(getLogLevel())
}

func getLogLevel() fiberlog.Level {
	return LogLevel[conf.ErrorLogLevel]
}

func fileWriter(logType fiberlog.Level) io.Writer {
	logPath := conf.ErrorLogPath
	if logType == fiberlog.LevelInfo {
		logPath = conf.AccessLogPath
	}
	//standard output
	if logPath == "/dev/stdout" {
		return os.Stdout
	}
	if logPath == "/dev/stderr" {
		return os.Stderr
	}

	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fiberlog.Panicf("error: %v", err)
	}
	return file
}
