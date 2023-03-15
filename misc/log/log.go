// Package log contains a set of functions and structs for creating and working with zap logger
package log

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewZap creates a new instance of zap logger with the provided level and encoding type.
// lvl: level of the logger (debug, info, warn, error, etc)
// encType: encoding type of the logger (json, console, etc)
// Returns: new instance of zap logger and error if any
func NewZap(lvl, encType string) (*zap.Logger, error) {
	var (
		err    error
		logger *zap.Logger
	)

	config := zap.NewProductionConfig()
	config.Encoding = encType
	config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.EncodeDuration = zapcore.StringDurationEncoder
	config.EncoderConfig.LevelKey = "level"
	config.EncoderConfig.CallerKey = "c"
	config.EncoderConfig.StacktraceKey = "s"

	if err = config.Level.UnmarshalText([]byte(lvl)); err != nil && lvl != "" {
		return nil, err
	}

	if logger, err = config.Build(); err != nil {
		return nil, err
	}

	return logger, nil
}

// ctxKey is a struct that is used as the key for storing logger in the context
type ctxKey struct{} // or exported to use outside the package

// CtxWithLogger adds the provided logger to the context and returns the new context
func CtxWithLogger(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, ctxKey{}, logger)
}

// FromContext retrieves the logger from the context, returns a Nop logger if not found
func FromContext(ctx context.Context) *zap.Logger {
	if ctxLogger, ok := ctx.Value(ctxKey{}).(*zap.Logger); ok {
		return ctxLogger
	}
	return zap.NewNop()
}
