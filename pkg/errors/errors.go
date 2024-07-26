package errors

import "github.com/pkg/errors"

func Wrap(e error, msg string) error {
	return errors.Wrap(e, msg)
}
