package repository

import (
	"context"

	"github.com/alf4ridzi/library-crud-ent-echo/ent"
)

type BookRepository interface {
}

type bookRepositoryImpl struct {
	DB *ent.Client
}

func NewBookRepository(client *ent.Client) BookRepository {
	return &bookRepositoryImpl{DB: client}
}

func (r *bookRepositoryImpl) Create(ctx context.Context, book *ent.Books) error {
	return r.DB.Books.Create().
		SetAuthor(book.Author).
		SetDescription(book.Description).
		SetTitle(book.Title).
		SetQuantity(book.Quantity).
		SetAvailableQuantity(book.AvailableQuantity).
		SetPublishDate(book.PublishDate).
		Exec(ctx)
}
