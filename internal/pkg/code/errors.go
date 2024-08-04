package code

import (
	"errors"
	"fmt"
)

type ErrCode struct {
	ICode    string
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

func (err ErrCode) Code() string {
	return err.ICode
}

func (err ErrCode) String() string {
	return err.Message
}

func (err ErrCode) HTTPStatus() int {
	if err.HTTPCode == 0 {
		return 500
	}

	return err.HTTPCode
}

func Decode(err error) (int, string, string) {
	if err == nil {
		return OK.HTTPCode, OK.ICode, OK.Message
	}

	var e *ErrCode
	if errors.As(err, &e) {
		return e.HTTPCode, e.ICode, e.Message
	}
	return InternalServerError.HTTPCode, InternalServerError.ICode, err.Error()
}
