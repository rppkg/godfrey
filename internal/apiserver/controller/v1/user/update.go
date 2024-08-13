package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Update(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
