package buy_handler

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	var in BuyIn
	in.Item = chi.URLParam(r, "item")
	
}
