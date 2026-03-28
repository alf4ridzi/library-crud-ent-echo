package handler

import (
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

	return nil
}
