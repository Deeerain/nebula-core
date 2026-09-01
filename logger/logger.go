package logger

import (
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

type Logger interface {
	Info(msg string, args ...any)
	Debug(msg string, args ...any)
	Error(msg string, args ...any)
}

func Init(env string, minLevel slog.Level) {
	var handler slog.Handler

	opts := &slog.HandlerOptions{
		Level:     minLevel,
		AddSource: true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.SourceKey {
				if source, ok := a.Value.Any().(*slog.Source); ok {
					source.File = shortPath(source.File)
				}
			}
			return a
		},
	}

	if strings.ToLower(env) == "prod" || strings.ToLower(env) == "production" {
		handler = slog.NewTextHandler(os.Stdout, opts)
	} else {
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	slog.SetDefault(slog.New(handler))
}

func shortPath(path string) string {
	dir := filepath.Base(filepath.Dir(path))
	file := filepath.Base(path)
	if dir == "." || dir == "/" {
		return file
	}

	return filepath.Join(dir, file)
}
