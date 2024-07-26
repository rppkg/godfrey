package godfrey

import (
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"

	"github.com/rppkg/godfrey/internal/godfrey/controller/v1/user"
	"github.com/rppkg/godfrey/pkg/log"
)

func initRouters(g *gin.Engine) error {
	pprof.Register(g)

	g.GET("/healthz", func(c *gin.Context) {
		log.Info("Healthz function called")

		c.JSON(http.StatusOK, "ok")
	})

	userCtrl := user.NewUserController()

	api := g.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			userv1 := v1.Group("/users")
			userv1.POST("", userCtrl.Create)
			userv1.PUT(":id", userCtrl.Update)
			userv1.GET(":id", userCtrl.Get)
			userv1.GET("", userCtrl.List)
			userv1.DELETE(":id", userCtrl.Delete)
		}
	}

	return nil
}
