package routes

import "github.com/labstack/echo/v5"

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

	r.AuthRoute.Register(api)
}
