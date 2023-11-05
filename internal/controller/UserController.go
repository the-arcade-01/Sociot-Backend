package controller

import (
	"net/http"
	entity "sociot/internal/entity"
	service "sociot/internal/service"
)

type UserController struct {
	service service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		service: userService,
	}
}

func (controller *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	response := entity.Response{
		Data: []string{"user1", "user2"},
	}
	entity.ResponseWithJSON(w, http.StatusOK, response)
}

func (controller *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	response := entity.Response{
		Data: []string{"user1", "user2"},
	}
	entity.ResponseWithJSON(w, http.StatusOK, response)
}

func (controller *UserController) UpdateUserById(w http.ResponseWriter, r *http.Request) {
	response := entity.Response{
		Data: []string{"user1", "user2"},
	}
	entity.ResponseWithJSON(w, http.StatusOK, response)
}

func (controller *UserController) DeleteUserById(w http.ResponseWriter, r *http.Request) {
	response := entity.Response{
		Data: []string{"user1", "user2"},
	}
	entity.ResponseWithJSON(w, http.StatusOK, response)
}

func (controller *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	response := entity.Response{
		Data: []string{"user1", "user2"},
	}
	entity.ResponseWithJSON(w, http.StatusOK, response)
}

func (controller *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	response := entity.Response{
		Data: []string{"user1", "user2"},
	}
	entity.ResponseWithJSON(w, http.StatusOK, response)
}

func (controller *UserController) GetUserPosts(w http.ResponseWriter, r *http.Request) {

}

func (controller *UserController) GetUserComments(w http.ResponseWriter, r *http.Request) {

}
