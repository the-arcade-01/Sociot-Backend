package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func LoadDBConfig(dbUrl string, dbDriver string) (*sql.DB, error) {
	db, err := sql.Open(dbDriver, dbUrl)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	log.Println("Connected to DB")
	return db, nil
}
