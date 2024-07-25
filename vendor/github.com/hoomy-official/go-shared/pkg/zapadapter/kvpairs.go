// Package zapadapter provides a wrapper around the Zap logging library,
// allowing to adapt a *zap.Logger to a Logger interface that works with key/value pairs.
package zapadapter

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// logFunc is a private type that adapts a standard logging function
// to the Logger interface.
type logFunc func(kv ...interface{}) error

// Log calls the underlying logging function of logFunc.
func (f logFunc) Log(kv ...interface{}) error {
	return f(kv...)
}

// Logger is the fundamental interface for all log operations.
// Log creates a log event from keyvals, a variadic sequence of
// alternating keys and values. Implementations must be safe for
// concurrent use by multiple goroutines.
type Logger interface {
	Log(keyvals ...interface{}) error
}

// ZapAdapter adapts a *zap.Logger to the Logger interface, using
// the service name as a constant key.
func ZapAdapter(service string, logger *zap.Logger) Logger {
	var l func(msg string, keysAndValues ...interface{})

	switch logger.Level() {
	case zapcore.DebugLevel:
		l = logger.Sugar().Debugw
	case zapcore.InfoLevel:
		l = logger.Sugar().Infow
	case zapcore.WarnLevel:
		l = logger.Sugar().Warnw
	case zapcore.ErrorLevel:
		l = logger.Sugar().Errorw
	case zapcore.DPanicLevel:
		l = logger.Sugar().DPanicw
	case zapcore.PanicLevel:
		l = logger.Sugar().Panicw
	case zapcore.FatalLevel:
		l = logger.Sugar().Fatalw
	case zapcore.InvalidLevel:
		l = logger.Sugar().Debugw
	default:
		l = logger.Sugar().Debugw
	}

	return logFunc(func(kv ...interface{}) error {
		l(service, kv...)
		return nil
	})
}
