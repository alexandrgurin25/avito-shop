package wallet

import (
	"avito-shop/internal/entity"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) GetAmountByUserId(ctx context.Context, tx pgx.Tx, userID int) (*entity.User, error) {
	var amount int
	err := tx.QueryRow(
		ctx,
		`SELECT amount FROM wallet WHERE user_id = $1`,
		userID,
	).Scan(&amount)

	if err != nil {
		log.Printf("%v", err)
		return nil, fmt.Errorf("repository GetAmountByUserId error %v", err)
	}

	user := &entity.User{
		ID:     userID,
		Amount: amount,
	}

	

	return user, nil
}
