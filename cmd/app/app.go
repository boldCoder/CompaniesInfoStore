package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/CompaniesInfoStore/internal/config"
	// m1 "github.com/CompaniesInfoStore/internal/services/company/model"
	m2 "github.com/CompaniesInfoStore/internal/services/user/model"
	"github.com/CompaniesInfoStore/pkg/database"
	"github.com/CompaniesInfoStore/pkg/logger"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

const timeout = 5 * time.Second

type Application struct {
	logger     zerolog.Logger
	config     *config.Config
	db         *gorm.DB
	services   *services
	router     *mux.Router
	httpServer *http.Server
}

func (a *Application) Init(ctx context.Context) {
	// Initialize logger instance
	logger := logger.NewLogger()
	a.logger = logger

	// Load Config
	config, err := config.Load(logger, "")
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to read config")
	}
	a.config = config

	zerolog.SetGlobalLevel(zerolog.Level(config.Logging.Level))
	
	// Initialize DB
	db, err := database.NewDB(config.DBConfig, logger)
	if err != nil {
		log.Fatalf("error connecting db: %s ", err)
		return
	}

	if err = db.AutoMigrate(&m2.User{}); err != nil {
		fmt.Println("[ERROR]: Migrating Database")
		return
	}
	a.db = db

	// Link to Service Layer
	services := buildServices(config, db)
	a.services = services

	router := mux.NewRouter()
	a.router = router

	// Register Handlers
	a.SetupHandlers()

}

func (a *Application) Start() {
	a.httpServer = &http.Server{
		Addr:              ":" + fmt.Sprintf("%v", a.config.Server.Port),
		Handler:           a.router,
		ReadHeaderTimeout: timeout,
	}

	defer a.logger.Info().Msg("server stopped running....")
	a.logger.Info().Msgf("server running at port %d...", a.config.Server.Port)
	if err := a.httpServer.ListenAndServe(); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}
