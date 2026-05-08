package handler

import (
	"errors"
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

func (h *BookHandler) GetBookCategories(c *echo.Context) error {
	categories, err := h.bookService.GetBookCategories(c.Request().Context())
	if err != nil {
		return response.Error(
			c,
			http.StatusInternalServerError,
			"internal server error",
		)
	}

	return response.Success(c, categories)
}

func (h *BookHandler) GetOneBookBorrow(c *echo.Context) error {
	id := c.Param("id")

	borrows, err := h.bookService.GetOneBookBorrows(c.Request().Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrBookNotFound):
			return response.Fail(
				c,
				http.StatusNotFound,
				response.Message(err.Error()),
			)
		default:
			c.Logger().Error(err.Error())
			return response.Error(
				c,
				http.StatusInternalServerError,
				"internal server error",
			)
		}

	}

	return response.Success(
		c,
		borrows,
	)
}

func (h *BookHandler) GetOneBook(c *echo.Context) error {
	id := c.Param("id")

	book, err := h.bookService.GetOneBook(c.Request().Context(), id)

	if err != nil {
		switch {
		case errors.Is(err, service.ErrBookNotFound):
			return response.Fail(
				c,
				http.StatusNotFound,
				response.Message(err.Error()),
			)
		default:
			c.Logger().Error(err.Error())
			return response.Error(
				c,
				http.StatusInternalServerError,
				"internal server error",
			)
		}

	}

	return response.Success(
		c,
		book,
	)
}

func (h *BookHandler) DeleteBook(c *echo.Context) error {
	id := c.Param("id")

	err := h.bookService.DeleteBook(c.Request().Context(), id)
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
		response.Message("success delete book"),
	)
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
