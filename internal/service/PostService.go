package service

import (
	"fmt"
	"net/http"
	"sociot/internal/entity"
	repo "sociot/internal/repository"

	"github.com/go-chi/jwtauth/v5"
)

type PostService struct {
	repo  repo.PostRepository
	token *jwtauth.JWTAuth
}

func NewPostService(postRepo repo.PostRepository, token *jwtauth.JWTAuth) PostService {
	return PostService{
		repo:  postRepo,
		token: token,
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
	post := &entity.Post{
		UserId:  postBody.UserId,
		Title:   postBody.Title,
		Content: postBody.Content,
	}
	err := service.repo.CreatePost(post)
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusInternalServerError)
		return response
	}
	response := entity.NewResponseObject(nil, "Post created successfully", http.StatusCreated)
	return response
}

func (service *PostService) UpdatePostById(postId int, postBody *entity.UpdatePostRequestBody) entity.Response {
	post := &entity.Post{
		PostId:  postId,
		Title:   postBody.Title,
		Content: postBody.Content,
	}
	err := service.repo.UpdatePostById(post)
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusInternalServerError)
		return response
	}
	response := entity.NewResponseObject(nil, fmt.Sprintf("Post updated successfully, PostId: %v", postId), http.StatusOK)
	return response
}

func (service *PostService) DeletePostById(postId int, userId int) entity.Response {
	post, err := service.repo.GetPostById(postId)
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusBadRequest)
		return response
	}
	if post == nil || post.UserId != userId {
		response := entity.NewResponseObject(nil, "invalid user, user is not authorized to delete post", http.StatusUnauthorized)
		return response
	}
	err = service.repo.DeletePostById(postId)
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusInternalServerError)
		return response
	}
	response := entity.NewResponseObject(nil, fmt.Sprintf("Post deleted successfully, PostId: %v", postId), http.StatusOK)
	return response
}

func (service *PostService) UpdatePostViewsById(postId int) entity.Response {
	err := service.repo.UpdatePostViewsById(postId)
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusInternalServerError)
		return response
	}
	response := entity.NewResponseObject(nil, nil, http.StatusOK)
	return response
}

func (service *PostService) GetUserPosts(userId int) entity.Response {
	posts, err := service.repo.GetUserPosts(userId)

	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusInternalServerError)
		return response
	}

	response := entity.NewResponseObject(posts, nil, http.StatusOK)
	return response
}
