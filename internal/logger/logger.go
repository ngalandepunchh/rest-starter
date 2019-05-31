package logger

import (
	// Internal
	"github.com/hunterpunchh/rest-starter/internal/config"

	//External
	"go.uber.org/zap"
)

func setupLogger(cfg *config.Config) {
	return
}

// New creates and returns *zap.SugaredLogger
// Call logger.Desugar() to get the structured logger
// For more info: https://godoc.org/go.uber.org/zap
func New(cfg *config.Config) *zap.SugaredLogger {
	l, _ := zap.NewDevelopment()
	sugared := l.Sugar()
	return sugared
}
