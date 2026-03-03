package routes

import (
	"github.com/alf4ridzi/library-crud-ent-echo/internal/delivery/http/middleware"
	"github.com/labstack/echo/v5"
)

type Routes struct {
	AuthRoute *AuthRoute
	UserRoute *UserRoute
}

func NewRoutes(
	authRoute *AuthRoute,
	userRoute *UserRoute,
) *Routes {
	return &Routes{
		AuthRoute: authRoute,
		UserRoute: userRoute,
	}
}

func (r *Routes) Register(router *echo.Echo) {
	api := router.Group("/api")

	r.AuthRoute.Register(api)
	api.Use(middleware.JwtAuth)
	r.UserRoute.Register(api)
}
