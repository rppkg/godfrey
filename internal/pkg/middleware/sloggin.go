package middleware

import (
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
)

func SlogInPrint() gin.HandlerFunc {
	return sloggin.NewWithConfig(slog.New(slog.NewJSONHandler(os.Stdout, nil)), sloggin.Config{
		DefaultLevel:     slog.LevelInfo,
		ClientErrorLevel: slog.LevelWarn,
		ServerErrorLevel: slog.LevelError,

		WithUserAgent:     true,
		WithRequestID:     true,
		WithRequestBody:   true,
		WithRequestHeader: true,

		WithResponseBody:   true,
		WithResponseHeader: true,

		WithSpanID:  true,
		WithTraceID: true,
	})
}

func SlogInFilter() gin.HandlerFunc {
	return sloggin.NewWithFilters(
		slog.New(slog.NewJSONHandler(os.Stdout, nil)),
		sloggin.Accept(func(c *gin.Context) bool {
			return true
		}),
		sloggin.IgnoreStatus(401, 404),
	)
}
