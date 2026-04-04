package service

import "errors"

var (
	ErrInvalidCredentials  = errors.New("username/email/password is incorrect")
	ErrInvalidPassword     = errors.New("password is incorrect")
	ErrBookNotFound        = errors.New("book not found")
	ErrBookAlreadyBorrow   = errors.New("you already borrowing this book")
	ErrUserIsNotBorrowBook = errors.New("user is not borrowing this book")
)
