package handler

import (
	"github.com/labstack/echo/v5"
)

func bindAndValidateReq(c *echo.Context, i any) error {
	if err := c.Bind(i); err != nil {
		return ErrBind
	}

	if err := c.Validate(i); err != nil {
		return err
	}

	return nil
}
