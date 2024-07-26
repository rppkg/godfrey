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

	userCtrl := user.NewController()

	api := g.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			userv1 := v1.Group("/users")
			userv1.POST("", userCtrl.Create)
			userv1.PUT(":username", userCtrl.Update)
			userv1.GET(":username", userCtrl.Get)
			userv1.GET("", userCtrl.List)
			userv1.DELETE(":username", userCtrl.Delete)
		}
	}

	return nil
}
