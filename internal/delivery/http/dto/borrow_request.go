package dto

import "time"

type BorrowRequest struct {
	UserID  uint      `json:"user_id" validate:"required"`
	DueDate time.Time `json:"due_date" validate:"required"`
}
