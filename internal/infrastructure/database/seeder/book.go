package seeder

import (
	"context"
	"log"
	"time"

	"github.com/alf4ridzi/library-crud-ent-echo/ent"
)

var books = []ent.Books{
	{
		Author:            "Jack Hanma",
		Description:       "Ini adalah buku fiksi tentang boneka tua di rumah terbengkalai",
		Title:             "Old Doll",
		Quantity:          10,
		AvailableQuantity: 10,
		PublishDate:       time.Now(),
		Edges: ent.BooksEdges{
			Categories: []*ent.Categories{
				{
					ID: 1,
				},
				{
					ID: 6,
				},
				{
					ID: 8,
				},
			},
		},
	},
}

func BookSeeder(client *ent.Client) {
	var bookscreate []*ent.BooksCreate

	for _, book := range books {
		b := book

		create := client.Books.Create().
			SetAuthor(b.Author).
			SetTitle(b.Title).
			SetDescription(b.Description).
			SetQuantity(b.Quantity).
			SetAvailableQuantity(b.AvailableQuantity).
			SetPublishDate(b.PublishDate)

		for _, cat := range b.Edges.Categories {
			create = create.AddCategoryIDs(cat.ID)
		}

		bookscreate = append(bookscreate, create)
	}

	_, err := client.Books.CreateBulk(bookscreate...).Save(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
