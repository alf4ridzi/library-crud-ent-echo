package test

import (
	"context"
	"testing"

	"github.com/alf4ridzi/library-crud-ent-echo/internal/config"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/infrastructure/database"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/repository"
)

func TestDeleteBook(t *testing.T) {
	config.LoadEnv()

	db, err := database.NewMysqlEnt()
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	bookRepo := repository.NewBookRepository(db)

	err = bookRepo.DeleteByID(context.Background(), 1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("ok")
}
