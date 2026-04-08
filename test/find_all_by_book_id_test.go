package test

import (
	"context"
	"testing"

	"github.com/alf4ridzi/library-crud-ent-echo/internal/config"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/delivery/http/dto"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/infrastructure/database"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/repository"
)

func TestFindAllByBookIDBorrows(t *testing.T) {
	config.LoadEnv()

	db, err := database.NewMysqlEnt()
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	borrowRepo := repository.NewBorrowRepository(db)

	borrowsQuery, err := borrowRepo.FindAllByBookIDBorrow(context.Background(), 3)
	if err != nil {
		t.Fatal(err)
	}

	var borrows []dto.BorrowResponse

	for _, borrowQuery := range borrowsQuery {
		name := borrowQuery.Edges.User.Name

		borrow := dto.BorrowResponse{
			ID: borrowQuery.ID,
			User: dto.UserBorrowResponse{
				Name: *name,
			},
			ReleaseDate: borrowQuery.ReleaseDate,
			DueDate:     borrowQuery.DueDate,
		}

		borrows = append(borrows, borrow)
	}

	t.Log(borrows)
}
