package controller

import (
	"encoding/json"
	"net/http"
	"sociot/internal/entity"
	service "sociot/internal/service"
	"sociot/internal/utils"
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

// GetPosts
// @Summary 	Get all posts
// @Description Fetches all posts
// @Tags 		Posts
// @Accept 		json
// @Produce 	json
// @Success 	200 	{object} 	entity.Response 	"List of all posts"
// @Failure 	400		{object}	entity.Response 	"Bad request"
// @Failure 	500 	{object} 	entity.Response 	"Internal server error"
// @Router 		/posts [get]
func (controller *PostController) GetPosts(w http.ResponseWriter, r *http.Request) {
	response := controller.service.GetPosts()
	entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
}

// CreatePost
// @Summary		Creates a new post
// @Description	Creates a new post
// @Tags		Posts
// @Accept		json
// @Produce		json
// @Param		Authorization	header	string	true	"Authentication header passed like this Bearer T"
// @Param		postBody	body	entity.PostRequestBody	true	"Post request body"
// @Success		200		{object}	entity.Response		"Post success response"
// @Failure		400		{object}	entity.Response		"Bad request"
// @Failure		401		{object}	entity.Response		"Unauthorized"
// @Failure		500		{object}	entity.Response		"Internal server error"
// @Router		/posts [post]
func (controller *PostController) CreatePost(w http.ResponseWriter, r *http.Request) {
	postBody := new(entity.PostRequestBody)
	if err := json.NewDecoder(r.Body).Decode(&postBody); err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusBadRequest)
		entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
		return
	}
	defer r.Body.Close()

	response, err := utils.ValidateRequestBody(postBody)
	if err != nil {
		entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
		return
	}

	err = utils.ValidateAuthToken(postBody.UserId, r.Context())
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusUnauthorized)
		entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
		return
	}

	response = controller.service.CreatePost(postBody)
	entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
}

// GetPostById
// @Summary		Get post details by Id
// @Description Get post details by Id
// @Tags		Posts
// @Accept		json
// @Produce		json
// @Param		id		path		uint64		true	"Post Id"
// @Success		200		{object}	entity.Response		"Post details by Id"
// @Failure		400		{object}	entity.Response		"Bad request"
// @Failure		500		{object}	entity.Response		"Internal server error"
// @Router		/posts/{id} [get]
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

// UpdatePostById
// @Summary		Get post details by Id
// @Description Get post details by Id
// @Tags		Posts
// @Accept		json
// @Produce		json
// @Param		Authorization	header	string	true	"Authentication header passed like this Bearer T"
// @Param		id		path		uint64		true	"Post Id"
// @Param		postBody	body	entity.UpdatePostRequestBody	true	"Update post request body"
// @Success		200		{object}	entity.Response		"Post update success response"
// @Failure		400		{object}	entity.Response		"Bad request"
// @Failure		401		{object}	entity.Response		"Unauthorized"
// @Failure		500		{object}	entity.Response		"Internal server error"
// @Router		/posts/{id} [put]
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

	response, err := utils.ValidateRequestBody(postBody)
	if err != nil {
		entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
		return
	}

	err = utils.ValidateAuthToken(postBody.UserId, r.Context())
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusUnauthorized)
		entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
		return
	}

	response = controller.service.UpdatePostById(postId, postBody)
	entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
}

// DeletePostById
// @Summary		Deletes a post by Id
// @Description Deletes a post by Id
// @Tags		Posts
// @Accept		json
// @Produce		json
// @Param		Authorization	header	string	true	"Authentication header passed like this Bearer T"
// @Param		id		path		uint64		true	"Post Id"
// @Param		postBody	body	entity.DeletePostRequestBody	true	"Delete post request body"
// @Success		200		{object}	entity.Response		"Deletes a Post by Id"
// @Failure		400		{object}	entity.Response		"Bad request"
// @Failure		401		{object}	entity.Response		"Unauthorized"
// @Failure		500		{object}	entity.Response		"Internal server error"
// @Router		/posts/{id} [delete]
func (controller *PostController) DeletePostById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	postId, err := strconv.Atoi(id)
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusBadRequest)
		entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
		return
	}

	postBody := new(entity.DeletePostRequestBody)
	if err := json.NewDecoder(r.Body).Decode(&postBody); err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusBadRequest)
		entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
		return
	}
	defer r.Body.Close()

	response, err := utils.ValidateRequestBody(postBody)
	if err != nil {
		entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
		return
	}

	err = utils.ValidateAuthToken(postBody.UserId, r.Context())
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusUnauthorized)
		entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
		return
	}

	response = controller.service.DeletePostById(postId, postBody.UserId)
	entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
}

// UpdatePostViewsById
// @Summary		Updates views of a post by Id
// @Description Updates views of a post by Id
// @Tags		Posts
// @Accept		json
// @Produce		json
// @Param		id		path		uint64		true	"Post Id"
// @Success		200		{object}	entity.Response		"Updates views of a Post by Id"
// @Failure		400		{object}	entity.Response		"Bad request"
// @Failure		500		{object}	entity.Response		"Internal server error"
// @Router		/posts/views/{id} [put]
func (controller *PostController) UpdatePostViewsById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	postId, err := strconv.Atoi(id)
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusBadRequest)
		entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
		return
	}

	response := controller.service.UpdatePostViewsById(postId)
	entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
}

// GetPostByUserId
// @Summary		Returns users post
// @Description	Returns users post
// @Tags		Posts
// @Accept		json
// @Produce		json
// @Param		Authorization	header	string	true	"Authentication header passed like this Bearer T"
// @Param		id		path		uint64		true	"User Id"
// @Success		200		{object}	entity.Response		"Login User success response with Token"
// @Failure		400		{object}	entity.Response		"Bad request"
// @Failure		401		{object}	entity.Response		"Unauthorized"
// @Failure		500		{object}	entity.Response		"Internal server error"
// @Router		/posts/users/{id} [get]
func (controller *PostController) GetUserPosts(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusBadRequest)
		entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
		return
	}

	err = utils.ValidateAuthToken(userId, r.Context())
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusUnauthorized)
		entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
		return
	}
	response := controller.service.GetUserPosts(userId)
	entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
}

// GetTags
// @Summary		Returns top trending tags
// @Description	Returns top trending tags
// @Tags		Posts
// @Accept		json
// @Produce		json
// @Success		200		{object}	entity.Response		"Normal response"
// @Failure		500		{object}	entity.Response		"Internal server error"
// @Router		/posts/tags [get]
func (controller *PostController) GetTags(w http.ResponseWriter, r *http.Request) {
	response := controller.service.GetTags()
	entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
}

func (controller *PostController) GetCommentsByPostId(w http.ResponseWriter, r *http.Request) {

}

func (controller *PostController) CreateCommentByPostId(w http.ResponseWriter, r *http.Request) {

}
