package user_repository

import (
	"avito-shop/internal/entity"
	"context"
	"fmt"
	"log"
)

func (r *Repository) CreateUser(ctx context.Context, username string, passwordHash string) (*entity.User, error) {
	var id int
	err := r.db.QueryRow(
		ctx,
		`INSERT INTO "users" (username, password_hash) VALUES ($1, $2) RETURNING "id"`,
		username,
		passwordHash,
	).Scan(&id)

	if err != nil {
		log.Printf("%v", err)
		return nil, fmt.Errorf("repository user create error %w", err)
	}

	result := &entity.User{
		ID:           id,
		Username:     username,
		PasswordHash: passwordHash,
	}

	return result, nil
}
