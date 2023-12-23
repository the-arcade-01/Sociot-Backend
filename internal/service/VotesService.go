package service

import (
	"net/http"
	"sociot/internal/entity"
	"sociot/internal/repository"
)

type VotesService struct {
	repo repository.VotesRepo
}

func NewVotesService(repo repository.VotesRepo) VotesService {
	return VotesService{
		repo: repo,
	}
}

func (service VotesService) UpdatePostVotesById(postId int, userId int, voteType string) entity.Response {
	err := service.repo.UpdatePostVotesById(postId, userId, voteType)
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusInternalServerError)
		return response
	}
	response := entity.NewResponseObject(nil, nil, http.StatusOK)
	return response
}

func (service VotesService) GetVotesCountById(postId int) entity.Response {
	count, err := service.repo.GetVotesCountById(postId)
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusInternalServerError)
		return response
	}
	response := entity.NewResponseObject(count, nil, http.StatusOK)
	return response
}

func (service VotesService) GetUserVoted(postId int, userId int) entity.Response {
	value, err := service.repo.GetUserVoted(postId, userId)
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusInternalServerError)
		return response
	}
	response := entity.NewResponseObject(value, nil, http.StatusOK)
	return response
}
