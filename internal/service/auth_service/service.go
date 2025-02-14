package auth_service

import (
	"avito-shop/internal/common"
	"avito-shop/internal/entity"
	"avito-shop/internal/repository/user_repository"
	"avito-shop/internal/repository/wallet_repository"
	"context"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo   *user_repository.Repository
	wallet *wallet_repository.Repository
}

func New(repo *user_repository.Repository, wallet *wallet_repository.Repository) *Service {
	return &Service{repo: repo, wallet: wallet}
}

func (s *Service) Auth(ctx context.Context, username string, password string) (*entity.Auth, error) {
	// Поиск пользователя в бд по username
	userAuth, err := s.repo.FindUserByUsername(ctx, nil, username)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	var jwt entity.Auth
	// Проверка существования пользователя
	if userAuth != nil {
		// Если есть -> проверка пароля
		if CheckPasswordHash(password, userAuth.PasswordHash) {
			jwt.AccessToken, err = createToken(userAuth)
			if err != nil {
				return nil, err
			}

		} else {
			// 401 -> Пароль не подошел
			return nil, common.ErrIncorrectPassword
		}

	} else {
		// Если нет пользователя -> создаем пользователя
		password_hash, err := HashPassword(password)
		if err != nil {
			log.Printf("HashPassword could not generare password_hash %v", err)
			return nil, err
		}

		userAuth, err := s.repo.CreateUser(ctx, username, password_hash)
		if err != nil {
			return nil, err
		}

		// Создаем jwt по username(user'a)
		jwt.AccessToken, err = createToken(userAuth)
		if err != nil {
			return nil, err
		}

		// Cоздание кошелька
		err = s.wallet.CreateWallet(ctx, userAuth.ID)
		if err != nil {
			log.Printf("Здесь ошибка%v", err)
			return nil, err
		}

	}

	return &jwt, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func createToken(user *entity.User) (string, error) {
	secretKey := os.Getenv("AUTH_SECRET_KEY")

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"iat": time.Now().Unix(),
	})
	accessToken, err := t.SignedString([]byte(secretKey))

	if err != nil {
		log.Printf("JWT token could not be signed %v", err)
		return "", err
	}

	return accessToken, nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
