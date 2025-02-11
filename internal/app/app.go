package app

import (
	"avito-shop/internal/database"
	"avito-shop/internal/http/handlers/auth_handler"
	"avito-shop/internal/http/handlers/buy_handler"
	"avito-shop/internal/http/middlewares"
	"avito-shop/internal/repository/user_repository"
	"avito-shop/internal/service/auth_service"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type app struct {
}

func New() *app {
	return &app{}
}

func (a *app) Start() {
	db, err := database.New(database.WithConn())
	if err != nil {
		// Ошибка
		return
	}

	userRepository := user_repository.New(db)
	authService := auth_service.New(userRepository)
	authHendler := auth_handler.New(authService)

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)

	router.Post("/api/auth", authHendler.Handle)

	router.With(middlewares.AuthMiddleware).Get("/api/buy/{item}", buy_handler.Handle)
	//router.Post("/api/auth", auth_handler.Handle)
	//router.Post("/api/auth", auth_handler.Handle)
	//router.Post("/api/auth", auth_handler.Handle)

	log.Fatal(http.ListenAndServe(":8080", router))
}
