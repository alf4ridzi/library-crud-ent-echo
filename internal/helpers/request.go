package helpers

import (
	"github.com/labstack/echo/v5"
)

func BindAndValidateReq(c *echo.Context, i any) error {
	if err := c.Bind(i); err != nil {
		return err
	}

	if err := c.Validate(i); err != nil {
		return err
	}

	return nil
}
