package controller

import (
	"net/http"
	"sociot/internal/entity"
	service "sociot/internal/service"
)

type PostController struct {
	service service.PostService
}

func NewPostController(postService service.PostService) *PostController {
	return &PostController{
		service: postService,
	}
}

func (controller *PostController) GetPosts(w http.ResponseWriter, r *http.Request) {
	response := controller.service.GetPosts()
	entity.ResponseWithJSON(w, http.StatusOK, response)
}

func (controller *PostController) CreatePost(w http.ResponseWriter, r *http.Request) {

}

func (controller *PostController) GetPostById(w http.ResponseWriter, r *http.Request) {

}

func (controller *PostController) UpdatetPostById(w http.ResponseWriter, r *http.Request) {

}

func (controller *PostController) DeletePostById(w http.ResponseWriter, r *http.Request) {

}

func (controller *PostController) GetCommentsByPostId(w http.ResponseWriter, r *http.Request) {

}

func (controller *PostController) CreateCommentByPostId(w http.ResponseWriter, r *http.Request) {

}
