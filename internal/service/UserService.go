package service

import (
	"net/http"
	"sociot/internal/entity"
	repo "sociot/internal/repository"
	"strconv"
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
	users := service.repo.GetUsers()
	response := entity.Response{
		Data: users,
		Meta: entity.Meta{
			StatusCode: http.StatusOK,
		},
	}
	return response
}

func (service *UserService) GetUserById(userId string) entity.Response {
	id, err := strconv.Atoi(userId)
	if err != nil {
		response := entity.Response{
			Meta: entity.Meta{
				Message:    err.Error(),
				StatusCode: http.StatusBadRequest,
			},
		}
		return response
	}
	user := service.repo.GetUserById(id)
	response := entity.Response{
		Data: user,
		Meta: entity.Meta{
			StatusCode: http.StatusOK,
		},
	}
	return response
}
