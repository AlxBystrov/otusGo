package logger

import (
	"log/slog"
	"os"
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

func New(level string, handler string) *Logger {
	levelInt, ok := levelMap[strings.ToLower(level)]
	if !ok {
		levelInt = 2 // use info as default level
	}
	var loggerHandler *slog.Logger
	switch handler {
	case "text":
		loggerHandler = slog.New(slog.NewTextHandler(os.Stdout, nil))
	case "json":
		loggerHandler = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	default:
		loggerHandler = slog.New(slog.NewTextHandler(os.Stdout, nil))
	}
	return &Logger{
		level:  levelInt,
		logger: *loggerHandler,
	}
}

func (l Logger) Debug(msg string, args ...any) {
	if l.level <= levelMap["debug"] {
		l.logger.Debug(msg, args...)
	}
}

func (l Logger) Info(msg string, args ...any) {
	if l.level <= levelMap["info"] {
		l.logger.Info(msg, args...)
	}
}

func (l Logger) Warning(msg string, args ...any) {
	if l.level <= levelMap["warning"] {
		l.logger.Warn(msg, args...)
	}
}

func (l Logger) Error(msg string, args ...any) {
	if l.level <= levelMap["error"] {
		l.logger.Error(msg, args...)
	}
}
