package controller

import (
	"encoding/json"
	"net/http"
	"sociot/internal/entity"
	service "sociot/internal/service"
	"strconv"

	"github.com/go-chi/chi/v5"
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
	entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
}

func (controller *PostController) CreatePost(w http.ResponseWriter, r *http.Request) {
	postBody := new(entity.PostRequestBody)
	if err := json.NewDecoder(r.Body).Decode(&postBody); err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusBadRequest)
		entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
		return
	}
	defer r.Body.Close()
	response := controller.service.CreatePost(postBody)
	entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
}

func (controller *PostController) GetPostById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	postId, err := strconv.Atoi(id)
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusBadRequest)
		entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
		return
	}
	response := controller.service.GetPostById(postId)
	entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
}

func (controller *PostController) UpdatePostById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	postId, err := strconv.Atoi(id)
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusBadRequest)
		entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
		return
	}
	postBody := new(entity.UpdatePostRequestBody)
	if err := json.NewDecoder(r.Body).Decode(&postBody); err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusBadRequest)
		entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
		return
	}
	defer r.Body.Close()

	response := controller.service.UpdatePostById(postId, postBody)
	entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
}

func (controller *PostController) DeletePostById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	postId, err := strconv.Atoi(id)
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusBadRequest)
		entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
		return
	}
	response := controller.service.DeletePostById(postId)
	entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
}

func (controller *PostController) GetCommentsByPostId(w http.ResponseWriter, r *http.Request) {

}

func (controller *PostController) CreateCommentByPostId(w http.ResponseWriter, r *http.Request) {

}
