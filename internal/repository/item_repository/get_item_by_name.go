package item_repository

import (
	"avito-shop/internal/common"
	"avito-shop/internal/entity"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) GetItemByName(ctx context.Context, tx pgx.Tx, nameItem string) (*entity.Item, error) {
	var amount, id int
	err := tx.QueryRow(
		ctx,
		`SELECT id, amount FROM item WHERE name = $1`,
		nameItem,
	).Scan(&id, &amount)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, common.ErrItemNotFound
		}
		return nil, fmt.Errorf("GetItemByName repository error -> %w", err)
	}

	user := &entity.Item{
		ID:     id,
		Name:   nameItem,
		Amount: amount,
	}

	return user, nil
}
