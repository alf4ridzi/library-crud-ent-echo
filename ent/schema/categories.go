package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Categories holds the schema definition for the Categories entity.
type Categories struct {
	ent.Schema
}

// Fields of the Categories.
func (Categories) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id"),
		field.String("name").Unique(),
	}
}

// Edges of the Categories.
func (Categories) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("books", Books.Type).Ref("categories"),
	}
}
