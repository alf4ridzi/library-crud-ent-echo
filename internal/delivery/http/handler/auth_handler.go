package handler

import (
	"net/http"
	"time"

	"github.com/alf4ridzi/library-crud-ent-echo/ent"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/delivery/http/dto"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/delivery/http/response"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/helpers"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/pkg/ctxutil"
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
			response.ValidationErrors(err),
		)
	}

	ctx, cancel := ctxutil.WithTimeout(c.Request().Context(), 2*time.Second)
	defer cancel()

	err := h.authService.Register(ctx, req)
	if err != nil {
		switch {
		case ent.IsConstraintError(err):
			return response.Fail(c, http.StatusConflict, helpers.ParseConstraintError(err))
		default:
			return response.Error(c, http.StatusInternalServerError, err.Error())
		}

	}

	return response.Success(c, req)
}

func (h *AuthHandler) Refresh(c *echo.Context) error {
	return nil
}
