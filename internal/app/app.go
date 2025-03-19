package app

import (
	"avito-shop/internal/config"
	"avito-shop/pkg/logger"
	postgres "avito-shop/pkg/postgtres"
	"context"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type app struct {
}

func New() *app {
	return &app{}
}

func (a *app) Start() {
	ctx := context.Background()

	log := logger.GetLoggerFromCtx(ctx)

	cfg, err := config.New()
	if err != nil {
		log.Fatal(ctx, "unable to load config", zap.Error(err))
		return
	}

	db, err := postgres.New(ctx, cfg)
	if err != nil {
		log.Fatal(ctx, "unable to connect db", zap.Error(err))
		return
	}

	log.Info(ctx, "Successful start!")

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      a.routers(db, log),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Fatal(ctx, "Ошибка при запуске сервера", zap.Error(srv.ListenAndServe()))
}
