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

func TestAddBorrow(t *testing.T) {
	config.LoadEnv()

	db, err := database.NewMysqlEnt()
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	bookRepo := repository.NewBookRepository(db)
	userRepo := repository.NewUserRepository(db)

	borrowService := service.NewBorrowService(bookRepo, userRepo, db)

	req := &dto.BorrowRequest{
		UserID:  3,
		DueDate: time.Now(),
	}

	err = borrowService.Borrow(context.Background(), "1", req)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("ok")
}
