package service

import (
	"context"

	"github.com/alf4ridzi/library-crud-ent-echo/ent"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/delivery/http/dto"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/repository"
)

type BookService interface {
	CreateNewBook(ctx context.Context, req *dto.CreateNewBookRequest) (*dto.BookResponse, error)
}

type bookServiceImpl struct {
	bookRepo repository.BookRepository
}

func NewBookService(
	bookRepo repository.BookRepository,
) BookService {
	return &bookServiceImpl{bookRepo: bookRepo}
}

func (s *bookServiceImpl) CreateNewBook(ctx context.Context,
	req *dto.CreateNewBookRequest) (*dto.BookResponse, error) {

	newBook := &ent.Books{
		Author:            req.Author,
		Description:       req.Description,
		Title:             req.Author,
		Quantity:          req.Quantity,
		AvailableQuantity: req.AvailableQuantity,
		PublishDate:       req.PublishDate,
	}

	book, err := s.bookRepo.Create(ctx, newBook, req.CategoryIDs)
	if err != nil {
		return nil, err
	}

	response := &dto.BookResponse{
		Author:            book.Author,
		Description:       book.Description,
		Title:             book.Title,
		Quantity:          book.Quantity,
		AvailableQuantity: book.AvailableQuantity,
		PublishDate:       book.PublishDate,
		Categories:        book.Edges.Categories,
	}

	return response, nil
}
