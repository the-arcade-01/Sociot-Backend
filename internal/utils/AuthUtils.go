package utils

import (
	"context"
	"errors"

	"github.com/go-chi/jwtauth/v5"
	"golang.org/x/crypto/bcrypt"
)

func GetHashPassword(password string) (string, error) {
	bytePassword := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CheckPassword(hashPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err
}

func ValidateAuthToken(userId int, ctx context.Context) error {
	_, claims, _ := jwtauth.FromContext(ctx)
	userIdFromClaims := int(claims[USER_ID].(float64))
	if userId != userIdFromClaims {
		return errors.New("invalid Auth token, please provide correct token")
	}
	return nil
}
