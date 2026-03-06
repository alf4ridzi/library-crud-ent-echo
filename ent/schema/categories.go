package schema

import "entgo.io/ent"

// Categories holds the schema definition for the Categories entity.
type Categories struct {
	ent.Schema
}

// Fields of the Categories.
func (Categories) Fields() []ent.Field {
	return nil
}

// Edges of the Categories.
func (Categories) Edges() []ent.Edge {
	return nil
}
