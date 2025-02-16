package auth_service

import (
	"avito-shop/internal/common"
	"avito-shop/internal/entity"
	"avito-shop/internal/service/auth_service/mocks"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestAuth_SuccessfulLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthRepo := mocks.NewMockauthRepository(ctrl)
	mockWalletRepo := mocks.NewMockwalletRepository(ctrl)

	service := New(mockAuthRepo, mockWalletRepo)

	username := "testuser"
	password := "testpassword"
	passwordHash, _ := HashPassword(password)
	user := &entity.User{ID: 1, Username: username, PasswordHash: passwordHash}

	// Настройка ожиданий
	mockAuthRepo.EXPECT().FindUserByUsername(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(user, nil)
	mockAuthRepo.EXPECT().CreateUser(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(user, nil)
	mockWalletRepo.EXPECT().CreateWallet(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)

	// Тестирование успешной аутентификации
	auth, err := service.Auth(context.Background(), username, password)
	assert.NoError(t, err)
	assert.NotNil(t, auth)
	assert.NotEmpty(t, auth.AccessToken)
}

func TestAuth_ErrorUserNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthRepo := mocks.NewMockauthRepository(ctrl)
	mockWalletRepo := mocks.NewMockwalletRepository(ctrl)

	service := New(mockAuthRepo, mockWalletRepo)

	username := "testuser"
	password := "testpassword"

	// Пользователь не найден
	mockAuthRepo.EXPECT().FindUserByUsername(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(nil, common.ErrUserNotFound)

	// Тестирование ошибки при отсутствии пользователя
	auth, err := service.Auth(context.Background(), username, password)
	assert.Error(t, err)
	assert.Nil(t, auth)
	assert.Equal(t, common.ErrUserNotFound, err)
}

func TestAuth_ErrorPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthRepo := mocks.NewMockauthRepository(ctrl)
	mockWalletRepo := mocks.NewMockwalletRepository(ctrl)
	service := New(mockAuthRepo, mockWalletRepo)
	username := "testuser"
	password := "testpassword"
	passwordHash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("failed to hash password: %v", err)
	}
	user := &entity.User{ID: 1, Username: username, PasswordHash: passwordHash}

	// Настройка ожиданий
	mockAuthRepo.EXPECT().FindUserByUsername(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(user, nil)
	mockAuthRepo.EXPECT().CheckPasswordHash(user.PasswordHash, gomock.Any()).AnyTimes().Return(false) // Неверный пароль

	err = nil
	// Тестирование ошибки при неверном пароле
	auth, err := service.Auth(context.Background(), username, "wrongpassword")
	assert.Error(t, err)
	assert.Nil(t, auth)
	assert.Equal(t, common.ErrIncorrectPassword, err)

}
