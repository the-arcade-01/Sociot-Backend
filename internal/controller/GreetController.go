package controller

import (
	"encoding/json"
	"net/http"
)

func Greet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Hello, World")
}
