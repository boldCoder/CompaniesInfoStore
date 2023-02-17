package model

import (
	"context"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type Service interface {
	AddCompany(context.Context, zerolog.Logger, []CompanyInfo) error
	GetCompanyByID(context.Context, zerolog.Logger, uuid.UUID) (CompanyInfo, error)
	UpdateCompanyDetails(context.Context, zerolog.Logger, CompanyInfo) (CompanyInfo, error)
	DeleteByCompanyID(context.Context, zerolog.Logger, uuid.UUID) (string, error)
}

type Repository interface {
	Add(context.Context, zerolog.Logger, []CompanyInfo) error
	Fetch(context.Context, zerolog.Logger, uuid.UUID) (CompanyInfo, error)
	Update(context.Context, zerolog.Logger, CompanyInfo) (CompanyInfo, error)
	Delete(context.Context, zerolog.Logger, uuid.UUID) (string, error)
}
