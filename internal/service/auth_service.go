package service

import (
	"context"
	"strings"

	"github.com/alf4ridzi/library-crud-ent-echo/ent"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/delivery/http/dto"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/pkg/cryptoutil"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/repository"
)

type AuthService interface {
	Login(ctx context.Context, req *dto.LoginRequest) (*dto.UserResponse, error)
	Register(ctx context.Context, reg *dto.RegisterRequest) error
}

type authServiceImpl struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authServiceImpl{userRepo: userRepo}
}

func (s *authServiceImpl) Login(ctx context.Context, req *dto.LoginRequest) (*dto.UserResponse, error) {
	var user *ent.User
	var err error

	if strings.Contains(req.Identifier, "@") {
		user, err = s.userRepo.FindByEmail(ctx, req.Identifier)
	} else {
		user, err = s.userRepo.FindByUsername(ctx, req.Identifier)
	}

	if err != nil {
		return nil, err
	}

	if !cryptoutil.ValidatePassword(user.Password, req.Password) {
		return nil, ErrInvalidCredentials
	}

	userRespons := &dto.UserResponse{
		Username: user.Username,
		Email:    user.Email,
		Name:     *user.Name,
	}

	return userRespons, nil
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
