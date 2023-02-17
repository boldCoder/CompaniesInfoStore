package user

import (
	"context"
	"fmt"
	"time"

	"github.com/CompaniesInfoStore/internal/config"
	"github.com/CompaniesInfoStore/internal/services/user/model"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rs/zerolog"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo model.Repository
	cfg  *config.Config
}

func NewService(repo model.Repository, cfg *config.Config) *Service {
	return &Service{repo: repo, cfg: cfg}
}

func (s *Service) Signup(ctx context.Context, logger zerolog.Logger, email string, password string) error {
	usrRepo := model.User{Email: email, Password: password}
	if err := s.repo.AddUser(ctx, logger, usrRepo); err != nil {
		return err
	}
	return nil
}

func (s *Service) ValidateUser(ctx context.Context, logger zerolog.Logger, email string, password string) (string, error) {
	user, err := s.repo.Validate(ctx, logger, email)
	if err != nil {
		logger.Error().Err(err).Msg("error validating user")
		return "", err
	}

	// compare hash received with password
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		logger.Error().Err(err).Msg("wrong password")
		return "", err
	}

	// generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Password,
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	fmt.Println("[SECRET]: ", s.cfg.Secret, []byte(s.cfg.Secret))
	tokenString, err := token.SignedString([]byte(s.cfg.Secret))
	if err != nil {
		logger.Error().Err(err).Msgf("error creating jwt string")
		return "", err
	}

	return tokenString, nil
}
