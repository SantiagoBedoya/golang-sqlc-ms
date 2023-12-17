package user

import (
	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user User) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}
