package schema

import "entgo.io/ent"

// Borrowings holds the schema definition for the Borrowings entity.
type Borrowings struct {
	ent.Schema
}

// Fields of the Borrowings.
func (Borrowings) Fields() []ent.Field {
	return nil
}

// Edges of the Borrowings.
func (Borrowings) Edges() []ent.Edge {
	return nil
}
