package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (controller *Handler) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
