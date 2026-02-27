package service

import (
	"github.com/alf4ridzi/library-crud-ent-echo/internal/repository"
)

type AuthService interface{}

type authServiceImpl struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authServiceImpl{userRepo: userRepo}
}
