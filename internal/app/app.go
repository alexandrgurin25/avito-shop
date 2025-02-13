package app

import (
	"avito-shop/internal/database"
	"avito-shop/internal/http/handlers/auth_handler"
	"avito-shop/internal/http/handlers/buy_handler"
	"avito-shop/internal/http/handlers/send_coin_handler"
	"avito-shop/internal/http/middlewares"
	"avito-shop/internal/repository/user_repository"
	"avito-shop/internal/repository/wallet"
	"avito-shop/internal/service/auth_service"
	"avito-shop/internal/service/send_coin_service"
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
	
	walletRepository := wallet.New(db)
	userRepository := user_repository.New(db)

	authService := auth_service.New(userRepository, walletRepository)
	sendService := send_coin_service.New(walletRepository, userRepository)

	authHendler := auth_handler.New(authService)
	sendHendler := send_coin_handler.New(sendService)

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)

	router.Post("/api/auth", authHendler.Handle)

	router.With(middlewares.AuthMiddleware).Get("/api/buy/{item}", buy_handler.Handle)
	router.With(middlewares.AuthMiddleware).Post("/api/sendCoin", sendHendler.Handle)
	//router.Post("/api/auth", auth_handler.Handle)
	//router.Post("/api/auth", auth_handler.Handle)

	log.Fatal(http.ListenAndServe(":8080", router))
}
