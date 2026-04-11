package handler

import "errors"

var (
	ErrBind     = errors.New("bind error")
	ErrValidate = errors.New("validate error")
)
