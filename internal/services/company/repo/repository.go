package repo

import (
	"context"
	"fmt"

	"github.com/CompaniesInfoStore/internal/services/company/model"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) Add(ctx context.Context, logger zerolog.Logger,
	companies []model.CompanyInfo) error {
	for _, details := range companies {
		if tx := r.db.WithContext(ctx).Create(details); tx.Error != nil {
			return tx.Error
		}
	}

	return nil
}

func (r *Repository) Fetch(ctx context.Context, logger zerolog.Logger,
	companyID uuid.UUID) (model.CompanyInfo, error) {
	company := model.CompanyInfo{}
	if tx := r.db.WithContext(ctx).Where("id=?", companyID).Find(&company); tx.Error != nil {
		return model.CompanyInfo{}, tx.Error
	}

	return company, nil
}

func (r *Repository) Update(ctx context.Context, logger zerolog.Logger,
	company model.CompanyInfo) (model.CompanyInfo, error) {

	var companyInfo model.CompanyInfo
	if tx := r.db.WithContext(ctx).Find(&companyInfo); tx.Error != nil {
		fmt.Println("error finding company details", tx.Error)
		return model.CompanyInfo{}, tx.Error
	}

	companyInfo.Id = company.Id
	companyInfo.CompanyName = company.CompanyName
	companyInfo.CompanyType = company.CompanyType
	companyInfo.Description = company.Description
	companyInfo.EmployeeStrength = company.EmployeeStrength
	companyInfo.Registered = company.Registered

	if tx := r.db.WithContext(ctx).Save(&companyInfo); tx.Error != nil {
		fmt.Println("error updating company details", tx.Error)
		return model.CompanyInfo{}, tx.Error
	}

	return companyInfo, nil
}

func (r *Repository) Delete(ctx context.Context, logger zerolog.Logger,
	companyID uuid.UUID) (string, error) {
	company := model.CompanyInfo{}
	if tx := r.db.WithContext(ctx).Where("id=?", companyID).Delete(&company); tx.Error != nil {
		return "record not deleted", tx.Error
	}

	return "record successfully deleted", nil

}
