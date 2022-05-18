package logger

import "go.uber.org/zap"

var Sugar *zap.SugaredLogger

func NewLogger() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	Sugar = logger.Sugar()
}
