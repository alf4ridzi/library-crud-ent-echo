package response

import "github.com/labstack/echo/v5"

func Error(c *echo.Context, statusCode int, data any) error {
	return c.JSON(statusCode, data)
}
