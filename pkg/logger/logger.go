package logger

import (
	"context"
	"io"
	"log/slog"
	"os"
	"sync"
)

var (
	once         sync.Once
	instance     *slog.Logger
	mu           sync.RWMutex
	outputWriter io.Writer // Stores the output writer for Sync
)

// Config holds logger configuration
type Config struct {
	Level     slog.Level
	AddSource bool
	Output    *os.File
}

// DefaultConfig returns a default configuration
func DefaultConfig() *Config {
	return &Config{
		Level:     slog.LevelInfo,
		AddSource: false,
		Output:    os.Stdout,
	}
}

func Init(cfg *Config) {
	once.Do(func() {
		handlerOpts := &slog.HandlerOptions{
			Level:     cfg.Level,
			AddSource: cfg.AddSource,
		}

		mu.Lock()
		defer mu.Unlock()

		outputWriter = cfg.Output
		instance = slog.New(slog.NewJSONHandler(outputWriter, handlerOpts))
		slog.SetDefault(instance)
	})
}

// getLogger returns the initialized logger instance
func getLogger() *slog.Logger {
	mu.RLock()
	defer mu.RUnlock()

	if instance == nil {
		Init(DefaultConfig())
	}
	return instance
}

// Logging methods with proper level checks and key/value validation

func Error(msg string, keyvals ...interface{}) {
	getLogger().Error(msg, validateKeyvals(keyvals...)...)
}

func Info(msg string, keyvals ...interface{}) {
	getLogger().Info(msg, validateKeyvals(keyvals...)...)
}

func Debug(msg string, keyvals ...interface{}) {
	if getLogger().Enabled(context.Background(), slog.LevelDebug) {
		getLogger().Debug(msg, validateKeyvals(keyvals...)...)
	}
}

func Warn(msg string, keyvals ...interface{}) {
	getLogger().Warn(msg, validateKeyvals(keyvals...)...)
}

// ErrorContext Context-aware logging
func ErrorContext(ctx context.Context, msg string, keyvals ...interface{}) {
	getLogger().ErrorContext(ctx, msg, validateKeyvals(keyvals...)...)
}

func InfoContext(ctx context.Context, msg string, keyvals ...interface{}) {
	getLogger().InfoContext(ctx, msg, validateKeyvals(keyvals...)...)
}

// Helper function to validate key/value pairs
func validateKeyvals(keyvals ...interface{}) []interface{} {
	if len(keyvals)%2 != 0 {
		extra := []interface{}{"!INVALID", "odd number of keyvals"}
		return append(keyvals, extra...)
	}
	return keyvals
}

// Sync flushes any buffered log entries
func Sync() error {
	mu.RLock()
	defer mu.RUnlock()

	if outputWriter == nil {
		return nil
	}

	if syncer, ok := outputWriter.(interface{ Sync() error }); ok {
		return syncer.Sync()
	}
	return nil
}
