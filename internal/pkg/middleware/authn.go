package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/rppkg/godfrey/pkg/token"
)

func Authn() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)
		if len(auth) != 2 || auth[0] != "Bearer" {
			c.Status(http.StatusUnauthorized)
			c.Abort()

			return
		}

		username, err := token.Parse(auth[1])
		if err != nil || len(username) == 0 {
			c.Status(http.StatusUnauthorized)
			c.Abort()

			return
		}

		c.Set("X-Username", username)
		c.Next()
	}
}
