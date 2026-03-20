package handler

import "github.com/alf4ridzi/library-crud-ent-echo/internal/service"

type BookHandler struct {
	bookService service.BookService
}

func NewBookHandler(bookService service.BookService) *BookHandler {
	return &BookHandler{
		bookService: bookService,
	}
}
