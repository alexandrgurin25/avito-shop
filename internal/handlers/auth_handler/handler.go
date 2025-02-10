package auth_handler

import (
	"encoding/json"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()

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

	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}

	json.NewEncoder(w).Encode(in.Username)
}
