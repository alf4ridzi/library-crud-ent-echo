package routes

import (
	"github.com/alf4ridzi/library-crud-ent-echo/delivery/http/handler"
	"github.com/labstack/echo/v5"
)

type AuthRoute struct {
	Handler *handler.AuthHandler
}

func NewAuthRoute(handler *handler.AuthHandler) *AuthRoute {
	return &AuthRoute{
		Handler: handler,
	}
}

func (r *AuthRoute) Register(rg *echo.Group) {
	auth := rg.Group("/auth")
	auth.POST("/login", r.Handler.Login)
	auth.POST("/register", r.Handler.Register)
	auth.POST("/refresh", r.Handler.Refresh)
}
