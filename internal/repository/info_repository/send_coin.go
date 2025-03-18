package info_repository

import (
	"avito-shop/internal/entity"
	"context"
	"fmt"
	"log"
)

func (r *Repository) GetInfoSendCoin(ctx context.Context, userID int) ([]entity.SentCoinTransaction, error) {
	rows, err := r.db.Query(
		ctx,
		`SELECT u.username, c.amount  FROM coinhistory c  JOIN users u ON u.id = c.touser_id where c.fromuser_id = $1`,
		userID,
	)

	defer rows.Close()

	if err != nil {
		log.Printf("%v", err)
		return nil, fmt.Errorf("repository GetInfoSendCoin error %v", err)
	}

	sentCoins := make([]entity.SentCoinTransaction, 0)

	for rows.Next() {
		sentCoin := entity.SentCoinTransaction{}

		err = rows.Scan(
			&sentCoin.ToUser,
			&sentCoin.Amount,
		)

		if err != nil {
			log.Printf("%v", err)
			return nil, fmt.Errorf("repository GetInfoSendCoin rows.Scan() error %w", err)
		}

		sentCoins = append(sentCoins, sentCoin)
	}

	return sentCoins, nil
}
