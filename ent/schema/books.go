package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Books holds the schema definition for the Books entity.
type Books struct {
	ent.Schema
}

// Fields of the Books.
func (Books) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id"),
		field.String("author"),
		field.String("description"),
		field.String("title"),
		field.Int("quantity"),
		field.Int("available_quantity"),
		field.Time("publis_date"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Books.
func (Books) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("borrowings", Borrowings.Type).Ref("book"),
		edge.To("categories", Categories.Type),
	}
}
