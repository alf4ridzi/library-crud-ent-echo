package repository

import (
	"context"

	"github.com/alf4ridzi/library-crud-ent-echo/ent"
	"github.com/alf4ridzi/library-crud-ent-echo/ent/books"
)

type BookRepository interface {
	RemoveBorrowings(ctx context.Context, id uint, borrow *ent.Borrowings) error
	AddAvailableQuantity(ctx context.Context, id uint, qty int) error
	AddBorrowings(ctx context.Context, id uint, borrow *ent.Borrowings) error
	AddQuantity(ctx context.Context, id uint, qty int) error
	FindOneByID(ctx context.Context, id uint) (*ent.Books, error)
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

func (r *bookRepositoryImpl) RemoveBorrowings(ctx context.Context, id uint, borrow *ent.Borrowings) error {
	return r.DB.Books.UpdateOneID(id).
		RemoveBorrowings(borrow).
		Exec(ctx)
}

func (r *bookRepositoryImpl) AddBorrowings(ctx context.Context, id uint, borrow *ent.Borrowings) error {
	return r.DB.Books.
		UpdateOneID(id).
		AddBorrowings(borrow).
		Exec(ctx)
}

func (r *bookRepositoryImpl) AddAvailableQuantity(ctx context.Context, id uint, qty int) error {
	return r.DB.Books.
		UpdateOneID(id).
		AddAvailableQuantity(qty).
		Exec(ctx)
}

func (r *bookRepositoryImpl) AddQuantity(ctx context.Context, id uint, qty int) error {
	return r.DB.Books.
		UpdateOneID(id).
		AddQuantity(qty).
		Exec(ctx)
}

func (r *bookRepositoryImpl) FindOneByID(ctx context.Context, id uint) (*ent.Books, error) {
	return r.DB.Books.Query().
		WithBorrowings().
		WithCategories().
		Where(
			books.ID(id),
		).First(ctx)
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
