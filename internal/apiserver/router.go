package apiserver

import (
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"

	"github.com/rppkg/godfrey/internal/apiserver/controller/v1/user"
)

func initRouters(g *gin.Engine) error {
	pprof.Register(g)

	g.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})

	userH := user.NewHandler()

	g.POST("/login", userH.Login)

	api := g.Group("/api")

	initUserRouters(api, userH)

	return nil
}

func initUserRouters(r *gin.RouterGroup, u *user.Handler) {
	v1 := r.Group("/v1")
	{
		userv1 := v1.Group("/users")
		userv1.POST("", u.Create)
		userv1.PUT(":username", u.Update)
		userv1.GET(":username", u.Get)
		userv1.GET("", u.List)
		userv1.DELETE(":username", u.Delete)
	}
}
