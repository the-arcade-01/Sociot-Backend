package service

import (
	repo "sociot/internal/repository"
)

type PostService struct {
	repo repo.PostRepository
}

func NewPostService(postRepo repo.PostRepository) *PostService {
	return &PostService{
		repo: postRepo,
	}
}
