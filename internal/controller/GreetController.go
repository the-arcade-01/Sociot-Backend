package controller

import (
	"net/http"
	"sociot/internal/entity"
)

func Greet(w http.ResponseWriter, r *http.Request) {
	response := entity.Response{
		Data: "Hello World!!",
	}
	entity.ResponseWithJSON(w, http.StatusOK, response)
}
