package repository

import (
	"context"
	"time"

	"github.com/alf4ridzi/library-crud-ent-echo/ent"
	"github.com/alf4ridzi/library-crud-ent-echo/ent/borrowings"
)

type BorrowRepository interface {
	FindByBookIDAndUserID(ctx context.Context, bookID uint, userID uint) (*ent.Borrowings, error)
	UpdateOneByID(ctx context.Context, id int, borrow *ent.Borrowings) error
	FindAllByBookIDBorrow(ctx context.Context, bookID uint) ([]*ent.Borrowings, error)
	FindOneByBookID(ctx context.Context, bookID uint) (*ent.Borrowings, error)
	Create(ctx context.Context, borrow *ent.Borrowings) (*ent.Borrowings, error)
}

type borrowRepositoryImpl struct {
	DB *ent.Client
}

func NewBorrowRepository(client *ent.Client) BorrowRepository {
	return &borrowRepositoryImpl{DB: client}
}

func (r *borrowRepositoryImpl) FindByBookIDAndUserID(ctx context.Context, bookID uint, userID uint) (*ent.Borrowings, error) {
	return r.DB.Borrowings.
		Query().
		Where(
			borrowings.UserID(userID),
			borrowings.BookID(bookID),
		).First(ctx)
}

func (r *borrowRepositoryImpl) UpdateOneByID(ctx context.Context, id int, borrow *ent.Borrowings) error {
	return r.DB.Borrowings.
		UpdateOneID(id).
		SetBookID(borrow.BookID).
		SetUserID(borrow.UserID).
		SetReleaseDate(*borrow.ReleaseDate).
		SetDueDate(borrow.DueDate).
		Exec(ctx)
}

func (r *borrowRepositoryImpl) FindAllByBookIDBorrow(ctx context.Context, bookID uint) ([]*ent.Borrowings, error) {
	return r.DB.Borrowings.
		Query().
		WithUser().
		WithBook().
		Where(
			borrowings.BookID(bookID),
			borrowings.HasUser(),
		).All(ctx)
}

func (r *borrowRepositoryImpl) FindOneByBookID(ctx context.Context, bookID uint) (*ent.Borrowings, error) {
	return r.DB.Borrowings.
		Query().
		WithUser().
		WithBook().
		Where(
			borrowings.BookID(bookID),
		).First(ctx)
}

func (r *borrowRepositoryImpl) UpdateReleaseDate(ctx context.Context, release time.Time) error {
	return nil
}

func (r *borrowRepositoryImpl) Create(ctx context.Context, borrow *ent.Borrowings) (*ent.Borrowings, error) {
	return r.DB.Borrowings.
		Create().
		SetBookID(borrow.BookID).
		SetUserID(borrow.UserID).
		SetDueDate(borrow.DueDate).
		Save(ctx)
}
