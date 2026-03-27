package service

import (
	"context"
	"strconv"

	"github.com/alf4ridzi/library-crud-ent-echo/ent"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/delivery/http/dto"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/repository"
)

type BookService interface {
	GetOneBook(ctx context.Context, id string) (*dto.BookResponse, error)
	DeleteBook(ctx context.Context, id string) error
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

func (s *bookServiceImpl) GetOneBook(ctx context.Context, id string) (*dto.BookResponse, error) {
	bookID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}

	bookQuery, err := s.bookRepo.FindOneByID(ctx, uint(bookID))
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrBookNotFound
		}

		return nil, err
	}

	book := &dto.BookResponse{
		Author:            bookQuery.Author,
		Description:       bookQuery.Description,
		Title:             bookQuery.Title,
		Quantity:          bookQuery.Quantity,
		AvailableQuantity: bookQuery.AvailableQuantity,
		PublishDate:       bookQuery.PublishDate,
		Categories:        bookQuery.Edges.Categories,
		Borrowings:        bookQuery.Edges.Borrowings,
		CreatedAt:         bookQuery.CreatedAt,
		UpdatedAt:         bookQuery.UpdatedAt,
	}

	return book, nil
}

func (s *bookServiceImpl) DeleteBook(ctx context.Context, id string) error {
	bookID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}

	return s.bookRepo.DeleteByID(ctx, uint(bookID))
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
