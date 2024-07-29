package middleware

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"

	"github.com/rppkg/godfrey/pkg/log"
)

func SlogInPrint() gin.HandlerFunc {
	return sloggin.NewWithConfig(log.RowL(), sloggin.Config{
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

		Filters: []sloggin.Filter{
			sloggin.IgnoreStatus(http.StatusUnauthorized, http.StatusNotFound, http.StatusOK),
		},
	})
}
