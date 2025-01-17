package users

import "errors"

var (
	ErrUserDontExist    = errors.New("user doesn't exist")
	ErrUserAlreadyExist = errors.New("user already exist")
)
