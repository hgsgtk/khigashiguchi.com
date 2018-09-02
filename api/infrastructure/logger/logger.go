package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var writer zapcore.WriteSyncer = os.Stdout

// Error write error log.
func Error(msg string) {
	logger := newLogger(writer)
	defer logger.Sync()
	logger.Error(msg)
}

func newLogger(writer zapcore.WriteSyncer) *zap.Logger {
	atom := zap.NewAtomicLevel()

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = ""

	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.Lock(writer),
		atom,
	))
	return logger
}
