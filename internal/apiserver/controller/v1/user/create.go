package user

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"

	"github.com/rppkg/godfrey/internal/pkg/core"
	v1 "github.com/rppkg/godfrey/pkg/api/v1"
)

func (h *Handler) Create(c *gin.Context) {
	var r v1.CreateUserRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.JSONResponse(c, core.ErrBindJSON, nil)
		return
	}

	if _, err := govalidator.ValidateStruct(r); err != nil {
		core.JSONResponse(c, core.ErrValidate.SetMessage("参数校验错误: %v", err.Error()), nil)
		return
	}

	resp, err := h.svc.Users().Create(c, &r)
	if err != nil {
		core.JSONResponse(c, err, nil)
		return
	}

	_, err = h.authz.AddNamedPolicy("p", r.Username, "/v1/users/"+r.Username, "(GET)|(POST)|(PUT)|(DELETE)")
	if err != nil {
		core.JSONResponse(c, err, nil)
		return
	}

	core.JSONResponse(c, nil, resp)
}
