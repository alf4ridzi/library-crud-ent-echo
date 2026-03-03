package middleware

import (
	"net/http"
	"strings"

	"github.com/alf4ridzi/library-crud-ent-echo/internal/delivery/http/response"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/pkg/tokenutil"
	"github.com/labstack/echo/v5"
)

func JwtAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		authorization := c.Request().Header.Get("Authorization")

		if authorization == "" {
			return response.Fail(
				c,
				http.StatusUnauthorized,
				"missing authorization header",
			)
		}

		authSplit := strings.Split(authorization, "Bearer")

		token := authSplit[1]

		claims, err := tokenutil.ClaimsAccessToken(token)
		if err != nil {
			return response.Fail(
				c,
				http.StatusUnauthorized,
				err.Error(),
			)
		}

		c.Set("user_id", claims.Subject)

		return next(c)
	}
}
