package repository

import (
	"context"

	"github.com/alf4ridzi/library-crud-ent-echo/ent"
	_ "github.com/alf4ridzi/library-crud-ent-echo/ent/runtime"
)

type UserRepository interface {
	Create(ctx context.Context, user ent.User) error
}

type userRepositoryImpl struct {
	DB *ent.Client
}

func NewUserRepository(client *ent.Client) UserRepository {
	return &userRepositoryImpl{DB: client}
}

func (r *userRepositoryImpl) Create(ctx context.Context, user ent.User) error {
	_, err := r.DB.User.Create().
		SetEmail(user.Email).
		SetName(*user.Name).
		SetUsername(user.Username).
		SetPassword(user.Password).
		Save(ctx)

	return err
}
