package response

import (
	"github.com/labstack/echo/v5"
)

func Fail(c *echo.Context, statusCode int, data any) error {
	return c.JSON(statusCode, FailResponse{
		Response: Response{Status: "fail"},
		Data:     data,
	})
}
