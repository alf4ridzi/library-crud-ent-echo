package response

import "github.com/labstack/echo/v5"

func Error(c *echo.Context, statusCode int, message string) error {
	return c.JSON(statusCode, ErrorResponse{
		Response: Response{Status: "error"},
		Message:  message,
	})
}
