package user

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"

	"github.com/rppkg/godfrey/internal/pkg/core"
	v1 "github.com/rppkg/godfrey/pkg/api/v1"
)

func (h *Handler) Update(c *gin.Context) {
	var r v1.UpdateUserRequest
	if err := c.ShouldBindQuery(&r); err != nil {
		core.JSONResponse(c, core.ErrBindJSON, nil)
		return
	}

	if _, err := govalidator.ValidateStruct(r); err != nil {
		core.JSONResponse(c, core.ErrValidate.SetMessage("参数校验错误: %v", err.Error()), nil)
		return
	}

	resp, err := h.svc.Users().Update(c, c.Param("username"), &r)
	if err != nil {
		core.JSONResponse(c, err, nil)
		return
	}

	core.JSONResponse(c, nil, resp)
}
