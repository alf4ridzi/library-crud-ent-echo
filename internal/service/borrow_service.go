package service

import (
	"context"
	"strconv"

	"github.com/alf4ridzi/library-crud-ent-echo/ent"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/delivery/http/dto"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/repository"
)

type BorrowService interface {
	ReleaseBorrow(ctx context.Context, bookIDStr string, req *dto.ReleaseBorrowRequest) error
	Borrow(ctx context.Context, bookID string, req *dto.BorrowRequest) error
}

type borrowServiceImpl struct {
	bookRepo   repository.BookRepository
	userRepo   repository.UserRepository
	borrowRepo repository.BorrowRepository
	DB         *ent.Client
}

func NewBorrowService(
	bookRepo repository.BookRepository,
	userRepo repository.UserRepository,
	borrowRepo repository.BorrowRepository,
	db *ent.Client,
) BorrowService {
	return &borrowServiceImpl{
		bookRepo:   bookRepo,
		userRepo:   userRepo,
		borrowRepo: borrowRepo,
		DB:         db}
}

func (s *borrowServiceImpl) ReleaseBorrow(ctx context.Context, bookIDStr string, req *dto.ReleaseBorrowRequest) error {
	bookID, err := strconv.ParseUint(bookIDStr, 10, 64)
	if err != nil {
		return err
	}

	book, err := s.bookRepo.FindOneByID(ctx, uint(bookID))
	if err != nil {
		if ent.IsNotFound(err) {
			return ErrBookNotFound
		}

		return err
	}

	user, err := s.bookRepo.FindOneByID(ctx, req.UserID)
	if err != nil {
		return err
	}

	var borrowToUpdate *ent.Borrowings

	for _, borrow := range book.Edges.Borrowings {
		if borrow.UserID == user.ID && borrow.BookID == book.ID {
			if borrow.ReleaseDate != nil {
				return ErrBorrowAlreadyRelease
			}

			borrow.ReleaseDate = &req.ReleaseDate

			borrowToUpdate = borrow

			break
		}
	}

	return s.borrowRepo.UpdateOneByID(ctx, borrowToUpdate.ID, borrowToUpdate)
}

func (s *borrowServiceImpl) Borrow(ctx context.Context, bookIDStr string, req *dto.BorrowRequest) error {
	bookID, err := strconv.ParseUint(bookIDStr, 10, 64)
	if err != nil {
		return err
	}

	tx, err := s.DB.Tx(ctx)
	if err != nil {
		return err
	}

	bookRepo := repository.NewBookRepository(tx.Client())
	userRepo := repository.NewUserRepository(tx.Client())
	borrowRepo := repository.NewBorrowRepository(tx.Client())

	// fmt.Println("executing find book")
	book, err := bookRepo.FindOneByID(ctx, uint(bookID))
	if err != nil {
		if ent.IsNotFound(err) {
			return ErrBookNotFound
		}

		return err
	}

	// fmt.Println("executing find user by id")
	user, err := userRepo.FindByID(ctx, req.UserID)
	if err != nil {
		return err
	}

	for _, borrows := range book.Edges.Borrowings {
		if user.ID == borrows.UserID {
			return ErrBookAlreadyBorrow
		}
	}

	borrowNew := &ent.Borrowings{
		BookID:  book.ID,
		UserID:  user.ID,
		DueDate: req.DueDate,
	}

	_, err = borrowRepo.Create(ctx, borrowNew)
	if err != nil {
		return err
	}

	// fmt.Println("executing add borrowings")
	// fmt.Println(borrow)
	// err = bookRepo.AddBorrowings(ctx, book.ID, borrow)
	// if err != nil {
	// 	return ent.Rollback(tx, err)
	// }

	// fmt.Println("executing add qty")
	err = bookRepo.AddAvailableQuantity(ctx, book.ID, -1)
	if err != nil {
		return ent.Rollback(tx, err)
	}

	tx.Commit()

	return nil
}
