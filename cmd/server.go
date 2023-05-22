package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	"github.com/go-finance/internal/app/controller"
	"github.com/go-finance/internal/app/repository"
	"github.com/go-finance/internal/app/service"
	"github.com/go-finance/internal/pkg/config"
	"github.com/go-finance/internal/pkg/middleware"
)

type Server struct {
	cfg    config.Config
	dbConn *sqlx.DB
	router *gin.Engine
}

func NewServer(cfg config.Config, DBConn *sqlx.DB) (*Server, error) {
	server := &Server{
		cfg:    cfg,
		dbConn: DBConn,
	}

	// setup router
	server.setupRouter()

	return server, nil
}

func (s *Server) setupRouter() {
	r := gin.New()

	r.Use(
		middleware.LoggingMiddleware(),
		middleware.RecoveryMiddleware(),
	)

	tokenService := service.NewGenerateToken(
		s.cfg.AccessTokenKey,
		s.cfg.RefreshTokenKey,
		s.cfg.AccessTokenDuration,
		s.cfg.RefreshTokenDuration,
	)

	userRepo := repository.NewUserRepo(DBConn)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	authRepo := repository.NewAuthRepo(DBConn)
	sessionService := service.NewSessionService(userRepo, authRepo, tokenService)
	sessionController := controller.NewSessionController(sessionService, tokenService)

	currencyRepo := repository.NewCurrencyRepo(DBConn)
	currencyService := service.NewCurrencyService(currencyRepo)
	currencyController := controller.NewCurrencyController(currencyService)

	r.POST("/users", userController.Create)
	r.POST("/login", sessionController.Login)

	r.Use(middleware.AuthMiddleware(tokenService))

	r.GET("/users", userController.GetLists)

	r.POST("/currency", currencyController.Create)
	r.GET("/currencies", currencyController.GetCurrenciesLists)

	s.router = r
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
