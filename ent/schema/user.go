package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/alf4ridzi/library-crud-ent-echo/ent/hook"
	"github.com/alf4ridzi/library-crud-ent-echo/ent/schema/schematype"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id").Positive(),
		field.String("name").Nillable(),
		field.String("email").Unique(),
		field.String("username").Unique(),
		field.String("password"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("borrowings", Borrowings.Type).Ref("user"),
		edge.To("role", Role.Type).Unique(),
	}
}

func (User) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(hook.HashPasswordHook(), ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schematype.SoftDeleteMixin{},
	}
}
