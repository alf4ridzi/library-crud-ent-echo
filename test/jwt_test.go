package test

import (
	"testing"
	"time"

	"github.com/alf4ridzi/library-crud-ent-echo/internal/config"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/delivery/http/dto"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/pkg/tokenutil"
)

func TestCreateJwtToken(t *testing.T) {
	config.LoadEnv()

	accessToken, err := tokenutil.GenerateAccessToken("123", time.Duration(1)*time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	refreshToken, err := tokenutil.GenerateAccessToken("123", time.Duration(7)*time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	token := dto.AuthJwt{
		Token: dto.AuthJwtResponse{
			Access:  accessToken,
			Refresh: refreshToken,
		},
	}

	t.Log(token)

}
