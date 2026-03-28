package dto

import "time"

type BorrowRequest struct {
	UserID      uint      `json:"user_id" validate:"required"`
	ReleaseDate time.Time `json:"release_date" validate:"required"`
	DueDate     time.Time `json:"due_date" validate:"required"`
}
