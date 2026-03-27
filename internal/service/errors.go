package service

import "errors"

var (
	ErrInvalidCredentials = errors.New("username/email/password is incorrect")
	ErrInvalidPassword    = errors.New("password is incorrect")
	ErrBookNotFound       = errors.New("book not found")
)
