package get_info_service_test

import (
	"avito-shop/internal/entity"
	"avito-shop/internal/service/get_info_service"
	"avito-shop/internal/service/get_info_service/mocks"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetInfoByUser_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockWalletRepo := mocks.NewMockwalletRepository(ctrl)
	mockInfoRepo := mocks.NewMockinfoRepository(ctrl)

	service := get_info_service.New(mockInfoRepo, mockWalletRepo)

	ctx := context.Background()
	userID := 1

	// Настройка моков
	mockWalletRepo.EXPECT().GetAmountByUserId(ctx, gomock.Any(), userID).Return(&entity.User{ID: userID, Amount: 100}, nil)
	mockInfoRepo.EXPECT().GetInvertoryByUserId(ctx, userID).Return([]entity.InventoryItem{{Type: "Item1", Quantity: 1}}, nil)
	mockInfoRepo.EXPECT().GetRecevedCoin(ctx, userID).Return([]entity.ReceivedCoinTransaction{{FromUser: "alex", Amount: 10}}, nil)
	mockInfoRepo.EXPECT().GetInfoSendCoin(ctx, userID).Return([]entity.SentCoinTransaction{{ToUser: "Avito", Amount: 25}}, nil)

	// Вызов метода
	result, err := service.GetInfoByUser(ctx, userID)

	// Проверка результата
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 100, result.Coins)
	assert.Len(t, result.Inventory, 1)
	assert.Equal(t, "Item1", result.Inventory[0].Type)
	assert.Len(t, result.CoinHistory.Received, 1)
	assert.Equal(t, 10, result.CoinHistory.Received[0].Amount)
	assert.Len(t, result.CoinHistory.Sent, 1)
	assert.Equal(t, 25, result.CoinHistory.Sent[0].Amount)
}

func TestGetInfoByUser_ErrorGettingAmount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockWalletRepo := mocks.NewMockwalletRepository(ctrl)
	mockInfoRepo := mocks.NewMockinfoRepository(ctrl)

	service := get_info_service.New(mockInfoRepo, mockWalletRepo)

	ctx := context.Background()
	userID := 1

	// Настройка моков для возврата ошибки
	mockWalletRepo.EXPECT().GetAmountByUserId(ctx, gomock.Any(), userID).Return(nil, errors.New("failed to get amount"))

	// Вызов метода
	result, err := service.GetInfoByUser(ctx, userID)

	// Проверка результата
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "failed to get amount", err.Error())
}

// Добавьте дополнительные тесты для других случаев ошибок, если необходимо
