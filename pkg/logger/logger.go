package logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/felipeversiane/donation-server/config"
	"gopkg.in/natefinch/lumberjack.v2"
)

type logger struct {
	slog *slog.Logger
}

type Interface interface {
	Logger() *slog.Logger
	WithContext(ctx context.Context) *slog.Logger
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
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     28,
		Compress:   true,
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

func (l *logger) WithContext(ctx context.Context) *slog.Logger {
	if ctx == nil {
		return l.slog
	}

	fields := []any{}
	if rid, ok := ctx.Value("request_id").(string); ok {
		fields = append(fields, "request_id", rid)
	}
	if uid, ok := ctx.Value("user_id").(string); ok {
		fields = append(fields, "user_id", uid)
	}

	return l.slog.With(fields...)
}

func (l *logger) Logger() *slog.Logger {
	return l.slog
}
