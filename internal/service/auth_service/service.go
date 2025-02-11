package auth_service

import (
	"avito-shop/internal/entity"
	"avito-shop/internal/repository/user_repository"
	"context"
	"log"
)

type Service struct {
	repo *user_repository.Repository
}

func New(repo *user_repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Auth(ctx context.Context, username string, password string) (*entity.Auth, error) {

	// Поиск пользователя в бд username
	userAuth, err := s.repo.FindUserByUsername(ctx, username)

	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	// Проверка существования пользователя
	if userAuth != nil {
		// Если есть -> проверка пароля

	} else {

	}

	// Создаем jwt по username(user'a)

	// Ecли пароль неверный -> вернуть ошибку, которая обработана в handler

	// Если нет пользователя -> создаем пользователя
	// Создаем jwt по username(user'a)
}
