package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(c *gin.Context) {

	c.JSON(http.StatusOK, nil)
}
