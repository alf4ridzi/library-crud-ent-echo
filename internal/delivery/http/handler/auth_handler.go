package handler

import (
	"github.com/alf4ridzi/library-crud-ent-echo/internal/service"
	"github.com/labstack/echo/v5"
)

type AuthHandler struct {
	authService service.AuthService
}

func (h *AuthHandler) Login(c *echo.Context) error {
	return nil
}

func (h *AuthHandler) Register(c *echo.Context) error {
	return nil
}

func (h *AuthHandler) Refresh(c *echo.Context) error {
	return nil
}
