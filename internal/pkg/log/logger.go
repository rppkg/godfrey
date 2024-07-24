package log

import (
	"context"
	"log/slog"
	"os"
	"sync"
)

type Logger struct {
	logger *slog.Logger
}

var (
	ins  *Logger
	once sync.Once
)

func L() *Logger {
	if ins != nil {
		return ins
	}

	once.Do(initLogger)

	return ins
}

func RowL() *slog.Logger {
	L()

	return ins.logger
}

func initLogger() {
	opts := GodfreyHandlerOptions{
		SlogOpts: slog.HandlerOptions{
			Level:       slog.LevelInfo,
			ReplaceAttr: replaceAttr,
		},
	}

	h := NewGodfreyHandler(os.Stdout, opts)
	logger := slog.New(h)

	l := &Logger{
		logger: logger,
	}

	ins = l
}

func (l *Logger) AppendCtx(parent context.Context, attr slog.Attr) context.Context {
	if parent == nil {
		parent = context.Background()
	}
	if v, ok := parent.Value(godfreyFields).([]slog.Attr); ok {
		v = append(v, attr)
		return context.WithValue(parent, godfreyFields, v)
	}
	var v []slog.Attr
	v = append(v, attr)
	return context.WithValue(parent, godfreyFields, v)
}

func (l *Logger) Debug(msg string, args ...any) {
	l.logger.Debug(msg, args...)
}

func (l *Logger) DebugContext(ctx context.Context, msg string, args ...any) {
	l.logger.DebugContext(ctx, msg, args...)
}

func (l *Logger) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}

func (l *Logger) InfoContext(ctx context.Context, msg string, args ...any) {
	l.logger.InfoContext(ctx, msg, args...)
}

func (l *Logger) Warn(msg string, args ...interface{}) {
	l.logger.Warn(msg, args...)
}

func (l *Logger) WarnContext(ctx context.Context, msg string, args ...any) {
	l.logger.WarnContext(ctx, msg, args...)
}

func (l *Logger) Error(msg string, args ...interface{}) {
	l.logger.Error(msg, args...)
}

func (l *Logger) ErrorContext(ctx context.Context, msg string, args ...any) {
	l.logger.ErrorContext(ctx, msg, args...)
}

func (l *Logger) Fatal(msg string, args ...interface{}) {
	l.logger.Error(msg, args...)
	os.Exit(1)
}

func (l *Logger) FatalContext(ctx context.Context, msg string, args ...any) {
	l.logger.ErrorContext(ctx, msg, args...)
	os.Exit(1)
}
