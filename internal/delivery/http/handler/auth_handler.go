package handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/alf4ridzi/library-crud-ent-echo/ent"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/delivery/http/dto"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/delivery/http/response"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/helpers"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/pkg/tokenutil"
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

	token, err := h.authService.Login(c.Request().Context(), req)
	if err != nil {
		c.Logger().Error(err.Error())
		switch {
		case errors.Is(err, service.ErrInvalidCredentials):
			return response.Fail(
				c,
				http.StatusUnauthorized,
				response.Message(err.Error()),
			)
		case ent.IsNotFound(err):
			return response.Fail(
				c,
				http.StatusUnauthorized,
				response.Message(service.ErrInvalidCredentials.Error()),
			)
		default:
			return response.Error(
				c,
				http.StatusInternalServerError,
				"internal server error",
			)
		}
	}

	return response.Success(
		c,
		token,
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
	req := new(dto.RefreshRequest)

	if err := c.Bind(req); err != nil {
		return response.Error(
			c,
			http.StatusInternalServerError,
			err.Error(),
		)
	}

	if err := c.Validate(req); err != nil {
		return response.Fail(
			c,
			http.StatusBadRequest,
			response.ValidationErrors(err),
		)
	}

	claims, err := tokenutil.ClaimsRefreshToken(req.Token)
	if err != nil {
		c.Logger().Error(err.Error())
		return response.Error(
			c,
			http.StatusInternalServerError,
			"internal server error",
		)
	}

	accessToken, err := tokenutil.GenerateAccessToken(claims.Subject, time.Duration(1)*time.Hour)
	if err != nil {
		c.Logger().Error(err.Error())
		return response.Error(
			c,
			http.StatusInternalServerError,
			"internal server error",
		)
	}

	return response.Success(
		c,
		dto.AuthJwt{
			Token: dto.AuthJwtResponse{
				Access: accessToken,
			},
		},
	)
}
