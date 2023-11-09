package service

import (
	"fmt"
	"net/http"
	"sociot/internal/entity"
	repo "sociot/internal/repository"
)

type UserService struct {
	repo repo.UserRepository
}

func NewUserService(userRepo repo.UserRepository) UserService {
	return UserService{
		repo: userRepo,
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
	user := service.repo.GetUserById(userId)
	response := entity.NewResponseObject(user, nil, http.StatusOK)
	return response
}

func (service *UserService) UpdateUserById(userId int, userBody *entity.UpdateUserRequestBody) entity.Response {
	err := service.repo.UpdateUserById(userId, userBody)
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusInternalServerError)
		return response
	}
	response := entity.NewResponseObject(nil, fmt.Sprintf("Details updated for User: %v", userId), http.StatusAccepted)
	return response
}

func (service *UserService) DeleteUserById(userId int) entity.Response {
	err := service.repo.DeleteUserById(userId)
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusInternalServerError)
		return response
	}
	response := entity.NewResponseObject(nil, fmt.Sprintf("Deleted User successfully, userId: %v", userId), http.StatusOK)
	return response
}

func (service *UserService) CreateUser(userBody *entity.CreateUserRequestBody) entity.Response {
	err := service.repo.CreateUser(userBody)
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusInternalServerError)
		return response
	}
	response := entity.NewResponseObject(nil, "User created successfully", http.StatusCreated)
	return response
}

func (service *UserService) LoginUser(userBody *entity.LoginUserRequestBody) entity.Response {
	err := service.repo.LoginUser(userBody)
	// here, actual code will, check user password in repo, and then generate jwt token and return it
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusInternalServerError)
		return response
	}
	response := entity.NewResponseObject(nil, "User login successfully", http.StatusOK)
	return response
}
