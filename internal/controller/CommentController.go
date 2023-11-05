package controller

import (
	"net/http"
	"sociot/internal/service"
)

type CommentController struct {
	service service.CommentService
}

func NewCommentController(commentService *service.CommentService) *CommentController {
	return &CommentController{
		service: *commentService,
	}
}

func (controller *CommentController) UpdateCommentById(w http.ResponseWriter, r *http.Request) {

}

func (controller *CommentController) DeleteCommentById(w http.ResponseWriter, r *http.Request) {

}
