package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Auther 用来定义授权接口实现.
// sub: 操作主题，obj：操作对象, act：操作
type Auther interface {
	Authorize(sub, obj, act string) (bool, error)
}

// Authz 是 Gin 中间件，用来进行请求授权.
func Authz(a Auther) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 可能这里使用用户的角色，基于角色控制权限会更好。查是否满足sub，act, obj的策略
		sub := c.GetString("X-Username")
		obj := c.Request.URL.Path
		act := c.Request.Method

		if allowed, _ := a.Authorize(sub, obj, act); !allowed {
			c.JSON(http.StatusForbidden, nil)

			c.Abort()
			return
		}
	}
}
