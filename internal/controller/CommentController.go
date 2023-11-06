package controller

import (
	"net/http"
	"sociot/internal/entity"
	"sociot/internal/service"
)

type CommentController struct {
	service service.CommentService
}

func NewCommentController(commentService service.CommentService) CommentController {
	return CommentController{
		service: commentService,
	}
}

func (controller *CommentController) GetCommentById(w http.ResponseWriter, r *http.Request) {
	response := controller.service.GetCommentById()
	entity.ResponseWithJSON(w, http.StatusOK, response)
}

func (controller *CommentController) UpdateCommentById(w http.ResponseWriter, r *http.Request) {

}

func (controller *CommentController) DeleteCommentById(w http.ResponseWriter, r *http.Request) {

}
