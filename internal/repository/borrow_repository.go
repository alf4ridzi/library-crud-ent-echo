package repository

import (
	"context"
	"time"

	"github.com/alf4ridzi/library-crud-ent-echo/ent"
)

type BorrowRepository interface {
	Create(ctx context.Context, borrow *ent.Borrowings) (*ent.Borrowings, error)
}

type borrowRepositoryImpl struct {
	DB *ent.Client
}

func NewBorrowRepository(client *ent.Client) BorrowRepository {
	return &borrowRepositoryImpl{DB: client}
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
