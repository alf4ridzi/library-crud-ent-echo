package handler

import (
	"errors"
	"net/http"

	"github.com/alf4ridzi/library-crud-ent-echo/ent"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/delivery/http/dto"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/delivery/http/response"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/helpers"
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
	req := new(dto.LoginRequest)

	if err := c.Bind(req); err != nil {
		return response.Fail(
			c,
			http.StatusBadRequest,
			err,
		)
	}

	if err := c.Validate(req); err != nil {
		return response.Fail(
			c,
			http.StatusBadRequest,
			response.ValidationErrors(err),
		)
	}

	user, err := h.authService.Login(c.Request().Context(), req)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidCredentials):
			return response.Fail(
				c,
				http.StatusUnauthorized,
				response.Message(err.Error()),
			)
		}
	}

	return response.Success(
		c,
		user,
	)
}

func (h *AuthHandler) Register(c *echo.Context) error {
	req := new(dto.RegisterRequest)

	if err := c.Bind(req); err != nil {
		return response.Fail(c, http.StatusBadRequest, err)
	}

	if err := c.Validate(req); err != nil {
		return response.Fail(c,
			http.StatusBadRequest,
			response.ValidationErrors(err),
		)
	}

	err := h.authService.Register(c.Request().Context(), req)
	if err != nil {
		switch {
		case ent.IsConstraintError(err):
			return response.Fail(c, http.StatusConflict, helpers.ParseConstraintError(err))
		default:
			return response.Error(c, http.StatusInternalServerError, err.Error())
		}

	}

	return response.Success(c, response.MessageResponse{
		Message: "user registered successfully",
	})
}

func (h *AuthHandler) Refresh(c *echo.Context) error {
	return nil
}
