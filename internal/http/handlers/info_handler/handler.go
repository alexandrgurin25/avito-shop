package info_handler

import (
	"encoding/json"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId, ok := ctx.Value("userId").(string)


	

	if !ok {
		http.Error(w, "", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(userId)
}
