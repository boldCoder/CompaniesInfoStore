package repo

import (
	"context"
	"errors"

	"github.com/CompaniesInfoStore/internal/services/user/model"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) AddUser(ctx context.Context, logger zerolog.Logger,
	user model.User) error {
	if tx := r.db.WithContext(ctx).Create(&user); tx.Error != nil {
		logger.Error().Err(tx.Error).Msgf("error while adding to database")
		return tx.Error
	}

	return nil
}

func (r *Repository) Validate(ctx context.Context, logger zerolog.Logger,
	email string) (model.User, error) {
	var user model.User
	if tx := r.db.WithContext(ctx).Where("email = ?", email).First(&user); tx.Error != nil {
		logger.Error().Err(tx.Error).Msgf("error fetching user from database")
		return model.User{}, tx.Error
	}

	if user.ID == 0 {
		logger.Error().Msgf("no user found")
		return model.User{}, errors.New("no user found")
	}

	return user, nil
}
