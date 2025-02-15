package coin_hisory_repository

import (
	"avito-shop/internal/database"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

type Repository struct {
	db database.DataBase
}

func New(db database.DataBase) *Repository {
	return &Repository{db: db}
}

func (r *Repository) StartTransaction(ctx context.Context) (pgx.Tx, error) {
	tx, err := r.db.Begin(ctx)
	return tx, err
}

func (r *Repository) AddCoinHisory(ctx context.Context, tx pgx.Tx, fromUserId int, toUserID int, amount int) error {
	_, err := tx.Exec(
		ctx,
		`INSERT INTO coinhistory(fromuser_id, touser_id, amount) VALUES ($1, $2, $3)`,
		fromUserId,
		toUserID,
		amount,
	)

	if err != nil {
		log.Printf("%v", err)
		return fmt.Errorf("repository AddCoinHisory error %v", err)
	}

	return nil
}
