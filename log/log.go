package log

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger() *zap.Logger {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger, err := cfg.Build()
	if err != nil {
		log.Fatal(err)
	}
	defer func(logger *zap.Logger) {
		if err := logger.Sync(); err != nil {
			log.Fatalln("service down", err.Error())
		}
	}(logger)
	logger.Info("Log initialized")

	return logger
}
