package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New initializes and returns a new Zap logger with the specified level and development mode.
// The development flag determines whether to use a production or development configuration.
// Additional zap.Options can also be provided to customize logger behavior.
func New(level zapcore.Level, development bool, opts ...zap.Option) (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	if development {
		config = zap.NewDevelopmentConfig()
	}

	config.Level = zap.NewAtomicLevelAt(level)

	return config.Build(opts...)
}
