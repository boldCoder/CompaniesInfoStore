package companyinfo

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CompaniesInfoStore/api"
	"github.com/CompaniesInfoStore/internal/services/company/model"
	"github.com/google/uuid"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

type resource struct {
	service model.Service
	logger  zerolog.Logger
}

func RegisterHandlers(router *mux.Router, 
	svc model.Service, logger zerolog.Logger, auth mux.MiddlewareFunc) {
	res := resource{service: svc, logger: logger}
	r := router.PathPrefix("/company").Subrouter()
	api.RegisterHandler(r, "POST", "/create", nil, res.addCompany, auth)
	api.RegisterHandler(r, "PATCH", "/update", nil, res.updateCompany, auth)
	api.RegisterHandler(r, "GET", "/get/{id}", nil, res.getCompanyByID, auth)
	api.RegisterHandler(r, "DELETE", "/delete/{id}", nil, res.deleteCompanyByID, auth)

}

func (res resource) addCompany(w http.ResponseWriter, r *http.Request) {
	req := []model.CompanyInfo{}

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println("[ERROR]: ", err)
		res.logger.Error().Err(err).Msg("failed to decode request")
		return
	}

	err := res.service.AddCompany(r.Context(), res.logger, req)
	if err != nil {
		res.logger.Err(err).Msg("unable to insert record")
		api.JsonResponse(w, http.StatusInternalServerError, err)
		return
	}

	api.JsonResponse(w, http.StatusOK, req)
}

func (res resource) getCompanyByID(w http.ResponseWriter, r *http.Request) {
	companyID := mux.Vars(r)["id"]

	uid, err := uuid.Parse(companyID)
	if err != nil {
		fmt.Println("[ERROR]: %w", err)
		return
	}

	response, err := res.service.GetCompanyByID(r.Context(), res.logger, uid)
	if err != nil {
		api.JsonResponse(w, http.StatusInternalServerError, err)
	}
	api.JsonResponse(w, http.StatusOK, response)
}

func (res resource) deleteCompanyByID(w http.ResponseWriter, r *http.Request) {
	companyID := mux.Vars(r)["id"]

	uid, err := uuid.Parse(companyID)
	if err != nil {
		fmt.Println("[ERROR]: %w", err)
		return
	}

	response, err := res.service.DeleteByCompanyID(r.Context(), res.logger, uid)
	if err != nil {
		api.JsonResponse(w, http.StatusInternalServerError, err)
	}

	api.JsonResponse(w, http.StatusOK, response)
}

func (res resource) updateCompany(w http.ResponseWriter, r *http.Request) {
	req := model.CompanyInfo{}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	if err := decoder.Decode(&req); err != nil {
		fmt.Println("[ERROR]: ", err)
		res.logger.Error().Err(err).Msg("failed to decode request")
		return
	}

	fmt.Println(">>>>>>>>>>", req)

	resp, err := res.service.UpdateCompanyDetails(r.Context(), res.logger, req)
	if err != nil {
		res.logger.Err(err).Msg("unable to update record")
		api.JsonResponse(w, http.StatusInternalServerError, err)
		return
	}

	api.JsonResponse(w, http.StatusOK, resp)
}
