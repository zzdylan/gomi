package logger

import (
	"testing"

	"go.uber.org/zap"
)

func TestLogger(t *testing.T) {
	// logger, _ := zap.NewProduction()
	logger, _ := zap.NewDevelopment()

	logger.Info("test")
	logger.Info("message", zap.String("key", "value"))
	// 性能稍低但更简洁的 SugaredLogger
	sugar := logger.Sugar()
	sugar.Infof("message with %s", "format")
	sugar.Infow("message with structured fields", "key", "value")
}
