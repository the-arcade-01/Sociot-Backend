package entity

import (
	"encoding/json"
	"net/http"
)

type Meta struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

type Response struct {
	Data interface{} `json:"data"`
	Meta Meta        `json:"meta"`
}

func ResponseWithJSON(w http.ResponseWriter, httpStatus int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(payload)
}
