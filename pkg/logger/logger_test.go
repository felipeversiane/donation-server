package logger

import (
	"context"
	"log/slog"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/felipeversiane/donation-server/config"
	"gopkg.in/natefinch/lumberjack.v2"
)

func TestLoggerInitialization(t *testing.T) {
	dir := t.TempDir()
	logPath := filepath.Join(dir, "app.log")
	cfg := config.Log{
		Level:      "error",
		Path:       logPath,
		Stdout:     false,
		MaxSize:    1,
		MaxBackups: 2,
		MaxAge:     3,
		Compress:   false,
	}

	l, err := New(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	l.Debug("debug message")
	l.Error("error message")

	data, err := os.ReadFile(logPath)
	if err != nil {
		t.Fatalf("failed to read log file: %v", err)
	}

	if len(data) == 0 {
		t.Fatalf("expected log file to contain data")
	}

	if string(data) == "" {
		t.Fatalf("logger did not write to file")
	}
	if !slogEnabled(l, slog.LevelError) {
		t.Fatalf("logger did not log at expected level")
	}

	lj := extractLumberjackWriter(t, l)
	expected := lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   cfg.Compress,
	}
	if !reflect.DeepEqual(*lj, expected) {
		t.Fatalf("lumberjack config mismatch")
	}
}

func slogEnabled(l Interface, level slog.Level) bool {
	return l.Handler().Enabled(context.TODO(), level)
}

func extractLumberjackWriter(t *testing.T, l Interface) *lumberjack.Logger {
	h := l.Handler()
	jh, ok := h.(*slog.JSONHandler)
	if !ok {
		t.Fatalf("unexpected handler type %T", h)
	}
	v := reflect.ValueOf(jh).Elem().FieldByName("commonHandler")
	v = reflect.Indirect(v)
	w := v.FieldByName("w").Interface()

	lj, ok := w.(*lumberjack.Logger)
	if !ok {
		t.Fatalf("writer is %T, want *lumberjack.Logger", w)
	}
	return lj
}
