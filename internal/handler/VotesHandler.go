package controller

import (
	"net/http"
	"sociot/internal/entity"
	"sociot/internal/service"
	"sociot/internal/utils"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

type VotesHandler struct {
	service service.VotesService
}

func NewVotesHandler(service service.VotesService) *VotesHandler {
	return &VotesHandler{
		service: service,
	}
}

// UpdatePostVotesById
// @Summary		Updates votes of a post by Id
// @Description Updates votes of a post by Id
// @Tags		Votes
// @Accept		json
// @Produce		json
// @Param		Authorization	header	string	true	"Authentication header passed like this Bearer T"
// @Param		postId	query		uint64		true	"Post Id"
// @Param		userId	query		uint64		true	"User Id"
// @Param		type	query		string		true	"vote type, u or d" Enums(u, d)
// @Success		200		{object}	entity.Response		"Updates votes of a Post by Id"
// @Failure		400		{object}	entity.Response		"Bad request"
// @Failure		500		{object}	entity.Response		"Internal server error"
// @Router		/votes [put]
func (handler VotesHandler) UpdatePostVotesById(w http.ResponseWriter, r *http.Request) {
	postId := r.URL.Query().Get("postId")
	userId := r.URL.Query().Get("userId")
	voteType := r.URL.Query().Get("type")

	if !(strings.Compare(voteType, "u") == 0 || strings.Compare(voteType, "d") == 0) {
		response := entity.NewResponseObject(nil, "Please provide the correct vote type", http.StatusBadRequest)
		entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
		return
	}

	UserId, err := strconv.Atoi(userId)
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusBadRequest)
		entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
		return
	}

	PostId, err := strconv.Atoi(postId)
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusBadRequest)
		entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
		return
	}

	err = utils.ValidateAuthToken(UserId, r.Context())
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusUnauthorized)
		entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
		return
	}

	response := handler.service.UpdatePostVotesById(PostId, UserId, voteType)
	entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
}

// GetVotesCountById
// @Summary		Get votes of a post by Id
// @Description Get votes of a post by Id
// @Tags		Votes
// @Accept		json
// @Produce		json
// @Param		postId	path		uint64		true	"Post Id"
// @Success		200		{object}	entity.Response		"Get votes of a Post by Id"
// @Failure		400		{object}	entity.Response		"Bad request"
// @Failure		500		{object}	entity.Response		"Internal server error"
// @Router		/votes/{postId} [get]
func (handler VotesHandler) GetVotesCountById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "postId")
	postId, err := strconv.Atoi(id)
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusBadRequest)
		entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
		return
	}
	response := handler.service.GetVotesCountById(postId)
	entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
}

// GetUserVoted
// @Summary		Returns 1, 0, -1 whether user voted or not
// @Description Returns 1, 0, -1 whether user voted or not
// @Tags		Votes
// @Accept		json
// @Produce		json
// @Param		Authorization	header	string	true	"Authentication header passed like this Bearer T"
// @Param		postId	query		uint64		true	"Post Id"
// @Param		userId	query		uint64		true	"User Id"
// @Success		200		{object}	entity.Response		"Get votes of a Post by Id"
// @Failure		400		{object}	entity.Response		"Bad request"
// @Failure		500		{object}	entity.Response		"Internal server error"
// @Router		/votes/status [get]
func (handler VotesHandler) GetUserVoted(w http.ResponseWriter, r *http.Request) {
	postId := r.URL.Query().Get("postId")
	userId := r.URL.Query().Get("userId")

	UserId, err := strconv.Atoi(userId)
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusBadRequest)
		entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
		return
	}

	PostId, err := strconv.Atoi(postId)
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusBadRequest)
		entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
		return
	}

	err = utils.ValidateAuthToken(UserId, r.Context())
	if err != nil {
		response := entity.NewResponseObject(nil, err.Error(), http.StatusUnauthorized)
		entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
		return
	}

	response := handler.service.GetUserVoted(PostId, UserId)
	entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
}
