package app

import (
	companyinfo "github.com/CompaniesInfoStore/api/company_info"
	"github.com/CompaniesInfoStore/api/users"
	usrService "github.com/CompaniesInfoStore/internal/services/user"
)

func (a *Application) SetupHandlers() {
	middleware := usrService.AuthMiddleware(a.logger, a.db, a.config)
	companyinfo.RegisterHandlers(
		a.router,
		a.services.svc,
		a.logger,
		middleware.Auth,
	)
	users.RegisterHandlers(
		a.logger,
		a.router,
		a.services.usrSvc,
		middleware.Auth,
	)

}
