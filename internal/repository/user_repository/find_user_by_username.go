package user_repository

import (
	"avito-shop/internal/entity"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) FindUserByUsername(ctx context.Context, tx pgx.Tx, username string) (*entity.User, error) {
	var err error

	var id int
	var amountCoint int
	var password_hash string

	db := r.db
	if tx != nil {
		db = tx
	}

	err = db.QueryRow(
		ctx,
		"SELECT u.id, u.password_hash, w.amount FROM users u JOIN wallet w ON w.user_id = u.id WHERE u.username = $1",
		username,
	).Scan(&id, &password_hash, &amountCoint)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("FindUserByUsername repository error -> %w", err)
	}
	user := &entity.User{
		ID:           id,
		Amount:       amountCoint,
		Username:     username,
		PasswordHash: password_hash,
	}

	return user, nil
}
