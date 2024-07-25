package cmd

import (
	"fmt"

	"github.com/hoomy-official/go-shared/pkg/logger"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Commons defines the common flags and embedded commands for printing version
// and licence information, utilized by the command-line interface.
type Commons struct {
	Development bool   `short:"D" env:"DEBUG,DEV,DEVELOPMENT" help:"Set to true to enable development mode with debug-level logging."`            //nolint:lll // lll do not understand tags https://github.com/golangci/golangci-lint/issues/3983#issue-1833630800
	Level       string `short:"l" env:"LOG_LEVEL" help:"Specify the logging level, options are: debug, info, warn, error, fatal." default:"info"` //nolint:lll // same

	Version Version `cmd:"" help:"Display version information."`
	Licence Licence `cmd:"" help:"Show the application's licence."`
}

// Logger initializes a new zap.Logger based on the Development and Level fields in the commons struct.
// It returns the configured logger or an error if the logging level is invalid or the logger cannot be created.
func (g *Commons) Logger() (*zap.Logger, error) {
	level, err := zapcore.ParseLevel(g.Level)
	if err != nil {
		return nil, fmt.Errorf("cannot parse logger level \"%s\": %w", g.Level, err)
	}

	if g.Development {
		level = zapcore.DebugLevel
	}

	return logger.New(level, g.Development)
}
