package buy_handler

import (
	"avito-shop/internal/database"
	"avito-shop/internal/entity"
	"avito-shop/internal/repository/item_repository"
	"avito-shop/internal/repository/wallet_repository"
	"avito-shop/internal/service/buy_service"
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Handler(t *testing.T) {
	err := os.Setenv("DB_HOST", "127.0.0.1")
	assert.NoError(t, err)
	err = os.Setenv("DB_PORT", "5432")
	assert.NoError(t, err)
	err = os.Setenv("DB_USER", "postgres")
	assert.NoError(t, err)
	err = os.Setenv("DB_PASSWORD", "102104")
	assert.NoError(t, err)
	err = os.Setenv("DB_NAME", "shop")
	assert.NoError(t, err)
	ctx := context.Background()
	db, err := database.New(database.WithConn())
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer func() {
        if err := db.Close(ctx); err != nil {
            t.Errorf("Ошибка при закрытии базы данных: %v", err)
        }
    }()

	user := entity.User{
		ID:           50,
		Username:     "alexandrAvito",
		PasswordHash: "password_hash",
		Amount:       1000,
	}

	_, err = db.Exec(
		ctx,
		`INSERT INTO users(id, username, password_hash) VALUES ($1,$2,$3)`,
		user.ID,
		user.Username,
		user.PasswordHash,
	)
	assert.NoError(t, err)

	_, err = db.Exec(
		ctx,
		`INSERT INTO wallet(user_id, amount) VALUES ($1,$2)`,
		user.ID,
		user.Amount,
	)
	assert.NoError(t, err)

	defer func() {
		_, err = db.Exec(ctx, `DELETE FROM wallet WHERE user_id = $1`, user.ID)
		assert.NoError(t, err)
		_, err = db.Exec(ctx, `DELETE FROM users WHERE id = $1`, user.ID)
		assert.NoError(t, err)
		_, err = db.Exec(ctx, `DELETE FROM orders WHERE user_id = $1`, user.ID)
		assert.NoError(t, err)
	}()

	walletRepository := wallet_repository.New(db)
	if walletRepository == nil {
		t.Fatal("walletRepository is not initialized")
	}
	itemRepository := item_repository.New(db)
	if itemRepository == nil {
		t.Fatal("itemRepository is not initialized")
	}
	buyService := buy_service.New(itemRepository, walletRepository)

	err = buyService.BuyItem(ctx, user.ID, "umbrella")
	assert.NoError(t, err)

	const itemPrice = 200

	// Проверяем, что сумма на кошельке уменьшилась
	var currentAmount int
	err = db.QueryRow(ctx, `SELECT amount FROM wallet WHERE user_id = $1`, user.ID).Scan(&currentAmount)
	assert.NoError(t, err)
	assert.Equal(t, user.Amount-itemPrice, currentAmount, "Wallet amount should be decreased by the item price")

	// Проверяем, что новая запись о покупке появилась
	var purchaseCount int
	err = db.QueryRow(ctx, `SELECT COUNT(*) FROM orders WHERE user_id = $1 AND item_id = $2`, user.ID, 7).Scan(&purchaseCount)
	assert.NoError(t, err)
	assert.Equal(t, 1, purchaseCount, "There should be one purchase record for the item")

}
