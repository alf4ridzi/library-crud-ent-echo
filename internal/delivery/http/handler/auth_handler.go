package handler

import (
	"net/http"

	"github.com/alf4ridzi/library-crud-ent-echo/internal/delivery/http/dto"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/delivery/http/response"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/service"
	"github.com/labstack/echo/v5"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(c *echo.Context) error {
	return c.JSON(200, map[string]string{
		"message": "ok",
	})
}

func (h *AuthHandler) Register(c *echo.Context) error {
	req := new(dto.RegisterRequest)

	if err := c.Bind(req); err != nil {
		return response.Fail(c, http.StatusBadRequest, err)
	}

	if err := c.Validate(req); err != nil {
		return response.Fail(c,
			http.StatusBadRequest,
			response.FormatValidationError(err, req),
		)
	}

	return nil
}

func (h *AuthHandler) Refresh(c *echo.Context) error {
	return nil
}
