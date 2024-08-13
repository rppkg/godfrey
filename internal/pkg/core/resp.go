package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func JSONResponse(c *gin.Context, err error, data interface{}) {
	hcode, message := Decode(err)
	if hcode != http.StatusOK {
		c.JSON(hcode, BaseResponse{
			Code:    hcode,
			Message: message,
		})

		return
	}

	c.JSON(http.StatusOK, data)
}
