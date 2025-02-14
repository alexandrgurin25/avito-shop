package wallet_repository

import (
	"avito-shop/internal/common"
	"context"
	"fmt"
	"log"
)

func (r *Repository) CreateWallet(ctx context.Context, userID int) error {
	_, err := r.db.Exec(
		ctx,
		`INSERT INTO "wallet"(user_id, amount) VALUES ($1, $2)`,
		userID,
		common.StartBalance,
	)

	if err != nil {
		log.Printf("%v", err)
		return fmt.Errorf("repository wallet create error %v", err)
	}
	return nil
}
