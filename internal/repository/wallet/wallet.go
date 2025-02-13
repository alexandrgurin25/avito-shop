package wallet

import (
	"avito-shop/internal/database"
	"context"

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
