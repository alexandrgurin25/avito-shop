package wallet_repository

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) SetAmount(ctx context.Context, tx pgx.Tx, usesId int, amount int) error {

	_, err := tx.Exec(
		ctx,
		`UPDATE wallet SET amount = $2 WHERE user_id = $1`,
		usesId,
		amount,
	)

	if err != nil {
		log.Printf("%v", err)
		return fmt.Errorf("repository SetAmount error %v", err)
	}
	return nil
}
