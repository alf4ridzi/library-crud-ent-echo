package test

import (
	"testing"
	"time"

	"github.com/alf4ridzi/library-crud-ent-echo/internal/config"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/pkg/tokenutil"
)

func TestCreateJwtToken(t *testing.T) {
	config.LoadEnv()

	token, err := tokenutil.GenerateAuthToken("123", time.Duration(1)*time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(token)

}
