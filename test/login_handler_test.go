package test

// import (
// 	"testing"

// 	"github.com/alf4ridzi/library-crud-ent-echo/internal/config"
// 	"github.com/alf4ridzi/library-crud-ent-echo/internal/delivery/http/handler"
// 	"github.com/alf4ridzi/library-crud-ent-echo/internal/infrastructure/database"
// 	"github.com/alf4ridzi/library-crud-ent-echo/internal/repository"
// 	"github.com/alf4ridzi/library-crud-ent-echo/internal/service"
// )

// func TestHandlerLogin(t *testing.T) {
// 	config.LoadEnv()

// 	db, err := database.NewMysqlEnt()
// 	if err != nil {
// 		t.Log(err)
// 	}

// 	defer db.Close()

// 	userRepo := repository.NewUserRepository(db)
// 	authService := service.NewAuthService(userRepo)
// 	authHandler := handler.NewAuthHandler(authService)

// }
