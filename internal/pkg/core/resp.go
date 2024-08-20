package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	HTTP200 = &GodfreyErr{HTTPCode: http.StatusOK, Message: "成功"}
	HTTP400 = &GodfreyErr{HTTPCode: http.StatusBadRequest, Message: "参数错误"}
	HTTP401 = &GodfreyErr{HTTPCode: http.StatusUnauthorized, Message: "认证未通过"}
	HTTP403 = &GodfreyErr{HTTPCode: http.StatusForbidden, Message: "权限不够"}
	HTTP404 = &GodfreyErr{HTTPCode: http.StatusNotFound, Message: "页面未找到"}
	HTTP500 = &GodfreyErr{HTTPCode: http.StatusInternalServerError, Message: "服务器内部错误"}

	ErrBindJSON = &GodfreyErr{HTTPCode: http.StatusBadRequest, Message: "参数绑定错误"}
	ErrValidate = &GodfreyErr{HTTPCode: http.StatusBadRequest, Message: "参数校验错误"}

	InternalServerError = &GodfreyErr{HTTPCode: http.StatusInternalServerError, Message: "服务器内部错误"}
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
