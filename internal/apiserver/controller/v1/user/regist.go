package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Regist(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
