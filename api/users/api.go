package users

import (
	"encoding/json"
	"net/http"

	"github.com/CompaniesInfoStore/api"
	"github.com/CompaniesInfoStore/internal/services/user/model"
	"github.com/rs/zerolog"
	"golang.org/x/crypto/bcrypt"

	"github.com/gorilla/mux"
)

type resource struct {
	services model.Service
	logger   zerolog.Logger
}

func RegisterHandlers(logger zerolog.Logger, 
	router *mux.Router, 
	svc model.Service, 
	auth mux.MiddlewareFunc) {
	res := resource{svc, logger}
	r := router.PathPrefix("/user").Subrouter()
	api.RegisterHandler(r, "POST", "/signup", nil, res.signup)
	api.RegisterHandler(r, "POST", "/login", nil, res.login)
}

func (res *resource) signup(w http.ResponseWriter, r *http.Request) {
	req := model.User{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		res.logger.Error().Err(err).Msg("failed to decode request")
		api.JsonResponse(w, http.StatusBadRequest, err)
		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		res.logger.Error().Err(err).Msg("failed to generate password hash")
		api.JsonResponse(w, http.StatusBadRequest, err)
		return
	}

	err = res.services.Signup(r.Context(), res.logger, req.Email, string(hash))
	if err != nil {
		res.logger.Error().Err(err).Msg("user signup failed")
		api.JsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	api.JsonResponse(w, http.StatusOK, "user created successfully")
}

func (res *resource) login(w http.ResponseWriter, r *http.Request) {
	req := model.User{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		res.logger.Error().Err(err).Msg("failed to decode request")
		api.JsonResponse(w, http.StatusBadRequest, err)
		return
	}

	if req.Email == "" || req.Password == "" {
		res.logger.Error().Msg("Invalid email or password")
		api.JsonResponse(w, http.StatusInternalServerError, "Invalid email or password")
		return
	}

	result, err := res.services.ValidateUser(r.Context(), res.logger, req.Email, req.Password)
	if err != nil {
		api.JsonResponse(w, http.StatusInternalServerError, nil)
		return
	}

	cookie := &http.Cookie{Name: "Authorization", Value: result,
		MaxAge:   3600,
		Secure:   false,
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)
	api.JsonResponse(w, http.StatusOK, nil)
}
