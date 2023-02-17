package user

import (
	"fmt"
	"net/http"
	"time"

	"github.com/CompaniesInfoStore/api"
	"github.com/CompaniesInfoStore/internal/config"
	"github.com/CompaniesInfoStore/internal/services/user/model"

	"github.com/golang-jwt/jwt/v4"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type Middleware struct {
	logger zerolog.Logger
	db     *gorm.DB
	cfg    *config.Config
}

func AuthMiddleware(logger zerolog.Logger, db *gorm.DB, cfg *config.Config) *Middleware {
	return &Middleware{logger: logger, db: db, cfg: cfg}
}

func (m *Middleware) Auth(f http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idToken := r.Header.Get("Authorization")
		if idToken == "" {
			m.logger.Error().Msg("token not found")
			api.JsonResponse(w, http.StatusUnauthorized, "token not found")
			return
		}

		token, _ := jwt.Parse(idToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(m.cfg.Secret), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				api.JsonResponse(w, http.StatusBadRequest, "token expired")
				return
			}

			var user model.User
			m.db.WithContext(r.Context()).First(&user, "password=?", claims["sub"])

			if user.ID == 0 {
				api.JsonResponse(w, http.StatusUnauthorized, "unauthorized user")
				return
			}

		} else {
			api.JsonResponse(w, http.StatusUnauthorized, "user not authorized")
			return
		}

		f.ServeHTTP(w, r)
	})

}
