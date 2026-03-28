package routes

import (
	"github.com/alf4ridzi/library-crud-ent-echo/internal/delivery/http/handler"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/delivery/http/middleware"
	"github.com/labstack/echo/v5"
)

type BorrowRoute struct {
	Handler *handler.BorrowHandler
}

func NewBorrowRoute(handler *handler.BorrowHandler) *BorrowRoute {
	return &BorrowRoute{
		Handler: handler,
	}
}

func (r *BorrowRoute) Register(rg *echo.Group) {
	borrow := rg.Group("/borrows")
	borrow.Use(middleware.JwtAuth)
	borrow.POST("/:id", r.Handler.Borrow)
}
