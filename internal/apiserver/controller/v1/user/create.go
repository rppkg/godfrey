package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (controller *Controller) Create(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
