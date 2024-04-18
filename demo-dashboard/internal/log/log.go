package log

import "github.com/gofiber/contrib/fiberzap/v2"

var Logger *fiberzap.LoggerConfig

func init() {
	Logger.SetOutput(fileWriter(ErrorLevel))
	Logger.SetLevel(getLogLevel())
}
