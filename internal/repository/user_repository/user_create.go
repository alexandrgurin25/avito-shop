package user_repository

import (
	"avito-shop/internal/database"
	"context"
)

type Repository struct {
	db database.DataBase
}

func New(db database.DataBase) *Repository {
	return &Repository{db: db}
}

func (r Repository) createUser(ctx context.Context, email string, passwordHash string) {
	err := r.db.ExecContext(
		ctx,
		`INSERT INTO "users" (email, password_hash) VALUES ($1, $2)`,
		email,
		passwordHash,
	)

	if err != nil {

	}
}
