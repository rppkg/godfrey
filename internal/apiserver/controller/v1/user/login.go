package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	v1 "github.com/rppkg/godfrey/pkg/api/v1"
)

func (h *Handler) Login(c *gin.Context) {
	var r v1.LoginUserRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	resp, err := h.svc.Users().Login(c, &r)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, resp)
}
