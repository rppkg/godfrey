package user

import (
	"github.com/gin-gonic/gin"

	"github.com/rppkg/godfrey/internal/pkg/core"
)

func (h *Handler) Get(c *gin.Context) {
	resp, err := h.svc.Users().Get(c, c.Param("username"))
	if err != nil {
		core.JSONResponse(c, err, nil)
		return
	}

	core.JSONResponse(c, nil, resp)
}
