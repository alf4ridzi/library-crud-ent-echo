package service

import "github.com/alf4ridzi/library-crud-ent-echo/internal/repository"

type BookService interface{}

type bookServiceImpl struct {
	bookRepo repository.BookRepository
}

func NewBookService(
	bookRepo repository.BookRepository,
) BookService {
	return &bookServiceImpl{bookRepo: bookRepo}
}
