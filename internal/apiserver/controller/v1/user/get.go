package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Get(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
