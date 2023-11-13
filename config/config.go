package config

import (
	"database/sql"
	"log"
	"os"
	"sociot/internal/utils"

	"github.com/go-chi/jwtauth/v5"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	Token *jwtauth.JWTAuth
	DB    *sql.DB
}

func LoadConfig() *AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error occurred while loading env file, %v\n", err)
	}
	jwtSecret := os.Getenv(utils.JWT_SECRET)
	dbUrl := os.Getenv(utils.DB_URL)
	dbDriver := os.Getenv(utils.DB_DRIVER)

	token := GenerateAuthToken(jwtSecret)
	db, err := LoadDBConfig(dbUrl, dbDriver)
	if err != nil {
		log.Fatalf("error occurred while loading DB Config, %v\n", err)
	}

	appConfig := &AppConfig{
		Token: token,
		DB:    db,
	}

	return appConfig
}
