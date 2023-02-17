package model

import (
	"context"

	"github.com/rs/zerolog"
)

type Service interface {
	Signup(context.Context, zerolog.Logger, string, string) error
	ValidateUser(context.Context, zerolog.Logger, string, string) (string, error)
}

type Repository interface {
	AddUser(context.Context, zerolog.Logger, User) error
	Validate(context.Context, zerolog.Logger, string) (User, error)
}
