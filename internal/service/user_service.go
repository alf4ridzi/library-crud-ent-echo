package service

import (
	"context"
	"errors"
	"strconv"

	"github.com/alf4ridzi/library-crud-ent-echo/internal/delivery/http/dto"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/repository"
)

type UserService interface {
	GetByID(ctx context.Context, id any) (*dto.UserResponse, error)
}

type userServiceImpl struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userServiceImpl{userRepo: userRepo}
}

func (s *userServiceImpl) GetByID(ctx context.Context, id any) (*dto.UserResponse, error) {
	userIDString, ok := id.(string)
	if !ok {
		return nil, errors.New("id is not string")
	}

	userID, err := strconv.ParseUint(userIDString, 10, 64)
	if err != nil {
		return nil, err

	}

	userIDUInt := uint(userID)

	user, err := s.userRepo.FindByID(ctx, userIDUInt)
	if err != nil {
		return nil, err
	}

	userResponse := &dto.UserResponse{
		Name:      *user.Name,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return userResponse, nil
}
