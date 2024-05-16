package global

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var G_Logger *Logger

func InitLogger() (err error) {
	logger, err := NewLogger()
	if err != nil {
		return err
	}
	G_Logger = logger
	return
}

type Logger struct {
	NormalLog *zap.Logger
	SugerLog  *zap.SugaredLogger
}

func NewLogger() (*Logger, error) {

	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	return &Logger{
		NormalLog: logger,
		SugerLog:  logger.Sugar(),
	}, nil
}
