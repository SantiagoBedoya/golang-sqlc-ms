package user

import (
	"context"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repo   UserRepository
	logger *zap.Logger
}

func NewUserService(repo UserRepository, logger *zap.Logger) UserService {
	return &service{
		repo:   repo,
		logger: logger,
	}
}

func (s *service) SignIn(ctx context.Context, email, password string) (*User, error) {
	return nil, nil
}

func (s *service) SignUp(ctx context.Context, username, email, password string) (*User, error) {
	d, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return s.repo.CreateUser(ctx, User{
		Username: username,
		Email:    email,
		Password: string(d),
	})
}
