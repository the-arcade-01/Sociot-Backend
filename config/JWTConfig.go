package config

import (
	"sociot/internal/utils"

	"github.com/go-chi/jwtauth/v5"
)

func GenerateAuthToken(jwtSecret string) *jwtauth.JWTAuth {
	tokenAuth := jwtauth.New(utils.JWT_ALGO, []byte(jwtSecret), nil)
	return tokenAuth
}
