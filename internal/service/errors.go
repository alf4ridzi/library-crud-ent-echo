package service

import "errors"

var (
	ErrInvalidCredentials = errors.New("username/email/password is wrong")
)
