package service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth/v5"

	"sociot/internal/entity"
	repo "sociot/internal/repository"
)

type UserService struct {
	token *jwtauth.JWTAuth
	repo  repo.UserRepository
}

func NewUserService(userRepo repo.UserRepository, token *jwtauth.JWTAuth) UserService {
	return UserService{
		repo:  userRepo,
		token: token,
	}
}

func (service *UserService) GetUsers() entity.Response {
	users, err := service.repo.GetUsers()
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusInternalServerError)
		return response
	}
	response := entity.NewResponseObject(users, nil, http.StatusOK)
	return response
}

func (service *UserService) GetUserById(userId int) entity.Response {
	user, err := service.repo.GetUserById(userId)
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusInternalServerError)
		return response
	}
	response := entity.NewResponseObject(user, nil, http.StatusOK)
	return response
}

func (service *UserService) UpdateUserById(
	userId int,
	userBody *entity.UpdateUserNameReqBody,
) entity.Response {
	user := &entity.User{
		UserName: userBody.UserName,
	}
	err := service.repo.CheckExistingUser(user)
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusBadRequest)
		return response
	}
	err = service.repo.UpdateUserNameById(userId, user)
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusInternalServerError)
		return response
	}
	response := entity.NewResponseObject(
		nil,
		fmt.Sprintf("Username updated for User: %v", userId),
		http.StatusAccepted,
	)
	return response
}

func (service *UserService) UpdateUserPasswordById(
	userId int,
	userBody *entity.UpdateUserPasswordReqBody,
) entity.Response {
	user := &entity.User{
		Password: userBody.Password,
	}
	err := service.repo.UpdateUserPasswordById(userId, user)
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusInternalServerError)
		return response
	}
	response := entity.NewResponseObject(
		nil,
		fmt.Sprintf("Password updated for User: %v", userId),
		http.StatusAccepted,
	)
	return response
}

func (service *UserService) DeleteUserById(userId int) entity.Response {
	err := service.repo.DeleteUserById(userId)
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusInternalServerError)
		return response
	}
	response := entity.NewResponseObject(
		nil,
		fmt.Sprintf("Deleted User successfully, userId: %v", userId),
		http.StatusOK,
	)
	return response
}

func (service *UserService) CreateUser(userBody *entity.CreateUserRequestBody) entity.Response {
	user := &entity.User{
		UserName: userBody.UserName,
		Email:    userBody.Email,
		Password: userBody.Password,
	}
	err := service.repo.CheckExistingUser(user)
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusBadRequest)
		return response
	}
	err = service.repo.CreateUser(user)
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusInternalServerError)
		return response
	}

	response := entity.NewResponseObject(nil, "User created successfully", http.StatusCreated)
	return response
}

func (service *UserService) LoginUser(userBody *entity.LoginUserRequestBody) entity.Response {
	user := &entity.User{
		Email:    userBody.Email,
		Password: userBody.Password,
	}

	userDetails, err := service.repo.LoginUser(user)
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusUnauthorized)
		return response
	}

	expirationTime := time.Now().Add(72 * time.Hour)
	claims := map[string]interface{}{
		"userId":   userDetails.UserId,
		"email":    userDetails.Email,
		"userName": userDetails.UserName,
		"exp":      expirationTime,
	}
	_, tokenString, err := service.token.Encode(claims)
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusInternalServerError)
		return response
	}

	userLoginDetails := entity.UserLoginDetails{
		Token:  tokenString,
		UserId: userDetails.UserId,
	}

	response := entity.NewResponseObject(userLoginDetails, "User login successfully", http.StatusOK)
	return response
}

func (service *UserService) GetUserStats(userId int) entity.Response {
	user, err := service.repo.GetUserStats(userId)
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusInternalServerError)
		return response
	}
	response := entity.NewResponseObject(user, nil, http.StatusOK)
	return response
}
