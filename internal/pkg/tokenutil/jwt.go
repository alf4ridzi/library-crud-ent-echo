package tokenutil

import (
	"errors"
	"time"

	"github.com/alf4ridzi/library-crud-ent-echo/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken(userID string, expired time.Duration) (string, error) {
	claims := ClaimsAccessJWT{
		jwt.RegisteredClaims{
			Subject:   userID,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expired)),
			Issuer:    config.AppConfig.AppName,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString([]byte(config.AppConfig.JwtAccessSecret))
	if err != nil {
		return "", err
	}

	return signed, nil
}

func GenerateRefreshToken(userID string, expired time.Duration) (string, error) {
	claims := ClaimsRefreshJWT{
		jwt.RegisteredClaims{
			Subject:   userID,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expired)),
			Issuer:    config.AppConfig.AppName,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString([]byte(config.AppConfig.JwtRefreshSecret))
	if err != nil {
		return "", err
	}

	return signed, nil
}

func ClaimsAccessToken(tokenJwt string) (*ClaimsAccessJWT, error) {
	token, err := jwt.ParseWithClaims(tokenJwt, &ClaimsAccessJWT{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(config.AppConfig.JwtAccessSecret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*ClaimsAccessJWT)
	if !ok {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func ClaimsRefreshToken(tokenJwt string) (*ClaimsRefreshJWT, error) {
	token, err := jwt.ParseWithClaims(tokenJwt, &ClaimsRefreshJWT{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(config.AppConfig.JwtAccessSecret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*ClaimsRefreshJWT)
	if !ok {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
