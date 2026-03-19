package hook

import (
	"context"
	"errors"

	"github.com/alf4ridzi/library-crud-ent-echo/ent"
	entCrud "github.com/alf4ridzi/library-crud-ent-echo/ent"
	"github.com/alf4ridzi/library-crud-ent-echo/ent/role"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/pkg/cryptoutil"
)

func DefaultRoleHook(client *entCrud.Client) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			um, ok := m.(*entCrud.UserMutation)
			if !ok {
				return next.Mutate(ctx, m)
			}

			role, err := client.Role.Query().Where(
				role.Name("user"),
			).First(context.Background())

			if err != nil {
				return err, errors.New("user role not found")
			}

			um.SetRoleID(role.ID)

			return next.Mutate(ctx, m)
		})
	}
}
func HashPasswordHook() ent.Hook {
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

			hashed := cryptoutil.HashPassword(password)

			um.SetPassword(hashed)

			return next.Mutate(ctx, m)
		})
	}
}
