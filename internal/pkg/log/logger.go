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

func init() {
	if ins != nil {
		return
	}
	once.Do(initLogger)
}

func L() *Logger {
	return ins
}

func RowL() *slog.Logger {
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

func Debug(msg string, args ...any) {
	ins.logger.Debug(msg, args...)
}

func DebugContext(ctx context.Context, msg string, args ...any) {
	ins.logger.DebugContext(ctx, msg, args...)
}

func Info(msg string, args ...any) {
	ins.logger.Info(msg, args...)
}

func InfoContext(ctx context.Context, msg string, args ...any) {
	ins.logger.InfoContext(ctx, msg, args...)
}

func Warn(msg string, args ...interface{}) {
	ins.logger.Warn(msg, args...)
}

func WarnContext(ctx context.Context, msg string, args ...any) {
	ins.logger.WarnContext(ctx, msg, args...)
}

func Error(msg string, args ...interface{}) {
	ins.logger.Error(msg, args...)
}

func ErrorContext(ctx context.Context, msg string, args ...any) {
	ins.logger.ErrorContext(ctx, msg, args...)
}

func Fatal(msg string, args ...interface{}) {
	ins.logger.Error(msg, args...)
	os.Exit(1)
}

func FatalContext(ctx context.Context, msg string, args ...any) {
	ins.logger.ErrorContext(ctx, msg, args...)
	os.Exit(1)
}
