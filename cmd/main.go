package main

import (
	"fmt"

	"github.com/go-finance/internal/pkg/config"
	"github.com/go-finance/internal/pkg/db"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

var (
	cfg    config.Config
	DBConn *sqlx.DB
)

func init() {
	configLoad, err := config.LoadConfig(".")
	if err != nil {
		log.Println("Unable to load config file")
		return
	}
	cfg = configLoad

	db, err := db.ConnectDB(cfg.DBDriver, cfg.DBConnection)
	if err != nil {
		log.Println("Database Unavailable")
		return
	}
	DBConn = db

	logLevel, err := log.ParseLevel("debug")
	if err != nil {
		logLevel = log.InfoLevel
	}

	log.SetLevel(logLevel)
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	server, err := NewServer(cfg, DBConn)
	if err != nil {
		log.Error(fmt.Errorf("Server Unavailable"))
		return
	}

	appPort := fmt.Sprintf(":%s", cfg.ServerPort)
	err = server.Start(appPort)
	if err != nil {
		log.Error(fmt.Errorf("Server Unable To Start"))
		return

	}
}
