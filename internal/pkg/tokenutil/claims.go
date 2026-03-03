package tokenutil

import "github.com/golang-jwt/jwt/v5"

type ClaimsUserJWT struct {
	jwt.RegisteredClaims
}
