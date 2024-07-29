package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (controller *Controller) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
