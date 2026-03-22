package service

import (
	"context"

	"github.com/alf4ridzi/library-crud-ent-echo/ent"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/delivery/http/dto"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/repository"
)

type BookService interface {
	GetAllBooks(ctx context.Context) ([]dto.BookResponse, error)
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

func (s *bookServiceImpl) GetAllBooks(ctx context.Context) ([]dto.BookResponse, error) {
	queries, err := s.bookRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var books []dto.BookResponse

	for _, b := range queries {
		book := dto.BookResponse{
			Author:            b.Author,
			Description:       b.Description,
			Title:             b.Title,
			Quantity:          b.Quantity,
			AvailableQuantity: b.AvailableQuantity,
			PublishDate:       b.PublishDate,
			Categories:        b.Edges.Categories,
			Borrowings:        b.Edges.Borrowings,
			CreatedAt:         b.CreatedAt,
			UpdatedAt:         b.UpdatedAt,
		}

		books = append(books, book)
	}

	return books, nil
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
		Borrowings:        book.Edges.Borrowings,
		CreatedAt:         book.CreatedAt,
		UpdatedAt:         book.UpdatedAt,
	}

	return response, nil
}
