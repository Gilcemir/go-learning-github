package web

import (
	"encoding/json"
	"net/http"
	"secure-password/internal/validator"
)

type HttpHandler interface {
	CheckPassword(w http.ResponseWriter, r *http.Request)
}

type HttpHandlerImpl struct{}

type RequestBody struct {
	Password string
}

func (h *HttpHandlerImpl) CheckPassword(w http.ResponseWriter, r *http.Request) {
	var password RequestBody
	err := json.NewDecoder(r.Body).Decode(&password)
	if err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	_, valErr := validator.IsValidPassword(&password.Password)

	if valErr != nil {
		http.Error(w, valErr.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
