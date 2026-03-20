package dto

import (
	"time"

	"github.com/alf4ridzi/library-crud-ent-echo/ent"
)

type BookResponse struct {
	Author            string            `json:"author" validate:"required"`
	Description       string            `json:"description" validate:"required"`
	Title             string            `json:"title" validate:"required"`
	Quantity          int               `json:"quantity" validate:"required"`
	AvailableQuantity int               `json:"available_quantity" validate:"required"`
	PublishDate       time.Time         `json:"publish_date" validate:"required"`
	Categories        []*ent.Categories `json:"category_ids" validate:"required,min=1"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
