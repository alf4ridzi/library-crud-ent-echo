package schema

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/alf4ridzi/library-crud-ent-echo/ent/hook"
	"golang.org/x/crypto/bcrypt"

	entCrud "github.com/alf4ridzi/library-crud-ent-echo/ent"
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
		hook.On(hashPasswordHook(), ent.OpCreate),
	}
}

func hashPasswordHook() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {

			um, ok := m.(*entCrud.UserMutation)
			if !ok {
				return next.Mutate(ctx, m)
			}

			password, exists := um.Password()
			if !exists {
				return next.Mutate(ctx, m)
			}

			hashed, err := bcrypt.GenerateFromPassword(
				[]byte(password),
				bcrypt.DefaultCost,
			)
			if err != nil {
				return nil, fmt.Errorf("hash password failed: %w", err)
			}

			um.SetPassword(string(hashed))

			return next.Mutate(ctx, m)
		})
	}
}
