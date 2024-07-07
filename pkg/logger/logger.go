package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Info(msg string, keysAndValues ...interface{})
	Error(msg string, keysAndValues ...interface{})
}

type ZapLogger struct {
	*zap.SugaredLogger
}

func (z *ZapLogger) Info(msg string, keysAndValues ...interface{}) {
	z.SugaredLogger.Infow(msg, keysAndValues...)
}

func (z *ZapLogger) Error(msg string, keysAndValues ...interface{}) {
	z.SugaredLogger.Errorw(msg, keysAndValues...)
}

func New(level string) Logger {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	config.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	if level == "debug" {
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}

	logger, _ := config.Build()
	return &ZapLogger{logger.Sugar()}
}
