package auth_handler

import (
	"avito-shop/internal/service/auth_service"
	"encoding/json"
	"net/http"
)

type Handler struct {
	service *auth_service.Service
}

func New(service *auth_service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var in AuthDtoIn

	err := json.NewDecoder(r.Body).Decode(&in)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if in.Password == "" || in.Username == "" {
		http.Error(w, "Username and password is required", http.StatusBadRequest)
		return
	}

	AuthDtoOut, err := h.service.Auth(ctx, in.Username, in.Password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(AuthDtoOut)
}
