package tokenutil

import "github.com/golang-jwt/jwt/v5"

type ClaimsAccessJWT struct {
	jwt.RegisteredClaims
}

type ClaimsRefreshJWT struct {
	jwt.RegisteredClaims
}
