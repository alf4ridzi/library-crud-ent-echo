package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Borrowings holds the schema definition for the Borrowings entity.
type Borrowings struct {
	ent.Schema
}

// Fields of the Borrowings.
func (Borrowings) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("book_id"),
		field.Uint("user_id"),
		field.Time("release_date"),
		field.Time("due_date"),
	}
}

// Edges of the Borrowings.
func (Borrowings) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Field("user_id").Unique().Required(),
		edge.To("book", Books.Type).Field("book_id").Unique().Required(),
	}
}
