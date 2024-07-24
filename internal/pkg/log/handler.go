package log

import (
	"context"
	"io"
	"log"
	"log/slog"
)

type ctxKey string

const (
	godfreyFields ctxKey = "godfrey_fields"
)

type GodfreyHandlerOptions struct {
	SlogOpts slog.HandlerOptions
}

type GodfreyHandler struct {
	slog.Handler
	l *log.Logger
}

func (h *GodfreyHandler) Handle(ctx context.Context, r slog.Record) error {
	if attrs, ok := ctx.Value(godfreyFields).([]slog.Attr); ok {
		for _, v := range attrs {
			r.AddAttrs(v)
		}
	}

	fields := make(map[string]interface{}, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()
		return true
	})

	return h.Handler.Handle(ctx, r)
}

func NewGodfreyHandler(out io.Writer, opts GodfreyHandlerOptions) *GodfreyHandler {
	h := &GodfreyHandler{
		Handler: slog.NewJSONHandler(out, &opts.SlogOpts),
		l:       log.New(out, "", 0),
	}
	return h
}
