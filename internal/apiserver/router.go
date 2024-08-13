package apiserver

import (
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"

	"github.com/rppkg/godfrey/internal/apiserver/controller/v1/user"
	"github.com/rppkg/godfrey/internal/apiserver/dal"
	"github.com/rppkg/godfrey/internal/pkg/auth"
	"github.com/rppkg/godfrey/internal/pkg/middleware"
)

func initRouters(g *gin.Engine) error {
	pprof.Register(g)

	g.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})

	g.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, "Page Not Found.")
	})

	authz, err := auth.NewAuthz(dal.Cli().DB())
	if err != nil {
		return err
	}

	userH := user.New(dal.Cli(), authz)

	g.POST("/login", userH.Login)

	api := g.Group("/api")

	initAPIUserRouters(api, userH, authz)

	return nil
}

func initAPIUserRouters(r *gin.RouterGroup, u *user.Handler, a *auth.Authz) {
	v1 := r.Group("/v1")
	{
		userv1 := v1.Group("/users")
		userv1.POST("", u.Regist)

		userv1.Use(middleware.Authn(), middleware.Authz(a))
		userv1.PUT(":username", u.Update)
		userv1.GET(":username", u.Get)
		userv1.GET("", u.List)
		userv1.DELETE(":username", u.Delete)
	}
}
