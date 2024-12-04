package myerrors

import "errors"

var (
	ErrInternal = errors.New("internal server error, please try again later")
)
