package buy_handler

import (
	"avito-shop/internal/common"
	"avito-shop/internal/service/buy_service"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handle struct {
	service *buy_service.Service
}

func New(service *buy_service.Service) *Handle {
	return &Handle{service: service}
}

func (h *Handle) Handle(w http.ResponseWriter, r *http.Request) {
	var in BuyIn
	ctx := r.Context()
	in.Item = chi.URLParam(r, "item")

	userCtxID := ctx.Value("userId")
	if userCtxID == nil {
		http.Error(w, common.ErrUserIdNotFoundContext.Error(), http.StatusInternalServerError)
		return
	}

	err := h.service.BuyItem(ctx, userCtxID.(int), in.Item)

	if errors.Is(err, common.ErrItemNotFound) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
