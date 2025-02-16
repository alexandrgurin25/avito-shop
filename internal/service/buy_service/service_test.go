package buy_service

import (
	"avito-shop/internal/common"
	"avito-shop/internal/database"
	"avito-shop/internal/entity"
	"avito-shop/internal/service/buy_service/mocks"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestBuyItem_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockItemRepo := mocks.NewMockitemRepository(ctrl)
	mockWalletRepo := mocks.NewMockwalletRepository(ctrl)

	service := New(mockItemRepo, mockWalletRepo)

	ctx := context.Background()
	userID := 1
	itemName := "TestItem"
	item := &entity.Item{ID: 1, Amount: 100}
	user := &entity.User{ID: userID, Amount: 200}

	var a database.PgTxMock
	// Настройка ожиданий
	mockItemRepo.EXPECT().GetItemByName(gomock.Any(), nil, itemName).Return(item, nil)
	mockWalletRepo.EXPECT().StartTransaction(gomock.Any()).Return(&a, nil)
	mockWalletRepo.EXPECT().GetAmountByUserId(gomock.Any(), gomock.Any(), userID).Return(user, nil)
	mockItemRepo.EXPECT().AddItemByUserId(gomock.Any(), gomock.Any(), userID, item.ID).Return(nil)
	mockWalletRepo.EXPECT().SetAmount(gomock.Any(), gomock.Any(), userID, user.Amount-item.Amount).Return(nil)

	err := service.BuyItem(ctx, userID, itemName)
	assert.NoError(t, err)
}

func TestBuyItem_Fail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockItemRepo := mocks.NewMockitemRepository(ctrl)
	mockWalletRepo := mocks.NewMockwalletRepository(ctrl)

	service := New(mockItemRepo, mockWalletRepo)

	ctx := context.Background()
	userID := 1
	itemName := "TestItem"
	item := &entity.Item{ID: 1, Amount: 100}
	user := &entity.User{ID: userID, Amount: 50}

	var a database.PgTxMock
	// Настройка ожиданий
	mockItemRepo.EXPECT().GetItemByName(gomock.Any(), nil, itemName).AnyTimes().Return(item, nil)
	mockWalletRepo.EXPECT().StartTransaction(gomock.Any()).AnyTimes().Return(&a, nil)
	mockWalletRepo.EXPECT().GetAmountByUserId(gomock.Any(), gomock.Any(), userID).AnyTimes().Return(user, nil)
	mockItemRepo.EXPECT().AddItemByUserId(gomock.Any(), gomock.Any(), userID, item.ID).AnyTimes().Return(nil)
	mockWalletRepo.EXPECT().SetAmount(gomock.Any(), gomock.Any(), userID, user.Amount-item.Amount).AnyTimes().Return(nil)

	err := service.BuyItem(ctx, userID, itemName)
	assert.ErrorIs(t, err, common.ErrLowBalance)
}
