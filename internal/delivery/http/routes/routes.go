package routes

import (
	"github.com/alf4ridzi/library-crud-ent-echo/internal/delivery/http/middleware"
	"github.com/labstack/echo/v5"
)

type Routes struct {
	AuthRoute *AuthRoute
}

func NewRoutes(
	authRoute *AuthRoute,
) *Routes {
	return &Routes{
		AuthRoute: authRoute,
	}
}

func (r *Routes) Register(router *echo.Echo) {
	api := router.Group("/api")

	api.Use(middleware.JwtAuth)

	r.AuthRoute.Register(api)
}
