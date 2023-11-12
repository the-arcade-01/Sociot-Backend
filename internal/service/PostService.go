package service

import (
	"fmt"
	"net/http"
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
	posts, err := service.repo.GetPosts()
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusInternalServerError)
		return response
	}
	response := entity.NewResponseObject(posts, nil, http.StatusOK)
	return response
}

func (service *PostService) GetPostById(postId int) entity.Response {
	post, err := service.repo.GetPostById(postId)
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusInternalServerError)
		return response
	}
	response := entity.NewResponseObject(post, nil, http.StatusOK)
	return response
}

func (service *PostService) CreatePost(postBody *entity.PostRequestBody) entity.Response {
	err := service.repo.CreatePost(postBody)
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusInternalServerError)
		return response
	}
	response := entity.NewResponseObject(nil, "Post created successfully", http.StatusCreated)
	return response
}

func (service *PostService) UpdatePostById(postId int, postBody *entity.UpdatePostRequestBody) entity.Response {
	err := service.repo.UpdatePostById(postId, postBody)
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusInternalServerError)
		return response
	}
	response := entity.NewResponseObject(nil, fmt.Sprintf("Post updated successfully, PostId: %v", postId), http.StatusOK)
	return response
}

func (service *PostService) DeletePostById(postId int) entity.Response {
	err := service.repo.DeletePostById(postId)
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusInternalServerError)
		return response
	}
	response := entity.NewResponseObject(nil, fmt.Sprintf("Post deleted successfully, PostId: %v", postId), http.StatusOK)
	return response
}
