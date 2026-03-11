package service

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/alf4ridzi/library-crud-ent-echo/ent"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/delivery/http/dto"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/pkg/cryptoutil"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/pkg/tokenutil"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/repository"
)

type AuthService interface {
	RefreshToken(ctx context.Context, accessToken string) (*dto.AuthJwt, error)
	Login(ctx context.Context, req *dto.LoginRequest) (*dto.AuthJwt, error)
	Register(ctx context.Context, reg *dto.RegisterRequest) error
}

type authServiceImpl struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authServiceImpl{userRepo: userRepo}
}

func (s *authServiceImpl) RefreshToken(ctx context.Context, accessToken string) (*dto.AuthJwt, error) {
	claims, err := tokenutil.ClaimsRefreshToken(accessToken)
	if err != nil {
		return nil, err
	}

	userID, err := strconv.ParseUint(claims.Subject, 10, 64)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepo.FindByID(ctx, uint(userID))
	if err != nil {
		return nil, err
	}

	userIDString := strconv.FormatUint(uint64(user.ID), 10)

	generateAccessToken, err := tokenutil.GenerateAccessToken(userIDString, user.Edges.Role.Name, time.Duration(1)*time.Hour)
	if err != nil {
		return nil, err
	}

	return &dto.AuthJwt{
		Token: dto.AuthJwtResponse{
			Access: generateAccessToken,
		},
	}, nil
}

func (s *authServiceImpl) Login(ctx context.Context, req *dto.LoginRequest) (*dto.AuthJwt, error) {
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

	if user == nil {
		return nil, ErrInvalidCredentials
	}

	if !cryptoutil.ValidatePassword(user.Password, req.Password) {
		return nil, ErrInvalidCredentials
	}

	userID := strconv.FormatUint(uint64(user.ID), 10)

	accessToken, err := tokenutil.GenerateAccessToken(userID, user.Edges.Role.Name, time.Duration(1)*time.Hour)
	if err != nil {
		return nil, err
	}

	refreshToken, err := tokenutil.GenerateRefreshToken(userID, time.Duration(7)*24*time.Hour)
	if err != nil {
		return nil, err
	}

	return &dto.AuthJwt{
		Token: dto.AuthJwtResponse{
			Access:  accessToken,
			Refresh: refreshToken,
		},
	}, nil
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
