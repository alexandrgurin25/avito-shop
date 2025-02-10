package app

import (
	"avito-shop/internal/handlers/auth_handler"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type app struct {
}

func New() *app {
	return &app{}
}

func (a *app) Start() {
	router := chi.NewRouter()
	router.Post("/api/auth", auth_handler.Handle)
	log.Fatal(http.ListenAndServe(":8080", router))
}
