package service

import (
	repo "sociot/internal/repository"
)

type UserService struct {
	repo repo.UserRepository
}

func NewUserService(userRepo repo.UserRepository) *UserService {
	return &UserService{
		repo: userRepo,
	}
}
