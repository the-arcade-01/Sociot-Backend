package controller

import (
	"encoding/json"
	"net/http"
	entity "sociot/internal/entity"
	service "sociot/internal/service"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type UserController struct {
	service service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return UserController{
		service: userService,
	}
}

func (controller *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	response := controller.service.GetUsers()
	entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
}

func (controller *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusBadRequest)
		entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
		return
	}

	response := controller.service.GetUserById(userId)
	entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
}

func (controller *UserController) UpdateUserById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusBadRequest)
		entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
		return
	}

	var userBody *entity.UpdateUserRequestBody
	if err := json.NewDecoder(r.Body).Decode(&userBody); err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusBadRequest)
		entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
		return
	}
	defer r.Body.Close()

	response := controller.service.UpdateUserById(userId, userBody)
	entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
}

func (controller *UserController) DeleteUserById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusBadRequest)
		entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
		return
	}
	response := controller.service.DeleteUserById(userId)
	entity.ResponseWithJSON(w, http.StatusOK, response)
}

func (controller *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userBody *entity.CreateUserRequestBody
	if err := json.NewDecoder(r.Body).Decode(&userBody); err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusBadRequest)
		entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
		return
	}
	response := controller.service.CreateUser(userBody)
	entity.ResponseWithJSON(w, http.StatusOK, response)
}

func (controller *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	var userBody *entity.LoginUserRequestBody
	if err := json.NewDecoder(r.Body).Decode(&userBody); err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusBadRequest)
		entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
		return
	}
	response := controller.service.LoginUser(userBody)
	entity.ResponseWithJSON(w, http.StatusOK, response)
}

func (controller *UserController) GetUserPosts(w http.ResponseWriter, r *http.Request) {

}

func (controller *UserController) GetUserComments(w http.ResponseWriter, r *http.Request) {

}
