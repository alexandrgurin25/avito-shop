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
	_, err := r.db.Exec(
		ctx,
		`INSERT INTO coinHistory(fromuser_id int, touser_id int, amount int) VALUES ($1, $2, $3)`,
		toUserID,
		fromUserId,
		amount,
	)

	if err != nil {
		log.Printf("%v", err)
		return fmt.Errorf("repository AddCoinHisory error %v", err)
	}

	return nil
}
