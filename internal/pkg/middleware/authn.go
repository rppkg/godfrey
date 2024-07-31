package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rppkg/godfrey/pkg/token"
)

func Authn() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 解析 JWT Token
		username, err := token.ParseRequest(c)
		if err != nil {
			c.JSON(http.StatusForbidden, nil)
			c.Abort()

			return
		}

		c.Set("X-Username", username)
		c.Next()
	}
}
