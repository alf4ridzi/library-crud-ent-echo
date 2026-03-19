package repository

import (
	"context"

	"github.com/alf4ridzi/library-crud-ent-echo/ent"
	_ "github.com/alf4ridzi/library-crud-ent-echo/ent/runtime"
	"github.com/alf4ridzi/library-crud-ent-echo/ent/user"
)

type UserRepository interface {
	UpdateUserPassword(ctx context.Context, user *ent.User) error
	UpdateByUser(ctx context.Context, user *ent.User) error
	FindByEmail(ctx context.Context, email string) (*ent.User, error)
	FindByUsername(ctx context.Context, username string) (*ent.User, error)
	FindByUsernameOrEmail(ctx context.Context, username *string, email *string) (*ent.User, error)
	FindByID(ctx context.Context, id uint) (*ent.User, error)
	Create(ctx context.Context, user *ent.User) error
}

type userRepositoryImpl struct {
	DB *ent.Client
}

func NewUserRepository(client *ent.Client) UserRepository {
	return &userRepositoryImpl{DB: client}
}

func (r *userRepositoryImpl) UpdateUserPassword(ctx context.Context, user *ent.User) error {
	return r.DB.User.UpdateOne(user).
		SetPassword(user.Password).
		Exec(ctx)
}

func (r *userRepositoryImpl) UpdateByUser(ctx context.Context, user *ent.User) error {
	return r.DB.User.UpdateOne(user).
		SetEmail(user.Email).
		SetUsername(user.Username).
		SetName(*user.Name).
		Exec(ctx)
}

func (r *userRepositoryImpl) FindByID(ctx context.Context, id uint) (*ent.User, error) {
	return r.DB.User.Get(ctx, id)
}

func (r *userRepositoryImpl) FindByEmail(ctx context.Context, email string) (*ent.User, error) {
	return r.DB.User.Query().Where(
		user.Email(email),
	).WithRole().
		First(ctx)
}

func (r *userRepositoryImpl) FindByUsername(ctx context.Context, username string) (*ent.User, error) {
	return r.DB.User.Query().Where(
		user.Username(username),
	).WithRole().
		First(ctx)
}

func (r *userRepositoryImpl) FindByUsernameOrEmail(ctx context.Context, username *string, email *string) (*ent.User, error) {
	return r.DB.User.Query().Where(
		user.Or(
			user.Username(*username),
			user.Email(*email),
		),
	).First(ctx)
}

func (r *userRepositoryImpl) Create(ctx context.Context, user *ent.User) error {
	_, err := r.DB.User.Create().
		SetEmail(user.Email).
		SetName(*user.Name).
		SetUsername(user.Username).
		SetPassword(user.Password).
		Save(ctx)

	return err
}
