package routes

import (
	"github.com/alf4ridzi/library-crud-ent-echo/internal/delivery/http/middleware"
	"github.com/labstack/echo/v5"
)

type Routes struct {
	AuthRoute *AuthRoute
	UserRoute *UserRoute
	BookRoute *BookRoute
}

func NewRoutes(
	authRoute *AuthRoute,
	userRoute *UserRoute,
	bookRoute *BookRoute,
) *Routes {
	return &Routes{
		AuthRoute: authRoute,
		UserRoute: userRoute,
		BookRoute: bookRoute,
	}
}

func (r *Routes) Register(router *echo.Echo) {
	api := router.Group("/api")

	api.Use(middleware.TimeoutMiddleware)
	r.AuthRoute.Register(api)
	r.UserRoute.Register(api)
	r.BookRoute.Register(api)
}
