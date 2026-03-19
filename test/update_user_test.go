package test

import (
	"context"
	"testing"

	"github.com/alf4ridzi/library-crud-ent-echo/internal/config"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/infrastructure/database"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/repository"
)

func TestUpdateUser(t *testing.T) {
	config.LoadEnv()

	db, err := database.NewMysqlEnt()
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	userRepo := repository.NewUserRepository(db)

	user, err := userRepo.FindByID(context.Background(), 1)

	user.Email = "test@gmail.com"
	user.Username = "YMWrSEB"

	err = userRepo.UpdateByUser(context.Background(), user)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("ok")
}
