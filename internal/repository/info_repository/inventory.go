package info_repository

import (
	"avito-shop/internal/entity"
	"context"
	"fmt"
	"log"
)

func (repo *Repository) GetInvertoryByUserId(ctx context.Context, userID int) ([]entity.InventoryItem, error) {
	rows, err := repo.db.Query(
		ctx,
		`SELECT  i.name, COUNT(i.name) FROM orders o JOIN item i ON i.id = o.item_id WHERE o.user_id  = $1 group by i.name `,
		userID,
	)

	if err != nil {
		log.Printf("%v", err)
		return nil, fmt.Errorf("repository GetInvertoryByUserId error %v", err)
	}

	defer rows.Close()

	inventoris := make([]entity.InventoryItem, 0)

	for rows.Next() {
		inventory := entity.InventoryItem{}

		err = rows.Scan(
			&inventory.Type,
			&inventory.Quantity,
		)

		if err != nil {
			log.Printf("%v", err)
			return nil, fmt.Errorf("repository GetInvertoryByUserId rows.Scan() error %w", err)
		}

		inventoris = append(inventoris, inventory)
	}

	return inventoris, nil
}
