package handler

import (
	"errors"
	"net/http"

	"github.com/alf4ridzi/library-crud-ent-echo/internal/delivery/http/dto"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/delivery/http/response"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/service"
	"github.com/labstack/echo/v5"
)

type BorrowHandler struct {
	bs service.BorrowService
}

func NewBorrowHandler(borrowService service.BorrowService) *BorrowHandler {
	return &BorrowHandler{bs: borrowService}
}

func (h *BorrowHandler) ReleaseBorrow(c *echo.Context) error {
	return nil
}

func (h *BorrowHandler) Borrow(c *echo.Context) error {
	req := new(dto.BorrowRequest)

	if err := c.Bind(req); err != nil {
		c.Logger().Error(err.Error())
		return response.Error(
			c,
			http.StatusInternalServerError,
			"internals server error",
		)
	}

	if err := c.Validate(req); err != nil {
		return response.Fail(
			c,
			http.StatusBadRequest,
			response.ValidationErrors(err),
		)
	}

	bookID := c.Param("id")

	err := h.bs.Borrow(c.Request().Context(), bookID, req)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrBookAlreadyBorrow):
			return response.Fail(
				c,
				http.StatusConflict,
				err.Error(),
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
		response.Message("borrow book success"),
	)
}
