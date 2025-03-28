// Code generated by MockGen. DO NOT EDIT.
// Source: service.go
//
// Generated by this command:
//
//	mockgen -destination=mocks/service.go -package=mocks -source=service.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	entity "avito-shop/internal/entity"
	context "context"
	reflect "reflect"

	pgx "github.com/jackc/pgx/v5"
	gomock "go.uber.org/mock/gomock"
)

// MockwalletRepository is a mock of walletRepository interface.
type MockwalletRepository struct {
	ctrl     *gomock.Controller
	recorder *MockwalletRepositoryMockRecorder
	isgomock struct{}
}

// MockwalletRepositoryMockRecorder is the mock recorder for MockwalletRepository.
type MockwalletRepositoryMockRecorder struct {
	mock *MockwalletRepository
}

// NewMockwalletRepository creates a new mock instance.
func NewMockwalletRepository(ctrl *gomock.Controller) *MockwalletRepository {
	mock := &MockwalletRepository{ctrl: ctrl}
	mock.recorder = &MockwalletRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockwalletRepository) EXPECT() *MockwalletRepositoryMockRecorder {
	return m.recorder
}

// CreateWallet mocks base method.
func (m *MockwalletRepository) CreateWallet(ctx context.Context, userID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWallet", ctx, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateWallet indicates an expected call of CreateWallet.
func (mr *MockwalletRepositoryMockRecorder) CreateWallet(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWallet", reflect.TypeOf((*MockwalletRepository)(nil).CreateWallet), ctx, userID)
}

// GetAmountByUserId mocks base method.
func (m *MockwalletRepository) GetAmountByUserId(ctx context.Context, tx pgx.Tx, userID int) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAmountByUserId", ctx, tx, userID)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAmountByUserId indicates an expected call of GetAmountByUserId.
func (mr *MockwalletRepositoryMockRecorder) GetAmountByUserId(ctx, tx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAmountByUserId", reflect.TypeOf((*MockwalletRepository)(nil).GetAmountByUserId), ctx, tx, userID)
}

// SetAmount mocks base method.
func (m *MockwalletRepository) SetAmount(ctx context.Context, tx pgx.Tx, usesId, amount int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetAmount", ctx, tx, usesId, amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetAmount indicates an expected call of SetAmount.
func (mr *MockwalletRepositoryMockRecorder) SetAmount(ctx, tx, usesId, amount any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAmount", reflect.TypeOf((*MockwalletRepository)(nil).SetAmount), ctx, tx, usesId, amount)
}

// StartTransaction mocks base method.
func (m *MockwalletRepository) StartTransaction(ctx context.Context) (pgx.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartTransaction", ctx)
	ret0, _ := ret[0].(pgx.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartTransaction indicates an expected call of StartTransaction.
func (mr *MockwalletRepositoryMockRecorder) StartTransaction(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartTransaction", reflect.TypeOf((*MockwalletRepository)(nil).StartTransaction), ctx)
}

// MockinfoRepository is a mock of infoRepository interface.
type MockinfoRepository struct {
	ctrl     *gomock.Controller
	recorder *MockinfoRepositoryMockRecorder
	isgomock struct{}
}

// MockinfoRepositoryMockRecorder is the mock recorder for MockinfoRepository.
type MockinfoRepositoryMockRecorder struct {
	mock *MockinfoRepository
}

// NewMockinfoRepository creates a new mock instance.
func NewMockinfoRepository(ctrl *gomock.Controller) *MockinfoRepository {
	mock := &MockinfoRepository{ctrl: ctrl}
	mock.recorder = &MockinfoRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockinfoRepository) EXPECT() *MockinfoRepositoryMockRecorder {
	return m.recorder
}

// GetInfoSendCoin mocks base method.
func (m *MockinfoRepository) GetInfoSendCoin(ctx context.Context, userID int) ([]entity.SentCoinTransaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInfoSendCoin", ctx, userID)
	ret0, _ := ret[0].([]entity.SentCoinTransaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInfoSendCoin indicates an expected call of GetInfoSendCoin.
func (mr *MockinfoRepositoryMockRecorder) GetInfoSendCoin(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInfoSendCoin", reflect.TypeOf((*MockinfoRepository)(nil).GetInfoSendCoin), ctx, userID)
}

// GetInvertoryByUserId mocks base method.
func (m *MockinfoRepository) GetInvertoryByUserId(ctx context.Context, userID int) ([]entity.InventoryItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInvertoryByUserId", ctx, userID)
	ret0, _ := ret[0].([]entity.InventoryItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInvertoryByUserId indicates an expected call of GetInvertoryByUserId.
func (mr *MockinfoRepositoryMockRecorder) GetInvertoryByUserId(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInvertoryByUserId", reflect.TypeOf((*MockinfoRepository)(nil).GetInvertoryByUserId), ctx, userID)
}

// GetRecevedCoin mocks base method.
func (m *MockinfoRepository) GetRecevedCoin(ctx context.Context, userID int) ([]entity.ReceivedCoinTransaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRecevedCoin", ctx, userID)
	ret0, _ := ret[0].([]entity.ReceivedCoinTransaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecevedCoin indicates an expected call of GetRecevedCoin.
func (mr *MockinfoRepositoryMockRecorder) GetRecevedCoin(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecevedCoin", reflect.TypeOf((*MockinfoRepository)(nil).GetRecevedCoin), ctx, userID)
}
