package logger

import (
	"demo-base/internal/conf"
	"os"
	"time"

	"github.com/gofiber/contrib/fiberzap/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02T15:04:05.000Z"))
}

func customLevelEnabler(lvl zapcore.Level) bool {
	if conf.LogLevel == "info" {
		switch lvl {
		case zapcore.DebugLevel:
			return false
		default:
			return true
		}
	}
	if conf.LogLevel == "warn" {
		switch lvl {
		case zapcore.DebugLevel, zapcore.InfoLevel:
			return false
		default:
			return true
		}
	}
	if conf.LogLevel == "error" {
		switch lvl {
		case zapcore.DebugLevel, zapcore.InfoLevel, zapcore.WarnLevel:
			return false
		default:
			return true
		}
	}
	return true
}

func NewCustomLogger() *zap.Logger {
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.EncodeTime = customTimeEncoder
	enabler := zap.LevelEnablerFunc(customLevelEnabler)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encodeConfig),
		zapcore.Lock(os.Stdout),
		enabler,
	)
	return zap.New(core)
}

var (
	logger = fiberzap.NewLogger(fiberzap.LoggerConfig{
		ExtraKeys: []string{"requestid"},
		SetLogger: NewCustomLogger(),
	})
)

func Trace(v ...interface{}) {
	logger.Trace(v...)
}

func Debug(v ...interface{}) {
	logger.Debug(v...)
}

func Info(v ...interface{}) {
	logger.Info(v...)
}

func Warn(v ...interface{}) {
	logger.Warn(v...)
}

func Error(v ...interface{}) {
	logger.Error(v...)
}

func Fatal(v ...interface{}) {
	logger.Fatal(v...)
}

func Panic(v ...interface{}) {
	logger.Panic(v...)
}

func Tracef(format string, v ...interface{}) {
	logger.Tracef(format, v...)
}

func Debugf(format string, v ...interface{}) {
	logger.Debugf(format, v...)
}

func Infof(format string, v ...interface{}) {
	logger.Infof(format, v...)
}

func Warnf(format string, v ...interface{}) {
	logger.Warnf(format, v...)
}

func Errorf(format string, v ...interface{}) {
	logger.Errorf(format, v...)
}

func Fatalf(format string, v ...interface{}) {
	logger.Fatalf(format, v...)
}

func Panicf(format string, v ...interface{}) {
	logger.Panicf(format, v...)
}

func Tracew(msg string, keysAndValues ...interface{}) {
	logger.Tracew(msg, keysAndValues...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	logger.Debugw(msg, keysAndValues...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	logger.Infow(msg, keysAndValues...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	logger.Warnw(msg, keysAndValues...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	logger.Errorw(msg, keysAndValues...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	logger.Fatalw(msg, keysAndValues...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	logger.Panicw(msg, keysAndValues...)
}
