package schema

import "entgo.io/ent"

// Books holds the schema definition for the Books entity.
type Books struct {
	ent.Schema
}

// Fields of the Books.
func (Books) Fields() []ent.Field {
	return nil
}

// Edges of the Books.
func (Books) Edges() []ent.Edge {
	return nil
}
