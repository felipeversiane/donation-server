package logger_test

import (
	"log/slog"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/felipeversiane/donation-server/config"
	loggerpkg "github.com/felipeversiane/donation-server/pkg/logger"
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

	l := loggerpkg.New(cfg)
	l.Logger().Debug("debug message")
	l.Logger().Error("error message")

	data, err := os.ReadFile(logPath)
	if err != nil {
		t.Fatalf("failed to read log file: %v", err)
	}

	if !reflect.DeepEqual(true, len(data) > 0) {
		t.Fatalf("expected log file to contain data")
	}

	if string(data) == "" || !slogEnabled(l.Logger(), slog.LevelError) {
		t.Fatalf("logger did not log at expected level")
	}

	lj := extractLumberjackWriter(t, l)
	if lj.Filename != logPath || lj.MaxSize != cfg.MaxSize || lj.MaxBackups != cfg.MaxBackups || lj.MaxAge != cfg.MaxAge || lj.Compress != cfg.Compress {
		t.Fatalf("lumberjack config mismatch")
	}
}

func slogEnabled(l *slog.Logger, level slog.Level) bool {
	return l.Handler().Enabled(nil, level)
}

func extractLumberjackWriter(t *testing.T, l loggerpkg.Interface) *lumberjack.Logger {
	h := l.Logger().Handler()
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
