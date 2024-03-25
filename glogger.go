package glogger

import "log/slog"

type GLogger struct {
	test string
}

func New(test string) *GLogger {
	return &GLogger{test: test}
}

func (l *GLogger) Print() {
	slog.Info("hello from:", "v", l.test)
}
