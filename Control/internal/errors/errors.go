package errors

import "errors"

var (
	ErrUserExists       = errors.New("user already exists")
	ErrUnregisteredUser = errors.New("unregistered user")
	ErrInvalidPassword  = errors.New("invalid password")

	ErrTeamExists = errors.New("team already exists")
)
