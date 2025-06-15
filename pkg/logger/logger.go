package logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/felipeversiane/donation-server/config"
	"github.com/felipeversiane/donation-server/pkg/contextkey"
	"gopkg.in/natefinch/lumberjack.v2"
)

type logger struct {
	slog *slog.Logger
}

// Interface exposes common logging capabilities in a form similar to slog.Logger
// while also allowing log enrichment with request scoped values.
type Interface interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
	With(args ...any) Interface
	WithGroup(name string) Interface
	Handler() slog.Handler
	WithContext(ctx context.Context) Interface

	// Logger returns the underlying slog.Logger for advanced use cases.
	Logger() *slog.Logger
}

func New(config config.Log) Interface {
	var level slog.Level
	switch config.Level {
	case "debug":
		level = slog.LevelDebug
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     level,
	}

	dir := filepath.Dir(config.Path)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		panic(fmt.Sprintf("failed to create log directory: %v", err))
	}

	fileWriter := &lumberjack.Logger{
		Filename:   config.Path,
		MaxSize:    config.MaxSize,
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,
		Compress:   config.Compress,
	}

	var output io.Writer
	if config.Stdout {
		output = io.MultiWriter(os.Stdout, fileWriter)
	} else {
		output = fileWriter
	}

	var handler slog.Handler

	handler = slog.NewJSONHandler(output, opts)

	s := slog.New(handler)

	return &logger{slog: s}
}

func (l *logger) WithContext(ctx context.Context) Interface {
	if ctx == nil {
		return l
	}

	fields := []any{}
	if rid, ok := ctx.Value(contextkey.RequestID).(string); ok {
		fields = append(fields, "request_id", rid)
	}
	if uid, ok := ctx.Value(contextkey.UserID).(string); ok {
		fields = append(fields, "user_id", uid)
	}

	return &logger{slog: l.slog.With(fields...)}
}

func (l *logger) Logger() *slog.Logger {
	return l.slog
}

func (l *logger) Debug(msg string, args ...any) {
	l.slog.Debug(msg, args...)
}

func (l *logger) Info(msg string, args ...any) {
	l.slog.Info(msg, args...)
}

func (l *logger) Warn(msg string, args ...any) {
	l.slog.Warn(msg, args...)
}

func (l *logger) Error(msg string, args ...any) {
	l.slog.Error(msg, args...)
}

func (l *logger) With(args ...any) Interface {
	return &logger{slog: l.slog.With(args...)}
}

func (l *logger) WithGroup(name string) Interface {
	return &logger{slog: l.slog.WithGroup(name)}
}

func (l *logger) Handler() slog.Handler {
	return l.slog.Handler()
}
