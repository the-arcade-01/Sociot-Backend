package service

import (
	"sociot/internal/entity"
	repo "sociot/internal/repository"
)

type CommentService struct {
	repo repo.CommentRepository
}

func NewCommentService(commentRepo repo.CommentRepository) CommentService {
	return CommentService{
		repo: commentRepo,
	}
}

func (service *CommentService) GetCommentById() entity.Response {
	comment := service.repo.GetCommentById()
	response := entity.Response{
		Data: comment,
	}
	return response
}
