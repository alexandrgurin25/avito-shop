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
		`select name, count from item join (
		 select o.item_id , count(o.item_id) as count from orders o where o.user_id = $1 group by o.item_id) od on item.id = od.item_id`,
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
