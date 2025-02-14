package app

import (
	"avito-shop/internal/database"
	"avito-shop/internal/http/handlers/auth_handler"
	"avito-shop/internal/http/handlers/buy_handler"
	"avito-shop/internal/http/handlers/info_handler"
	"avito-shop/internal/http/handlers/send_coin_handler"
	"avito-shop/internal/http/middlewares"
	"avito-shop/internal/repository/info_repository"
	"avito-shop/internal/repository/item_repository"
	"avito-shop/internal/repository/user_repository"
	"avito-shop/internal/repository/wallet_repository"
	"avito-shop/internal/service/auth_service"
	"avito-shop/internal/service/buy_service"
	"avito-shop/internal/service/get_info_service"
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

	walletRepository := wallet_repository.New(db)
	userRepository := user_repository.New(db)
	infoRepository := info_repository.New(db)

	authService := auth_service.New(userRepository, walletRepository)
	sendService := send_coin_service.New(walletRepository, userRepository)
	infoService := get_info_service.New(infoRepository, walletRepository)

	authHendler := auth_handler.New(authService)
	sendHendler := send_coin_handler.New(sendService)
	infoHendler := info_handler.New(infoService)

	itemRepository := item_repository.New(db)
	buyService := buy_service.New(itemRepository, walletRepository)
	buyHandle := buy_handler.New(buyService)

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)

	router.Post("/api/auth", authHendler.Handle)

	router.With(middlewares.AuthMiddleware).Get("/api/buy/{item}", buyHandle.Handle)
	router.With(middlewares.AuthMiddleware).Post("/api/sendCoin", sendHendler.Handle)
	router.With(middlewares.AuthMiddleware).Get("/api/info", infoHendler.Handle)

	log.Fatal(http.ListenAndServe(":8080", router))
}
