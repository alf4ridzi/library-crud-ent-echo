package tokenutil

import "github.com/golang-jwt/jwt/v5"

type ClaimsAccessJWT struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}

type ClaimsRefreshJWT struct {
	jwt.RegisteredClaims
}
