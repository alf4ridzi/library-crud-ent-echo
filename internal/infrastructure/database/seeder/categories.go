package seeder

import (
	"context"
	"log"

	"github.com/alf4ridzi/library-crud-ent-echo/ent"
	entCategories "github.com/alf4ridzi/library-crud-ent-echo/ent/categories"
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
	var categoriesCreate []*ent.CategoriesCreate

	for _, category := range categories {
		c := category

		// check if categories exist
		cat, err := client.Categories.Query().Where(
			entCategories.Code(c.Code),
		).First(context.Background())

		if !ent.IsNotFound(err) {
			continue
		}

		if cat != nil {
			continue
		}

		create := client.Categories.
			Create().
			SetName(c.Name).
			SetCode(c.Code)

		categoriesCreate = append(categoriesCreate, create)
	}

	_, err := client.Categories.CreateBulk(categoriesCreate...).Save(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
