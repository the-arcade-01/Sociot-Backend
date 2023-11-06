package service

import (
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
	users := service.repo.GetUsers()
	response := entity.Response{
		Data: users,
	}
	return response
}
