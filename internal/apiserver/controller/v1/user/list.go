package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (controller *Controller) List(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
