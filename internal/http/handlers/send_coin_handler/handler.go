package send_coin_handler

import (
	"encoding/json"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	var in SendCoin
	err := json.NewDecoder(r.Body).Decode(&in)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}
