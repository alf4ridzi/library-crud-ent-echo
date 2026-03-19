package handler

import (
	"net/http"

	"github.com/alf4ridzi/library-crud-ent-echo/ent"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/delivery/http/dto"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/delivery/http/response"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/helpers"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/service"
	"github.com/labstack/echo/v5"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) UpdateUser(c *echo.Context) error {
	req := new(dto.UserUpdateRequest)

	if err := c.Bind(req); err != nil {
		c.Logger().Error(err.Error())
		return response.Error(
			c,
			http.StatusInternalServerError,
			"internal server error",
		)
	}

	if err := c.Validate(req); err != nil {
		return response.Fail(
			c,
			http.StatusBadRequest,
			response.ValidationErrors(err),
		)
	}

	userID := c.Get("user_id")

	user, err := h.userService.UpdateUser(c.Request().Context(), userID, req)
	if err != nil {
		switch {
		case ent.IsConstraintError(err):
			return response.Fail(
				c,
				http.StatusConflict,
				helpers.ParseConstraintError(err),
			)
		default:
			c.Logger().Error(err.Error())
			return response.Error(
				c,
				http.StatusInternalServerError,
				"internal server error",
			)
		}
	}

	return response.Success(
		c,
		user,
	)
}

func (h *UserHandler) GetMe(c *echo.Context) error {
	userID := c.Get("user_id")

	user, err := h.userService.GetByID(c.Request().Context(), userID)
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
		user,
	)
}
