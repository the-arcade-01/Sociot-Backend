package service

import (
	repo "sociot/internal/repository"
)

type CommentService struct {
	repo repo.CommentRepository
}

func NewCommentService(commentRepo repo.CommentRepository) *CommentService {
	return &CommentService{
		repo: commentRepo,
	}
}
