package user_repository

import (
	"avito-shop/internal/entity"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

//Дописать транзакцию

func (r Repository) FindUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	var id int
	var amountCoint int
	var password_hash string
	err := r.db.QueryRow(
		ctx,
		"SELECT u.id, u.password_hash, w.amount FROM users u JOIN wallet w ON w.user_id = u.id WHERE u.username = $1",
		username,
	).Scan(&id, &password_hash, &amountCoint)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		fmt.Println(err)
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
