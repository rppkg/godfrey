package core

import (
	"errors"
	"fmt"
)

type ErrCode struct {
	HTTPCode int
	Message  string
}

func (err ErrCode) Error() string {
	return err.Message
}

func (err ErrCode) SetMessage(format string, args ...interface{}) ErrCode {
	err.Message = fmt.Sprintf(format, args...)
	return err
}

func (err ErrCode) String() string {
	return err.Message
}

func (err ErrCode) HCode() int {
	if err.HTTPCode == 0 {
		return 500
	}

	return err.HTTPCode
}

func Decode(err error) (hcode int, message string) {
	if err == nil {
		return HTTP200.HTTPCode, HTTP200.Message
	}

	var e *ErrCode
	if errors.As(err, &e) {
		return e.HTTPCode, e.Message
	}
	return InternalServerError.HTTPCode, err.Error()
}
