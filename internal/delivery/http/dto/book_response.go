package dto

import (
	"time"

	"github.com/alf4ridzi/library-crud-ent-echo/ent"
)

type BookResponse struct {
	ID                uint              `json:"id"`
	Author            string            `json:"author"`
	Description       string            `json:"description"`
	Title             string            `json:"title"`
	Quantity          int               `json:"quantity"`
	AvailableQuantity int               `json:"available_quantity"`
	PublishDate       time.Time         `json:"publish_date"`
	Categories        []*ent.Categories `json:"categories"`
	//Borrowings        []*ent.Borrowings `json:"borrowings"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
