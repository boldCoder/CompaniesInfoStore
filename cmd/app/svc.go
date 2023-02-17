package app

import (
	"github.com/CompaniesInfoStore/internal/config"
	companySvc "github.com/CompaniesInfoStore/internal/services/company"
	companyModel "github.com/CompaniesInfoStore/internal/services/company/model"
	companyRepo "github.com/CompaniesInfoStore/internal/services/company/repo"
	userSvc "github.com/CompaniesInfoStore/internal/services/user"
	userModel "github.com/CompaniesInfoStore/internal/services/user/model"
	userRepo "github.com/CompaniesInfoStore/internal/services/user/repo"

	"gorm.io/gorm"
)

type services struct {
	svc    companyModel.Service
	usrSvc userModel.Service
}

type repos struct {
	companyRepo companyModel.Repository
	usrRepo     userModel.Repository
}

func buildServices(cfg *config.Config, db *gorm.DB) *services {
	svc := &services{}
	repo := &repos{}
	repo.buildRepos(db)
	svc.buildService(repo, cfg)

	return svc
}

func (r *repos) buildRepos(db *gorm.DB) {
	r.companyRepo = companyRepo.NewRepository(db)
	r.usrRepo = userRepo.NewRepository(db)
}

func (s *services) buildService(repo *repos, cfg *config.Config) {
	s.svc = companySvc.NewService(repo.companyRepo)
	s.usrSvc = userSvc.NewService(repo.usrRepo, cfg)
}
