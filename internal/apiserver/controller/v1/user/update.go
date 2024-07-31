package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (controller *Handler) Update(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
