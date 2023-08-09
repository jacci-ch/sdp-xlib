package errors

import "errors"

var (
	ErrNotImpl       = New("not implemented")
	ErrNoConfig      = New("configuration not found")
	ErrInvalidConfig = New("invalid configuration")
	ErrInvalidArgs   = New("invalid arguments")
)

func New(str string) error {
	return errors.New("grpcx: " + str)
}

func Wrapper(err error) error {
	return New(err.Error())
}
