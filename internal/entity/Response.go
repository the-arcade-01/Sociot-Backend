package entity

import (
	"encoding/json"
	"net/http"
)

type Meta struct {
	Message    any `json:"message"`
	StatusCode int `json:"statusCode"`
}

type Response struct {
	Data interface{} `json:"data"`
	Meta Meta        `json:"meta"`
}

func NewResponseObject(data interface{}, message any, statusCode int) Response {
	response := Response{
		Data: data,
		Meta: Meta{
			Message:    message,
			StatusCode: statusCode,
		},
	}
	return response
}

func ResponseWithJSON(w http.ResponseWriter, httpStatus int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(payload)
}
