package test

import (
	"context"
	"testing"

	"github.com/alf4ridzi/library-crud-ent-echo/internal/config"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/infrastructure/database"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/repository"
)

func TestFindByEmail(t *testing.T) {
	config.LoadEnv()

	db, err := database.NewMysqlEnt()
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	userRepo := repository.NewUserRepository(db)

	user, err := userRepo.FindByEmail(context.Background(), "IvUowLx@kOvltjR.biz")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(user)
}
