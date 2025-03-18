package app

import (
	"avito-shop/internal/config"
	"avito-shop/internal/http/handlers/auth_handler"
	"avito-shop/internal/http/handlers/buy_handler"
	"avito-shop/internal/http/handlers/info_handler"
	"avito-shop/internal/http/handlers/send_coin_handler"
	"avito-shop/internal/http/middlewares"
	"avito-shop/internal/repository/coin_hisory_repository"
	"avito-shop/internal/repository/info_repository"
	"avito-shop/internal/repository/item_repository"
	"avito-shop/internal/repository/user_repository"
	"avito-shop/internal/repository/wallet_repository"
	"avito-shop/internal/service/auth_service"
	"avito-shop/internal/service/buy_service"
	"avito-shop/internal/service/get_info_service"
	"avito-shop/internal/service/send_coin_service"
	postgres "avito-shop/pkg/postgtres"
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

type app struct {
}

func New() *app {
	return &app{}
}

func (a *app) Start() {
	logger, _ := zap.NewProduction()

	ctx := context.Background()

	cfg, err := config.New()
	if err != nil {
		logger.Fatal("unable to load config", zap.Error(err))
		return
	}

	db, err := postgres.New(ctx, cfg)
	if err != nil {
		logger.Fatal("unable to connect db", zap.Error(err))
		return
	}

	logger.Info("Succesful start!")

	walletRepository := wallet_repository.New(db)
	userRepository := user_repository.New(db)
	infoRepository := info_repository.New(db)
	coinRepository := coin_hisory_repository.New(db)

	authService := auth_service.New(userRepository, walletRepository)
	sendService := send_coin_service.New(walletRepository, userRepository, coinRepository)
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

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	logger.Fatal("Ошибка при запуске сервера", zap.Error(srv.ListenAndServe()))
}
