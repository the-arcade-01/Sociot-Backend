package service

import (
	"fmt"
	"net/http"
	"sociot/internal/entity"
	repo "sociot/internal/repository"
)

type GeneralService struct {
	repo repo.GeneralRepository
}

func NewGeneralService(repo repo.GeneralRepository) GeneralService {
	return GeneralService{
		repo: repo,
	}
}

func (service *GeneralService) Search(option string, search string) entity.Response {
	results, err := service.repo.Search(option, search)
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusInternalServerError)
		return response
	}
	message := `Results found, %v: %v`
	if option == "user" {
		message = fmt.Sprintf(message, "Users", len(results.Users))
	} else {
		message = fmt.Sprintf(message, "Posts", len(results.Posts))
	}

	response := entity.NewResponseObject(results, message, http.StatusOK)
	return response
}
