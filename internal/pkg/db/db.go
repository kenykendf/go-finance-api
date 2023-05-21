package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	log "github.com/sirupsen/logrus"
)

func ConnectDB(DBDriver string, DBConn string) (*sqlx.DB, error) {
	db, err := sqlx.Connect(DBDriver, DBConn)
	if err != nil {
		fmt.Println("CHECK ERR", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Database Now Available")
	return db, nil
}
