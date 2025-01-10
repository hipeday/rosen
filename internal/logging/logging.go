package logging

import (
	"context"
	"github.com/hipeday/rosen/conf"
	ctx2 "github.com/hipeday/rosen/internal/ctx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

var logger *zap.SugaredLogger

func Setup(cfg *conf.Logging) {
	createLogger(cfg)
}

func Logger() *zap.SugaredLogger {
	if logger == nil {
		log.Fatal("logger not initialized")
	}
	return logger
}

func LoggerWithRequestID(ctx context.Context) *zap.SugaredLogger {
	if logger == nil {
		log.Fatal("logger not initialized")
	}
	if requestId, ok := ctx2.GetRequestId(ctx); ok {
		return logger.With("requestId", requestId)
	}
	return logger
}

func createLogger(c *conf.Logging) *zap.SugaredLogger {
	if logger != nil {
		return logger
	}
	if c == nil {
		log.Fatal("logger not initialized")
	}

	productionConfig := zap.NewProductionConfig()
	productionConfig.EncoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	timeFormat := "2006-01-02T15:04:05.000Z0700"

	if c.TimeFormat != "" {
		timeFormat = c.TimeFormat
	}
	// 设置时间格式
	productionConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(timeFormat)

	if c.Encoding == "" || c.Encoding == "console" {
		productionConfig.Encoding = "console"
	}

	if c.Colors != nil && *c.Colors {
		productionConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	switch c.Level {
	case "debug":
		productionConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		productionConfig.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		productionConfig.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		productionConfig.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	}

	build, _ := productionConfig.Build()
	logger = build.Sugar()
	return logger
}
