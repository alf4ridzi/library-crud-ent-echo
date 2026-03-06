package test

import (
	"context"
	"testing"

	"github.com/alf4ridzi/library-crud-ent-echo/internal/config"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/delivery/http/dto"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/infrastructure/database"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/repository"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/service"
)

func TestServiceLogin(t *testing.T) {
	config.LoadEnv()

	db, err := database.NewMysqlEnt()
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)

	req := new(dto.LoginRequest)

	req.Identifier = "GMZyTUj@ZYkNirQ.net"
	req.Password = "user12345"

	user, err := authService.Login(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(user)
}
