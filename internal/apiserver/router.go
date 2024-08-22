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
	// NOTE: 性能分析，一般可以通过基准测试，代码生成，http服务生成性能数据分拣，再分析这些性能数据文件
	// 通过host:port/debug/pprof 打开profiles
	// 通过curl host:port/debug/pprof/profile -o cpu.profile 获取cpu性能30s数据文件
	// 通过curl host:port/debug/pprof/heap -o mem.profile 获取mem性能数据文件
	// 通过go tool pprof [mem|cpu].profile 分析http服务cpu和mem的性能
	pprof.Register(g)

	g.GET("/healthz", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	g.NoRoute(func(c *gin.Context) {
		c.Status(http.StatusNotFound)
	})

	authz, err := auth.NewAuthz(dal.Cli().DB())
	if err != nil {
		return err
	}

	userH := user.NewHandle(dal.Cli(), authz)

	g.POST("/login", userH.Login)

	api := g.Group("/api")

	initAPIUserRouters(api, userH, authz)

	return nil
}

func initAPIUserRouters(r *gin.RouterGroup, u user.IHandler, a *auth.Authz) {
	v1 := r.Group("/v1")
	{
		userv1 := v1.Group("/users")

		userv1.POST("", u.Create)
		userv1.Use(middleware.Authn(), middleware.Authz(a))
		userv1.GET(":username", u.Get)
		userv1.PUT(":username", u.Update)
		userv1.GET("", u.List)
		userv1.DELETE(":username", u.Delete)
	}
}
