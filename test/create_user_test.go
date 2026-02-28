package test

import (
	"context"
	"testing"
	"time"

	"github.com/alf4ridzi/library-crud-ent-echo/ent"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/config"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/infrastructure/database"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/repository"
	"github.com/go-faker/faker/v4"
)

func TestCreateUser(t *testing.T) {
	config.LoadEnv()

	client, err := database.NewMysqlEnt()
	if err != nil {
		t.Fatal(err)
	}

	defer client.Close()

	repo := repository.NewUserRepository(client)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	name := "test name"

	user := ent.User{
		Name:     &name,
		Email:    faker.Email(),
		Username: faker.Password(),
		Password: "test123",
	}

	err = repo.Create(ctx, user)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(user)
}
