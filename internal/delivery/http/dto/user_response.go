package dto

import "time"

type UserResponse struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserBorrowResponse struct {
	Name string `json:"name"`
	// Email     string    `json:"email"`
	// Username  string    `json:"username"`
	// CreatedAt time.Time `json:"created_at"`
	// UpdatedAt time.Time `json:"updated_at"`
}
