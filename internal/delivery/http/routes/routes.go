package routes

import (
	internalMiddleware "github.com/alf4ridzi/library-crud-ent-echo/internal/delivery/http/middleware"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type Routes struct {
	AuthRoute   *AuthRoute
	UserRoute   *UserRoute
	BookRoute   *BookRoute
	borrowRoute *BorrowRoute
}

func NewRoutes(
	authRoute *AuthRoute,
	userRoute *UserRoute,
	bookRoute *BookRoute,
	borrowRoute *BorrowRoute,
) *Routes {
	return &Routes{
		AuthRoute:   authRoute,
		UserRoute:   userRoute,
		BookRoute:   bookRoute,
		borrowRoute: borrowRoute,
	}
}

func (r *Routes) Register(router *echo.Echo) {
	api := router.Group("/api")

	api.Use(internalMiddleware.TimeoutMiddleware)
	api.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(5.0)))

	r.AuthRoute.Register(api)
	r.UserRoute.Register(api)
	r.BookRoute.Register(api)
	r.borrowRoute.Register(api)
}
