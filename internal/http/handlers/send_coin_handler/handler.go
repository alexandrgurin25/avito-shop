package send_coin_handler

import (
	"avito-shop/internal/common"
	"avito-shop/internal/service/send_coin_service"
	"encoding/json"
	"errors"
	"net/http"
)

type Handle struct {
	service *send_coin_service.Service
}

func New(service *send_coin_service.Service) *Handle {
	return &Handle{service: service}
}

func (h *Handle) Handle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var in SendCoin
	err := json.NewDecoder(r.Body).Decode(&in)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	senderUserID := ctx.Value("userId")
	if senderUserID == nil {
		http.Error(w, common.ErrUserIdNotFoundContext.Error(), http.StatusInternalServerError)
		return
	}

	err = h.service.SendCoin(ctx, in.Username, senderUserID.(int), in.Amount)

	if errors.Is(err, common.ErrLowBalance) {
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
