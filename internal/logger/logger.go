package logger

import (
	"go.uber.org/zap"
)

type logger struct {
	logger *zap.SugaredLogger
}

type Logger interface {
	Info(string, ...interface{})
}

func New(cfg interface{}) Logger {
	l, _ := zap.NewProduction()
	sugar := l.Sugar()
	return &logger{
		logger: sugar,
	}
}

func (l *logger) Info(msg string, fields ...interface{}) {
	l.logger.Infow(msg, fields...)
}
