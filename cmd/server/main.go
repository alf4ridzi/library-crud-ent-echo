package main

import (
	"log"

	"github.com/alf4ridzi/library-crud-ent-echo/internal/config"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/delivery/http/handler"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/delivery/http/middleware"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/delivery/http/routes"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/infrastructure/database"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/repository"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/service"
	"github.com/labstack/echo/v5"
)

func main() {
	e := echo.New()

	if err := config.LoadEnv(); err != nil {
		log.Fatalf("failed to load env : %v", err)
	}

	client, err := database.NewMysqlEnt()
	if err != nil {
		log.Fatalf("failed to connect to database : %v", err)
	}

	e.Validator = middleware.NewValidator()

	userRepo := repository.NewUserRepository(client)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)
	authRoute := routes.NewAuthRoute(authHandler)

	r := routes.NewRoutes(authRoute)
	r.Register(e)

	if err := e.Start(":8080"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
