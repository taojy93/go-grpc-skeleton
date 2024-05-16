package global

import (
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
	sugar  *zap.SugaredLogger
	once   sync.Once
)

func InitLogger() error {
	var err error
	once.Do(func() {
		config := zap.NewProductionConfig()
		config.EncoderConfig.TimeKey = "timestamp"
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

		logger, err = config.Build()
		if err == nil {
			sugar = logger.Sugar()
		}
	})
	return err
}

func Logger() *zap.Logger {
	return logger
}

func Sugar() *zap.SugaredLogger {
	return sugar
}
