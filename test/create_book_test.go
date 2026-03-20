package test

import (
	"context"
	"testing"
	"time"

	"github.com/alf4ridzi/library-crud-ent-echo/ent"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/config"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/infrastructure/database"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/repository"
)

func TestCreateBook(t *testing.T) {
	config.LoadEnv()

	db, err := database.NewMysqlEnt()
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	bookRepo := repository.NewBookRepository(db)

	bookCategories := []uint{
		1,
		2,
	}

	newBook := &ent.Books{
		Author:            "admin test",
		Description:       "ini adalah buku tentang satria baja hitam",
		Title:             "satria baja hitam",
		Quantity:          10,
		AvailableQuantity: 10,
		PublishDate:       time.Now(),
	}

	book, err := bookRepo.Create(context.Background(), newBook, bookCategories)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(book)
}
