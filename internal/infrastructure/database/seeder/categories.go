package seeder

import (
	"context"
	"log"

	"github.com/alf4ridzi/library-crud-ent-echo/ent"
)

var categories = []ent.Categories{
	{Name: "Fiction", Code: "fiction"},
	{Name: "Non-Fiction", Code: "non-fiction"},
	{Name: "Science", Code: "science"},
	{Name: "Technology", Code: "technology"},
	{Name: "History", Code: "history"},
	{Name: "Biography", Code: "biography"},
	{Name: "Fantasy", Code: "fantasy"},
	{Name: "Mystery", Code: "mystery"},
	{Name: "Self-Help", Code: "self-help"},
	{Name: "Children", Code: "children"},
}

func CategorySeed(client *ent.Client) {
	for _, category := range categories {
		_, err := client.Categories.Create().
			SetName(category.Name).
			SetCode(category.Code).Save(context.Background())

		if err != nil {
			log.Fatal(err)
		}
	}
}
