package buy_handler

import (
	"avito-shop/internal/database"
	"avito-shop/internal/entity"
	"avito-shop/internal/repository/item_repository"
	"avito-shop/internal/repository/wallet_repository"
	"avito-shop/internal/service/buy_service"
	"context"
	"log"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func init() {
	if err := godotenv.Load("../../../../.env"); err != nil {
		log.Print("No .env file found")
		panic(err)
	}
}

func Test_Handler(t *testing.T) {
	ctx := context.Background()
	db, err := database.New(database.WithConn())
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close(ctx)

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

	//сделать select в бд и проверить что списались деньги и появилась новая запись
}
