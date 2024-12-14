package errors

import "errors"

var (
	ErrUnknownUser     = errors.New("unregistered user")
	ErrInvalidPassword = errors.New("invalid password")
)
