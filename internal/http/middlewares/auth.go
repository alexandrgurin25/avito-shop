package middlewares

import (
	"avito-shop/internal/common"
	"avito-shop/pkg/logger"
	"context"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

type tokenData struct {
	jwt.RegisteredClaims       // техническое поле для пирсинга
	UserId               int   `json:"id"`
	CreatedAt            int64 `json:"iat"`
}

func AuthMiddleware(log *logger.Logger) func(http.Handler) http.Handler {
	secretKeyString := os.Getenv("AUTH_SECRET_KEY")

	secretKey := []byte(secretKeyString)

	if secretKey == nil {
		log.Fatal(context.Background(), "AUTH_SECRET_KEY not founded")
	}
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			ctx := r.Context()
			accessTokenHeader := r.Header.Get("Authorization") // получение данных из заголовка

			if len(accessTokenHeader) == 0 || !(strings.HasPrefix(accessTokenHeader, "Bearer ")) { // проверка, что токен начинается с корректного обозначения типа
				log.Error(ctx, "Invalid or missing Authorization header",
					zap.String("header", accessTokenHeader),
				)
				http.Error(w, "Некорректный jwt", http.StatusBadRequest)
				return
			}

			accessTokenString := accessTokenHeader[7:] // извлечение самой строки токена
			token, err := jwt.ParseWithClaims(accessTokenString, &tokenData{}, func(token *jwt.Token) (interface{}, error) {
				return secretKey, nil
			})

			// Проверка токена на валидность
			if data, ok := token.Claims.(*tokenData); ok && token.Valid {
				expirationTime := time.Unix(data.CreatedAt, 0).Add(common.ExpirationTime)

				// Проверка токена на срок действия
				if time.Now().After(expirationTime) {
					log.Info(ctx, "Token expired",
						zap.Int64("userId", int64(data.UserId)),
						zap.Time("expirationTime", expirationTime),
						zap.Time("currentTime", time.Now()),
					)

					http.Error(w, "AccessToken timed out", http.StatusUnauthorized)
					return
				}
				// Добавляем userId в контекст
				ctx = context.WithValue(ctx, "userId", data.UserId)
			} else {
				log.Info(ctx, "Invalid token",
					zap.String("token", accessTokenString),
					zap.Error(err),
				)
				http.Error(w, "Некорректный jwt", http.StatusBadRequest)
				return

			}

			next.ServeHTTP(w, r.WithContext(ctx))

		})

	}
}
