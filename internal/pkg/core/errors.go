package core

import (
	"errors"
	"fmt"
)

type GodfreyErr struct {
	HTTPCode int
	Message  string
}

func (err GodfreyErr) Error() string {
	return err.Message
}

func (err GodfreyErr) SetMessage(format string, args ...interface{}) GodfreyErr {
	err.Message = fmt.Sprintf(format, args...)
	return err
}

func (err GodfreyErr) String() string {
	return err.Message
}

func (err GodfreyErr) HCode() int {
	if err.HTTPCode == 0 {
		return 500
	}

	return err.HTTPCode
}

func Decode(err error) (hcode int, message string) {
	if err == nil {
		return HTTP200.HTTPCode, HTTP200.Message
	}

	var e *GodfreyErr
	if errors.As(err, &e) {
		return e.HTTPCode, e.Message
	}
	return InternalServerError.HTTPCode, err.Error()
}
