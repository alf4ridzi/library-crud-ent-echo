package routes

import (
	"github.com/alf4ridzi/library-crud-ent-echo/internal/delivery/http/handler"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/delivery/http/middleware"
	"github.com/labstack/echo/v5"
)

type BookRoute struct {
	Handler *handler.BookHandler
}

func NewBookRoute(handler *handler.BookHandler) *BookRoute {
	return &BookRoute{
		Handler: handler,
	}
}

func (r *BookRoute) Register(rg *echo.Group) {
	books := rg.Group("/books")
	books.GET("", r.Handler.GetAllBooks)
	books.Use(middleware.JwtAuth)
	books.POST("", r.Handler.Store)
	books.DELETE("/books/:id", r.Handler.DeleteBook)
}
