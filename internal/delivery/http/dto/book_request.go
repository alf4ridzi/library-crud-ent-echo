package dto

import "time"

type CreateNewBookRequest struct {
	Author            string    `json:"author" validate:"required"`
	Description       string    `json:"description" validate:"required"`
	Title             string    `json:"title" validate:"required"`
	Quantity          int       `json:"quantity" validate:"required"`
	AvailableQuantity int       `json:"available_quantity" validate:"required"`
	PublishDate       time.Time `json:"publish_date" validate:"required"`
	CategoryIDs       []uint    `json:"category_ids" validate:"required,min=1"`
}
