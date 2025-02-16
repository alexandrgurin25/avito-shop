package send_coin_service

import (
	"avito-shop/internal/database"
	"avito-shop/internal/entity"
	"avito-shop/internal/service/send_coin_service/mocks"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestSendCoin_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockWalletRepo := mocks.NewMockwalletRepository(ctrl)
	mockAuthRepo := mocks.NewMockauthRepository(ctrl)
	mockCoinHistoryRepo := mocks.NewMockcoinHistoryRepository(ctrl)

	service := New(mockWalletRepo, mockAuthRepo, mockCoinHistoryRepo)

	ctx := context.Background()
	var pgtx database.PgTxMock

	senderUserID := 1
	recipientUsername := "recipient"
	amount := 10

	// Настройка моков
	mockWalletRepo.EXPECT().StartTransaction(gomock.Any()).AnyTimes().Return(&pgtx, nil)
	mockWalletRepo.EXPECT().GetAmountByUserId(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&entity.User{ID: senderUserID, Amount: 100}, nil)
	mockAuthRepo.EXPECT().FindUserByUsername(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&entity.User{ID: 2, Amount: 50}, nil)
	mockWalletRepo.EXPECT().SetAmount(gomock.Any(), gomock.Any(), 2, 60).AnyTimes().Return(nil)
	mockWalletRepo.EXPECT().SetAmount(gomock.Any(), gomock.Any(), senderUserID, 90).AnyTimes().Return(nil)
	mockCoinHistoryRepo.EXPECT().AddCoinHisory(gomock.Any(), gomock.Any(), senderUserID, 2, amount).AnyTimes().Return(nil)

	// Вызов метода
	err := service.SendCoin(ctx, recipientUsername, senderUserID, amount)

	// Проверка результата
	assert.NoError(t, err)
}

func TestSendCoin_Failure_SetAmountError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockWalletRepo := mocks.NewMockwalletRepository(ctrl)
	mockAuthRepo := mocks.NewMockauthRepository(ctrl)
	mockCoinHistoryRepo := mocks.NewMockcoinHistoryRepository(ctrl)

	service := New(mockWalletRepo, mockAuthRepo, mockCoinHistoryRepo)

	ctx := context.Background()
	var pgtx database.PgTxMock

	senderUserID := 1
	recipientUsername := "recipient"
	amount := 10

	// Настройка моков
	mockWalletRepo.EXPECT().StartTransaction(gomock.Any()).AnyTimes().Return(&pgtx, nil)
	mockWalletRepo.EXPECT().GetAmountByUserId(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&entity.User{ID: senderUserID, Amount: 100}, nil)
	mockAuthRepo.EXPECT().FindUserByUsername(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&entity.User{ID: 2, Amount: 50}, nil)

	// Здесь мы настраиваем мок, чтобы он возвращал ошибку
	mockWalletRepo.EXPECT().SetAmount(gomock.Any(), gomock.Any(), 2, 60).Return(errors.New("failed to set amount"))
	mockWalletRepo.EXPECT().SetAmount(gomock.Any(), gomock.Any(), senderUserID, 90).AnyTimes().Return(nil)
	mockCoinHistoryRepo.EXPECT().AddCoinHisory(gomock.Any(), gomock.Any(), senderUserID, 2, amount).AnyTimes().Return(nil)

	// Вызов метода
	err := service.SendCoin(ctx, recipientUsername, senderUserID, amount)

	// Проверка результата
	assert.Error(t, err)                                 // Ожидаем, что ошибка будет возвращена
	assert.Equal(t, "failed to set amount", err.Error()) // Проверяем, что сообщение об ошибке соответствует ожидаемому
}
