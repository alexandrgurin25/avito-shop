package item_repository

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) AddItemByUserId(ctx context.Context, tx pgx.Tx, userID int, item int) error {
	_, err := r.db.Exec(
		ctx,
		`INSERT INTO orders(user_id, item_id) VALUES ($1, $2)`,
		userID,
		item,
	)

	if err != nil {
		log.Printf("%v", err)
		return fmt.Errorf("repository AddItemByUserId error %v", err)
	}

	return nil
}
