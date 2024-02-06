package logger

import (
	"io"
	"log"
	"log/slog"
	"strings"
)

type Logger struct {
	level  int
	logger slog.Logger
}

var levelMap = map[string]int{
	"debug":   1,
	"info":    2,
	"warning": 3,
	"error":   4,
}

func New(w io.Writer, level string, handler string) *Logger {
	levelInt, ok := levelMap[strings.ToLower(level)]
	if !ok {
		log.Printf("failed to set level for %s value, used INFO as default\n", level)
		levelInt = 2 // use info as default level
	}

	var loggerHandler *slog.Logger
	switch handler {
	case "text":
		loggerHandler = slog.New(slog.NewTextHandler(w, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case "json":
		loggerHandler = slog.New(slog.NewJSONHandler(w, &slog.HandlerOptions{Level: slog.LevelDebug}))
	default:
		loggerHandler = slog.New(slog.NewTextHandler(w, &slog.HandlerOptions{Level: slog.LevelDebug}))
	}
	return &Logger{
		level:  levelInt,
		logger: *loggerHandler,
	}
}

func (l *Logger) Debug(msg string, args ...any) {
	if l.level <= levelMap["debug"] {
		l.logger.Debug(msg, args...)
	}
}

func (l *Logger) Info(msg string, args ...any) {
	if l.level <= levelMap["info"] {
		l.logger.Info(msg, args...)
	}
}

func (l *Logger) Warning(msg string, args ...any) {
	if l.level <= levelMap["warning"] {
		l.logger.Warn(msg, args...)
	}
}

func (l *Logger) Error(msg string, args ...any) {
	if l.level <= levelMap["error"] {
		l.logger.Error(msg, args...)
	}
}
