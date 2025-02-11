package middlewares

import (
	"avito-shop/internal/common"
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type tokenData struct {
	jwt.RegisteredClaims        // техническое поле для пирсинга
	UserId               string `json:"sub"`
	CreatedAt            int64  `json:"iat"`
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
		panic(err)
	}
}

func AuthMiddleware(next http.Handler) http.Handler {
	secretKeyString := os.Getenv("AUTH_SECRET_KEY")

	secretKey := []byte(secretKeyString)

	if secretKey == nil {
		log.Fatal("AUTH_SECRET_KEY not founded")
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessTokenHeader := r.Header.Get("Authorization") // получение данных из заголовка

		if len(accessTokenHeader) == 0 || !(strings.HasPrefix(accessTokenHeader, "Bearer ")) { // проверка, что токен начинается с корректного обозначения типа
			log.Printf("Could not get token %s", accessTokenHeader)
			http.Error(w, "Некорректный jwt", http.StatusBadRequest)
			return
		}

		accessTokenString := accessTokenHeader[7:] // извлечение самой строки токена
		token, err := jwt.ParseWithClaims(accessTokenString, &tokenData{}, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})
		ctx := r.Context()

		if data, ok := token.Claims.(*tokenData); ok && token.Valid {
			expirationTime := time.Now().Add(common.ExpirationTime).Unix()
			if data.CreatedAt > expirationTime {
				log.Print("accessToken timed out")
				http.Error(w, "AccessToken timed out", http.StatusForbidden)
				return
			}
			ctx = context.WithValue(ctx, "userId", data.UserId)
		} else {
			log.Printf("%v", err)
			http.Error(w, "AccessToken timed out", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
