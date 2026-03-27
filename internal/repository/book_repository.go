package repository

import (
	"context"

	"github.com/alf4ridzi/library-crud-ent-echo/ent"
)

type BookRepository interface {
	DeleteByID(ctx context.Context, id uint) error
	FindAll(ctx context.Context) ([]*ent.Books, error)
	Create(ctx context.Context, book *ent.Books, categoryIDs []uint) (*ent.Books, error)
}

type bookRepositoryImpl struct {
	DB *ent.Client
}

func NewBookRepository(client *ent.Client) BookRepository {
	return &bookRepositoryImpl{DB: client}
}

func (r *bookRepositoryImpl) DeleteByID(ctx context.Context, id uint) error {
	return r.DB.Books.DeleteOneID(id).Exec(ctx)
}

func (r *bookRepositoryImpl) FindAll(ctx context.Context) ([]*ent.Books, error) {
	return r.DB.Books.Query().
		WithBorrowings().
		WithCategories().
		All(ctx)
}

func (r *bookRepositoryImpl) Create(ctx context.Context, book *ent.Books, categoryIDs []uint) (*ent.Books, error) {
	return r.DB.Books.Create().
		SetAuthor(book.Author).
		SetDescription(book.Description).
		SetTitle(book.Title).
		SetQuantity(book.Quantity).
		SetAvailableQuantity(book.AvailableQuantity).
		SetPublishDate(book.PublishDate).
		AddCategoryIDs(categoryIDs...).
		Save(ctx)
}
