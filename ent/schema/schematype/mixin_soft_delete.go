package schematype

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/alf4ridzi/library-crud-ent-echo/ent/hook"
)

type SoftDeleteMixin struct {
	mixin.Schema
}

func (SoftDeleteMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("deleted_at").
			Optional(),
	}
}

type softDeleteKey struct{}

func SkipSoftDelete(parent context.Context) context.Context {
	return context.WithValue(parent, softDeleteKey{}, true)
}

func (d SoftDeleteMixin) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		ent.TraverseFunc(func(ctx context.Context, q ent.Query) error {
			// Check for skip flag
			if skip, _ := ctx.Value(softDeleteKey{}).(bool); skip {
				return nil
			}

			// Use the P method you already defined
			// We cast to the interface you defined at the bottom of your file
			if wq, ok := q.(interface{ WhereP(...func(*sql.Selector)) }); ok {
				d.P(wq)
				return nil
			}

			// If the above fails, it's likely because the generated code
			// isn't matching the interface. Try the manual SQL injection:
			if traverser, ok := q.(interface {
				Modify(modifiers ...func(*sql.Selector))
			}); ok {
				traverser.Modify(func(s *sql.Selector) {
					s.Where(sql.IsNull(s.C("deleted_at")))
				})
				return nil
			}

			return nil // Or return the error if you want to be strict
		}),
	}
}

func (d SoftDeleteMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
					if skip, _ := ctx.Value(softDeleteKey{}).(bool); skip {
						return next.Mutate(ctx, m)
					}

					type softDeleteMutation interface {
						SetOp(ent.Op)
						SetDeletedAt(time.Time)
						WhereP(...func(*sql.Selector))
					}

					mx, ok := m.(softDeleteMutation)
					if !ok {
						return nil, fmt.Errorf("unexpected mutation type %T", m)
					}

					d.P(mx)
					mx.SetOp(ent.OpUpdate)
					mx.SetDeletedAt(time.Now())
					return next.Mutate(ctx, m)
				})
			},
			ent.OpDeleteOne|ent.OpDelete,
		),
	}
}

func (d SoftDeleteMixin) P(w interface{ WhereP(...func(*sql.Selector)) }) {
	w.WhereP(
		sql.FieldIsNull("deleted_at"),
	)
}
