package postgres

import (
	"auth/internal/user"
	"auth/postgres/sqlc"
	"context"
)

type repository struct {
	queries *sqlc.Queries
}

func NewRepository(queries *sqlc.Queries) user.UserRepository {
	return &repository{
		queries: queries,
	}
}
func (r *repository) CreateUser(ctx context.Context, u user.User) (*user.User, error) {
	newUser, err := r.queries.CreateUser(ctx, sqlc.CreateUserParams{
		Email:    u.Email,
		Username: u.Username,
		Password: u.Password,
	})
	if err != nil {
		return nil, err
	}
	return &user.User{
		ID:       newUser.ID,
		Username: newUser.Username,
		Email:    newUser.Email,
	}, nil
}
func (r *repository) GetUserByEmail(ctx context.Context, email string) (*user.User, error) {
	u, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return &user.User{
		ID:       u.ID,
		Username: u.Username,
		Password: u.Password,
		Email:    u.Email,
	}, nil
}
