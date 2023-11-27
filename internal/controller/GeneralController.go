package controller

import (
	"net/http"
	"sociot/internal/entity"
	"sociot/internal/service"
)

type GeneralController struct {
	service service.GeneralService
}

func NewGeneralController(service service.GeneralService) *GeneralController {
	return &GeneralController{
		service: service,
	}
}

// Search
// @Summary 	Search results whether users or posts
// @Description Get search results for users or posts
// @Tags 		Search
// @Accept 		json
// @Produce 	json
// @Param 		option	query 		string		true	"Searching options" Enums(post, user)
// @Param		search 	query 		string 		true	"Title of post or username"
// @Success 	200 	{object} 	entity.Response 	"List of relevant results"
// @Failure 	400		{object}	entity.Response 	"Bad request"
// @Failure 	500 	{object} 	entity.Response 	"Internal server error"
// @Router 		/search [get]
func (controller *GeneralController) Search(w http.ResponseWriter, r *http.Request) {
	option := r.URL.Query().Get("option")
	search := r.URL.Query().Get("search")

	response := controller.service.Search(option, search)
	entity.ResponseWithJSON(w, response.Meta.StatusCode, response)
}
