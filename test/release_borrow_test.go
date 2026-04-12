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

func TestReleaseBorrow(t *testing.T) {
	config.LoadEnv()

	db, err := database.NewMysqlEnt()
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	bookRepo := repository.NewBookRepository(db)
	borrowRepo := repository.NewBorrowRepository(db)
	userRepo := repository.NewUserRepository(db)

	bs := service.NewBorrowService(bookRepo, userRepo, borrowRepo, db)

	req := &dto.ReleaseBorrowRequest{
		UserID:      2,
		ReleaseDate: time.Now(),
	}

	err = bs.ReleaseBorrow(context.Background(), "1", req)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("ok")
}
