package user

import (
	"github.com/gin-gonic/gin"

	"github.com/rppkg/godfrey/internal/apiserver/dal"
	"github.com/rppkg/godfrey/internal/apiserver/service"
	"github.com/rppkg/godfrey/internal/pkg/auth"
)

type Handler struct {
	authz *auth.Authz
	svc   service.IService
}

type IHandler interface {
	Create(c *gin.Context)
	Login(c *gin.Context)
	Get(c *gin.Context)
	Update(c *gin.Context)
	List(c *gin.Context)
	Delete(c *gin.Context)
}

var _ IHandler = (*Handler)(nil)

func NewHandle(dal dal.IDal, authz *auth.Authz) IHandler {
	return &Handler{
		authz: authz,
		svc:   service.New(dal),
	}
}
