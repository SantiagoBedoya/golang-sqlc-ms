package user

import "context"

type UserService interface {
	SignIn(ctx context.Context, email, password string) (*User, error)
	SignUp(ctx context.Context, username, email, password string) (*User, error)
}
