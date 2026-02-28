package response

import "github.com/labstack/echo/v5"

func Success(c *echo.Context, data any) error {
	return c.JSON(200, SuccessResponse{
		Response: Response{Status: "success"},
		Data:     data,
	})
}
