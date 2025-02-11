package user_repository

import (
	"avito-shop/internal/entity"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx"
)

func (r Repository) FindUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	var id int
	var password_hash string
	err := r.db.QueryRow(
		ctx,
		`SELECT username, password FROM users WHERE username = $1 RETURNING "id", "password_hash"`,
		username,
	).Scan(&id, &password_hash)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("FindUserByUsername repository error -> %w", err)
	}
	user := &entity.User{
		ID:           id,
		Username:     username,
		PasswordHash: password_hash,
	}

	return user, nil
}
