package utils

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sociot/internal/entity"

	"github.com/go-chi/jwtauth/v5"
	"github.com/go-playground/validator/v10"
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

func ValidateRequestBody(body any) (entity.Response, error) {
	validate := validator.New()
	err := validate.Struct(body)
	var errMsg string
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errMsg += getErrMsg(err).Error()
		}
	}
	response := entity.NewResponseObject(nil, errMsg, http.StatusBadRequest)
	return response, err
}

func getErrMsg(err validator.FieldError) error {
	switch err.Tag() {
	case "required":
		return fmt.Errorf("%v: is required", err.Field())
	case "min":
		return fmt.Errorf("min length of %v is 4", err.Field())
	case "email":
		return fmt.Errorf("please provide a valid email")
	}
	return fmt.Errorf("please check the required params")
}
