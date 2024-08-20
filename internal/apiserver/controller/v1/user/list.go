package user

import (
	"github.com/gin-gonic/gin"

	"github.com/rppkg/godfrey/internal/pkg/core"
	v1 "github.com/rppkg/godfrey/pkg/api/v1"
)

func (h *Handler) List(c *gin.Context) {
	var r v1.ListUserRequest
	if err := c.ShouldBindQuery(&r); err != nil {
		core.JSONResponse(c, core.ErrBindJSON, nil)
		return
	}

	resp, err := h.svc.Users().List(c, &r)
	if err != nil {
		core.JSONResponse(c, err, nil)
		return
	}

	core.JSONResponse(c, nil, resp)
}
