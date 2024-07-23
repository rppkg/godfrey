package log

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"log/slog"

	"github.com/fatih/color"
)

type ctxKey string

const (
	godgreyFields ctxKey = "godfrey_fields"
)

type GodfreyHandlerOptions struct {
	SlogOpts slog.HandlerOptions
}

type GodfreyHandler struct {
	slog.Handler
	l *log.Logger
}

func (h *GodfreyHandler) Handle(ctx context.Context, r slog.Record) error {
	level := r.Level.String() + ":"
	switch r.Level {
	case slog.LevelDebug:
		level = color.MagentaString(level)
	case slog.LevelInfo:
		level = color.BlueString(level)
	case slog.LevelWarn:
		level = color.YellowString(level)
	case slog.LevelError:
		level = color.RedString(level)
	}

	if attrs, ok := ctx.Value(godgreyFields).([]slog.Attr); ok {
		for _, v := range attrs {
			r.AddAttrs(v)
		}
	}

	fields := make(map[string]interface{}, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()
		return true
	})

	b, err := json.MarshalIndent(fields, "", "  ")
	if err != nil {
		return err
	}

	timeStr := r.Time.Format("[15:05:05.000]")
	msg := color.CyanString(r.Message)
	h.l.Println(timeStr, level, msg, color.WhiteString(string(b)))

	return h.Handler.Handle(ctx, r)
}

func NewPrettyHandler(out io.Writer, opts GodfreyHandlerOptions) *GodfreyHandler {
	h := &GodfreyHandler{
		Handler: slog.NewJSONHandler(out, &opts.SlogOpts),
		l:       log.New(out, "", 0),
	}
	return h
}
