package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var ErrorLimitExceeded = errors.New("limit exceeded")

func Limit(maxEventsPerSec float64, maxBurstSize int) gin.HandlerFunc {
	limiter := rate.NewLimiter(rate.Limit(maxEventsPerSec), maxBurstSize)

	return func(c *gin.Context) {
		if limiter.Allow() {
			c.Next()

			return
		}

		_ = c.Error(ErrorLimitExceeded)
		c.AbortWithStatus(http.StatusTooManyRequests)
	}
}
