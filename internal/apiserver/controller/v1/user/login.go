package user

import (
	"github.com/gin-gonic/gin"

	"github.com/rppkg/godfrey/internal/pkg/core"
	v1 "github.com/rppkg/godfrey/pkg/api/v1"
)

func (h *Handler) Login(c *gin.Context) {
	var r v1.LoginUserRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.JSONResponse(c, core.ErrBind, nil)
		return
	}

	resp, err := h.svc.Users().Login(c, &r)
	if err != nil {
		core.JSONResponse(c, err, nil)
		return
	}

	core.JSONResponse(c, nil, resp)
}
