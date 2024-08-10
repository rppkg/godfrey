package user

import (
	"github.com/rppkg/godfrey/internal/apiserver/dal"
	"github.com/rppkg/godfrey/internal/apiserver/service"
	"github.com/rppkg/godfrey/internal/pkg/auth"
)

type Handler struct {
	authz *auth.Authz
	svc   service.IService
}

func New(dal dal.IDal, authz *auth.Authz) *Handler {
	return &Handler{
		authz: authz,
		svc:   service.New(dal),
	}
}
