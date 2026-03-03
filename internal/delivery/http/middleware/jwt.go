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
		authHeader := c.Request().Header.Get("Authorization")

		if authHeader == "" {
			return response.Fail(
				c,
				http.StatusUnauthorized,
				"missing authorization header",
			)
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := tokenutil.ClaimsAccessToken(tokenString)
		if err != nil {
			return response.Fail(
				c,
				http.StatusUnauthorized,
				response.Message(err.Error()),
			)
		}

		c.Set("user_id", claims.Subject)

		return next(c)
	}
}
