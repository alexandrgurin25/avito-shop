package info_repository

import (
	"avito-shop/internal/entity"
	"context"
	"fmt"
	"log"
)

func (repo *Repository) GetRecevedCoin(ctx context.Context, userID int) ([]entity.ReceivedCoinTransaction, error) {

	rows, err := repo.db.Query(
		ctx,
		`SELECT u.username, c.amount FROM coinhistory c JOIN users u ON u.id = c.fromuser_id WHERE c.touser_id = $1`,
		userID,
	)

	defer rows.Close()

	if err != nil {
		log.Printf("%v", err)
		return nil, fmt.Errorf("repository GetRecevedCoin error %v", err)
	}

	receivedCoins := make([]entity.ReceivedCoinTransaction, 0)

	for rows.Next() {
		receivedCoin := entity.ReceivedCoinTransaction{}

		err = rows.Scan(
			&receivedCoin.FromUser,
			&receivedCoin.Amount,
		)

		if err != nil {
			log.Printf("%v", err)
			return nil, fmt.Errorf("repository GetRecevedCoin rows.Scan() error %w", err)
		}

		receivedCoins = append(receivedCoins, receivedCoin)
	}

	return receivedCoins, nil
}
