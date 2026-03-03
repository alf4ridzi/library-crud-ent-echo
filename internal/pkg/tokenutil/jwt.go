package tokenutil

import (
	"time"

	"github.com/alf4ridzi/library-crud-ent-echo/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateAuthToken(userID string, expired time.Duration) (string, error) {
	claims := ClaimsUserJWT{
		jwt.RegisteredClaims{
			Subject:   userID,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expired)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString([]byte(config.AppConfig.JwtAuthSecret))
	if err != nil {
		return "", err
	}

	return signed, nil
}
