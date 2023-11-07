package config

import (
	"log"
	"os"
	"sociot/internal/utils"

	"github.com/go-chi/jwtauth/v5"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	Token *jwtauth.JWTAuth
}

func LoadConfig() *AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error occurred while loading env file, %v\n", err)
	}
	jwtSecret := os.Getenv(utils.JWT_SECRET)

	appConfig := &AppConfig{
		Token: GenerateAuthToken(jwtSecret),
	}

	return appConfig
}
