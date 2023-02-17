package services

import (
	"context"
	"fmt"

	"github.com/CompaniesInfoStore/internal/services/company/model"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type Service struct {
	repo model.Repository
}

func NewService(repo model.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) AddCompany(ctx context.Context, logger zerolog.Logger, companies []model.CompanyInfo) error {
	for i := range companies {
		ID := uuid.New()
		companies[i].Id = ID
	}

	err := s.repo.Add(ctx, logger, companies)
	if err != nil {
		return fmt.Errorf("failed to store company: %w", err)
	}

	return nil
}

func (s *Service) GetCompanyByID(ctx context.Context, logger zerolog.Logger, id uuid.UUID) (model.CompanyInfo, error) {
	resp, err := s.repo.Fetch(ctx, logger, id)
	if err != nil {
		return model.CompanyInfo{}, fmt.Errorf("failed to fetch company with ID: %w", err)
	}

	return resp, nil
}

func (s *Service) UpdateCompanyDetails(ctx context.Context, logger zerolog.Logger, company model.CompanyInfo) (model.CompanyInfo, error) {
	fmt.Println(">>>>>>>>>> coming in UpdateCompanyDetails")

	resp, err := s.repo.Update(ctx, logger, company)
	if err != nil {
		return model.CompanyInfo{}, fmt.Errorf("failed to update company details: %w", err)
	}

	return resp, nil
}

func (s *Service) DeleteByCompanyID(ctx context.Context, logger zerolog.Logger, id uuid.UUID) (string, error) {
	resp, err := s.repo.Delete(ctx, logger, id)
	if err != nil {
		return "", fmt.Errorf("failed to fetch company with ID: %w", err)
	}

	return resp, nil
}
