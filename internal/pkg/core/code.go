package core

import (
	"net/http"
)

var (
	HTTP200 = &ErrCode{HTTPCode: http.StatusOK, Message: ""}
	HTTP400 = &ErrCode{HTTPCode: http.StatusBadRequest, Message: ""}
	HTTP401 = &ErrCode{HTTPCode: http.StatusUnauthorized, Message: ""}
	HTTP403 = &ErrCode{HTTPCode: http.StatusForbidden, Message: ""}
	HTTP404 = &ErrCode{HTTPCode: http.StatusNotFound, Message: ""}
	HTTP500 = &ErrCode{HTTPCode: http.StatusInternalServerError, Message: ""}

	ErrBind = &ErrCode{HTTPCode: http.StatusBadRequest, Message: "参数绑定错误"}

	InternalServerError = &ErrCode{HTTPCode: http.StatusInternalServerError, Message: "服务器内部错误"}
)
