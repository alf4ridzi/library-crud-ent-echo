package test

import (
	"testing"

	"github.com/alf4ridzi/library-crud-ent-echo/internal/config"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/pkg/tokenutil"
)

func TestClaimsAccessToken(t *testing.T) {
	config.LoadEnv()

	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIyNyIsImV4cCI6MTc3MzE1NDQ2Nn0.4U0ncDed5Mq6qT3z4RZMKtFZKXuR9LtR1TQpNa97ifo"

	claims, err := tokenutil.ClaimsAccessToken(accessToken)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(claims)
}
