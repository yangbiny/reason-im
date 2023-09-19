package error

import "github.com/pkg/errors"

func Error(error error) error {
	return errors.WithStack(error)
}
