package test

import (
	"context"
	"testing"
	"time"

	"github.com/alf4ridzi/library-crud-ent-echo/internal/config"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/delivery/http/dto"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/infrastructure/database"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/repository"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/service"
)

func TestChangePasswordUser(t *testing.T) {
	config.LoadEnv()

	client, err := database.NewMysqlEnt()
	if err != nil {
		t.Fatal(err)
	}

	defer client.Close()

	userRepo := repository.NewUserRepository(client)
	userService := service.NewUserService(userRepo)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	pw := &dto.UserChangePasswordRequest{
		OldPassword: "user12345",
		NewPassword: "alfaridzi123",
	}

	err = userService.ChangeUserPassword(ctx, "2", pw)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("ok")
}
