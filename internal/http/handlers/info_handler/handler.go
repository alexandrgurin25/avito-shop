package info_handler

import (
	"avito-shop/internal/common"
	"avito-shop/internal/service/get_info_service"
	"encoding/json"
	"net/http"
)

type Handler struct {
	service *get_info_service.Service
}

func New(service *get_info_service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := ctx.Value("userId")

	if userId == nil {
		http.Error(w, common.ErrUserIdNotFoundContext.Error(), http.StatusInternalServerError)
		return
	}

	user, err := h.service.GetInfoByUser(ctx, userId.(int))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}
