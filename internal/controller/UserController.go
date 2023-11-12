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

// GetUsers
// @Summary		Get all users
// @Description Fetches all users
// @Tags		Users
// @Accept		json
// @Produce		json
// @Success		200		{object}	entity.Response		"List of all users"
// @Failure		400		{object}	entity.Response		"Bad request"
// @Failure		500		{object}	entity.Response		"Internal server error"
// @Router		/users [get]
func (controller *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	response := controller.service.GetUsers()
	entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
}

// GetUserById
// @Summary		Get user details by Id
// @Description Get user details by Id
// @Tags		Users
// @Accept		json
// @Produce		json
// @Param		id		path		uint64		true	"User Id"
// @Success		200		{object}	entity.Response		"User details by Id"
// @Failure		400		{object}	entity.Response		"Bad request"
// @Failure		401		{object}	entity.Response		"Unauthorized"
// @Failure		500		{object}	entity.Response		"Internal server error"
// @Router		/users/{id} [get]
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

// UpdateUserById
// @Summary		Update user details by Id
// @Description Update user details by Id
// @Tags		Users
// @Accept		json
// @Produce		json
// @Param		id		path		uint64		true	"User Id"
// @Param		userBody	body	entity.UpdateUserRequestBody	true	"Update user request body"
// @Success		200		{object}	entity.Response		"User update sucess response"
// @Failure		400		{object}	entity.Response		"Bad request"
// @Failure		401		{object}	entity.Response		"Unauthorized"
// @Failure		500		{object}	entity.Response		"Internal server error"
// @Router		/users/{id} [put]
func (controller *UserController) UpdateUserById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusBadRequest)
		entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
		return
	}

	userBody := new(entity.UpdateUserRequestBody)
	if err := json.NewDecoder(r.Body).Decode(&userBody); err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusBadRequest)
		entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
		return
	}
	defer r.Body.Close()

	response := controller.service.UpdateUserById(userId, userBody)
	entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
}

// DeleteUserById
// @Summary		Deletes a user by Id
// @Description Deletes a user by Id
// @Tags		Users
// @Accept		json
// @Produce		json
// @Param		id		path		uint64		true	"User Id"
// @Success		200		{object}	entity.Response		"Deletes a user by Id"
// @Failure		400		{object}	entity.Response		"Bad request"
// @Failure		401		{object}	entity.Response		"Unauthorized"
// @Failure		500		{object}	entity.Response		"Internal server error"
// @Router		/users/{id} [delete]
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

// CreateUser
// @Summary		Creates a new user
// @Description	Creates a new user
// @Tags		Users
// @Accept		json
// @Produce		json
// @Param		userBody	body	entity.CreateUserRequestBody	true	"User request body"
// @Success		200		{object}	entity.Response		"User success response"
// @Failure		400		{object}	entity.Response		"Bad request"
// @Failure		401		{object}	entity.Response		"Unauthorized"
// @Failure		500		{object}	entity.Response		"Internal server error"
// @Router		/users [post]
func (controller *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	userBody := new(entity.CreateUserRequestBody)
	if err := json.NewDecoder(r.Body).Decode(&userBody); err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusBadRequest)
		entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
		return
	}
	response := controller.service.CreateUser(userBody)
	entity.ResponseWithJSON(w, http.StatusOK, response)
}

func (controller *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	userBody := new(entity.LoginUserRequestBody)
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
