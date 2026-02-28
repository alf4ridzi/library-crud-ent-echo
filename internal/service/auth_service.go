package service

import (
	"context"

	"github.com/alf4ridzi/library-crud-ent-echo/ent"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/delivery/http/dto"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/repository"
)

type AuthService interface {
	Register(ctx context.Context, reg *dto.RegisterRequest) error
}

type authServiceImpl struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authServiceImpl{userRepo: userRepo}
}

func (s *authServiceImpl) Register(ctx context.Context, reg *dto.RegisterRequest) error {
	user := &ent.User{
		Name:     &reg.Name,
		Username: reg.Username,
		Email:    reg.Email,
		Password: reg.Password,
	}

	return s.userRepo.Create(ctx, user)
}
