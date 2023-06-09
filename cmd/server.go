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

	categoryRepo := repository.NewCategoryRepo(DBConn)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryController := controller.NewCategoryController(categoryService)

	typeTransactionRepo := repository.NewTypeTransactionRepo(DBConn)
	typeTransactionService := service.NewTypeTransactionService(typeTransactionRepo)
	typeTransactionController := controller.NewTypeTransactionController(typeTransactionService)

	r.POST("/users", userController.Create)
	r.POST("/login", sessionController.Login)

	r.Use(middleware.AuthMiddleware(tokenService))

	r.GET("/users", userController.GetLists)

	r.POST("/currency", currencyController.CreateCurrency)
	r.GET("/currencies", currencyController.GetCurrenciesLists)
	r.GET("/currency/:id", currencyController.GetCurrencyByID)
	r.PUT("/currency/:id", currencyController.UpdateCurrency)
	r.DELETE("/currency/:id", currencyController.DeleteCurrency)

	r.POST("/category", categoryController.CreateCategory)
	r.GET("/categories", categoryController.GetCategoriesLists)
	r.GET("/category/:id", categoryController.GetCategoryByID)
	r.PUT("/category/:id", categoryController.UpdateCategory)
	r.DELETE("/category/:id", categoryController.DeleteCategory)

	r.POST("/type_transaction", typeTransactionController.CreateTypeTransaction)
	r.GET("/type_transactions", typeTransactionController.GetTypeTransactionsLists)
	r.GET("/type_transaction/:id", typeTransactionController.GetTypeTransactionByID)
	r.PUT("/type_transaction/:id", typeTransactionController.UpdateTypeTransaction)
	r.DELETE("/type_transaction/:id", typeTransactionController.DeleteTypeTransacton)

	s.router = r
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
