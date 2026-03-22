package handler

import (
	"net/http"

	"github.com/alf4ridzi/library-crud-ent-echo/internal/delivery/http/dto"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/delivery/http/response"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/service"
	"github.com/labstack/echo/v5"
)

type BookHandler struct {
	bookService service.BookService
}

func NewBookHandler(bookService service.BookService) *BookHandler {
	return &BookHandler{
		bookService: bookService,
	}
}

func (h *BookHandler) GetAllBooks(c *echo.Context) error {
	books, err := h.bookService.GetAllBooks(c.Request().Context())
	if err != nil {
		c.Logger().Error(err.Error())
		return response.Error(
			c,
			http.StatusInternalServerError,
			"internal server error",
		)
	}

	return response.Success(
		c,
		books,
	)
}

func (h *BookHandler) Store(c *echo.Context) error {
	req := new(dto.CreateNewBookRequest)

	if err := c.Bind(req); err != nil {
		c.Logger().Error(err.Error())
		return response.Error(
			c,
			http.StatusInternalServerError,
			"internal server error",
		)
	}

	if err := c.Validate(req); err != nil {
		return response.Fail(
			c,
			http.StatusBadRequest,
			response.ValidationErrors(err),
		)
	}

	book, err := h.bookService.CreateNewBook(c.Request().Context(), req)
	if err != nil {
		c.Logger().Error(err.Error())
		return response.Error(
			c,
			http.StatusInternalServerError,
			"internal server error",
		)
	}

	return response.Success(
		c,
		book,
	)
}
