package test

import (
	"context"
	"testing"

	"github.com/alf4ridzi/library-crud-ent-echo/ent"
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

	name := "test edit"

	user := &ent.User{
		ID:   1,
		Name: &name,
	}

	err = userRepo.UpdateByUser(context.Background(), user)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("ok")
}
