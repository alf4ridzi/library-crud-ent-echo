package test

import (
	"testing"

	"github.com/alf4ridzi/library-crud-ent-echo/internal/config"
)

func TestEnv(t *testing.T) {
	if err := config.LoadEnv(); err != nil {
		t.Fatal(err)
	}

	t.Log(config.AppConfig.DBHost)
}
