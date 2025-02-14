package info_repository

import (
	"avito-shop/internal/database"
)

type Repository struct {
	db database.DataBase
}

func New(db database.DataBase) *Repository {
	return &Repository{db: db}
}
