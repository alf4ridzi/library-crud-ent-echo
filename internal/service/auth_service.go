package service

import "github.com/alf4ridzi/library-crud-ent-echo/ent"

type AuthService interface{}

type authServiceImpl struct {
	DB *ent.Client
}

func NewAuthService(client *ent.Client) AuthService {
	return &authServiceImpl{DB: client}
}
