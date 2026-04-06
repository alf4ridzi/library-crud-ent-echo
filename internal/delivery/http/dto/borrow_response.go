package dto

import "time"

type BorrowResponse struct {
	ID int `json:"id"`
	// Book        BookResponse       `json:"book"`
	User        UserBorrowResponse `json:"user"`
	ReleaseDate *time.Time         `json:"release_date"`
	DueDate     time.Time          `json:"due_date"`
}
