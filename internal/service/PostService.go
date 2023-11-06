package service

import (
	"sociot/internal/entity"
	repo "sociot/internal/repository"
)

type PostService struct {
	repo repo.PostRepository
}

func NewPostService(postRepo repo.PostRepository) PostService {
	return PostService{
		repo: postRepo,
	}
}

func (service *PostService) GetPosts() entity.Response {
	posts := service.repo.GetPosts()
	response := entity.Response{
		Data: posts,
	}
	return response
}
