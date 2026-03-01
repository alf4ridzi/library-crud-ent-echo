package middleware

import (
	"context"
	"time"

	"github.com/labstack/echo/v5"
)

func TimeoutMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		ctx, cancel := context.WithTimeout(c.Request().Context(), 2*time.Second)
		defer cancel()

		req := c.Request().WithContext(ctx)
		c.SetRequest(req)

		return next(c)
	}
}
