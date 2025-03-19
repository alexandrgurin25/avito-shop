package app

import (
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
	"avito-shop/pkg/logger"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

func (a *app) routers(db *pgxpool.Pool, log *logger.Logger) *chi.Mux {
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

	router.With(middlewares.RequestIDMiddleware).Post("/api/auth", authHendler.Handle)

	router.With(middlewares.RequestIDMiddleware, middlewares.AuthMiddleware(log)).Get("/api/buy/{item}", buyHandle.Handle)
	router.With(middlewares.RequestIDMiddleware, middlewares.AuthMiddleware(log)).Post("/api/sendCoin", sendHendler.Handle)
	router.With(middlewares.RequestIDMiddleware, middlewares.AuthMiddleware(log)).Get("/api/info", infoHendler.Handle)

	return router
}
