package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alf4ridzi/library-crud-ent-echo/internal/config"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/delivery/http/handler"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/delivery/http/middleware"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/delivery/http/routes"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/infrastructure/database"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/repository"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/service"
	"github.com/labstack/echo/v5"
)

func startServer(e *echo.Echo) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	s := http.Server{
		Addr:    ":8080",
		Handler: e,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			e.Logger.Error("failed to start server", "error", err)
		}
	}()

	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		e.Logger.Error("failed to stop server", "error", err)
	}
}

func main() {
	e := echo.New()

	if err := config.LoadEnv(); err != nil {
		log.Fatalf("failed to load env : %v", err)
	}

	client, err := database.NewMysqlEnt()
	if err != nil {
		log.Fatalf("failed to connect to database : %v", err)
	}

	defer client.Close()

	e.Validator = middleware.NewValidator()

	userRepo := repository.NewUserRepository(client)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)
	authRoute := routes.NewAuthRoute(authHandler)

	r := routes.NewRoutes(authRoute)
	r.Register(e)

	startServer(e)
}
