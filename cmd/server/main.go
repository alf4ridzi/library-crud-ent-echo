package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/alf4ridzi/library-crud-ent-echo/ent/runtime"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/config"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/delivery/http/handler"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/delivery/http/middleware"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/delivery/http/routes"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/infrastructure/database"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/repository"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/service"
	"github.com/labstack/echo/v5"
	echoMiddleware "github.com/labstack/echo/v5/middleware"
)

func startServer(e *echo.Echo) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	s := http.Server{
		Addr:    fmt.Sprintf(":%d", config.AppConfig.AppPort),
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
	e.Use(echoMiddleware.ContextTimeout(60 * time.Second))
	e.Use(middleware.TimeoutMiddleware)

	userRepo := repository.NewUserRepository(client)

	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)
	authRoute := routes.NewAuthRoute(authHandler)

	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	userRoute := routes.NewUserRoute(userHandler)

	bookRepo := repository.NewBookRepository(client)
	bookService := service.NewBookService(bookRepo)
	bookHandler := handler.NewBookHandler(bookService)
	bookRoute := routes.NewBookRoute(bookHandler)

	r := routes.NewRoutes(authRoute, userRoute, bookRoute)

	r.Register(e)

	startServer(e)
}
