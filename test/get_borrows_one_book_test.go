package test

import (
	"context"
	"testing"

	"github.com/alf4ridzi/library-crud-ent-echo/internal/config"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/infrastructure/database"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/repository"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/service"
)

func TestGetBorrowOneBook(t *testing.T) {
	config.LoadEnv()

	db, err := database.NewMysqlEnt()
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	borrowRepo := repository.NewBorrowRepository(db)
	bookRepo := repository.NewBookRepository(db)

	bookService := service.NewBookService(bookRepo, borrowRepo)

	borrows, err := bookService.GetOneBookBorrows(context.Background(), "3")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(borrows)
}
