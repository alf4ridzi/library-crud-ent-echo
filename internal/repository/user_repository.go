package repository

import "github.com/alf4ridzi/library-crud-ent-echo/ent"

type UserRepository interface {
}

type userRepositoryImpl struct {
	DB *ent.Client
}

func NewUserRepository(client *ent.Client) UserRepository {
	return &userRepositoryImpl{DB: client}
}
