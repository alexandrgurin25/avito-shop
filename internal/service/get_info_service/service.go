package get_info_service

import (
	"avito-shop/internal/entity"
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

//go:generate mockgen -destination=mocks/service.go -package=mocks -source=service.go

type walletRepository interface {
	StartTransaction(ctx context.Context) (pgx.Tx, error)
	GetAmountByUserId(ctx context.Context, tx pgx.Tx, userID int) (*entity.User, error)
	SetAmount(ctx context.Context, tx pgx.Tx, usesId int, amount int) error
	CreateWallet(ctx context.Context, userID int) error
}

type infoRepository interface {
	GetInvertoryByUserId(ctx context.Context, userID int) ([]entity.InventoryItem, error)
	GetRecevedCoin(ctx context.Context, userID int) ([]entity.ReceivedCoinTransaction, error)
	GetInfoSendCoin(ctx context.Context, userID int) ([]entity.SentCoinTransaction, error)
}

type Service struct {
	repo   infoRepository
	wallet walletRepository
}

func New(repo infoRepository, wallet walletRepository) *Service {
	return &Service{repo: repo, wallet: wallet}
}

func (s *Service) GetInfoByUser(ctx context.Context, userID int) (*entity.UserInfo, error) {
	//Получить кол-во денег на балансе
	user, err := s.wallet.GetAmountByUserId(ctx, nil, userID)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	invertoris, err := s.repo.GetInvertoryByUserId(ctx, userID)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	receivedCoin, err := s.repo.GetRecevedCoin(ctx, userID)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	sentCoin, err := s.repo.GetInfoSendCoin(ctx, userID)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	coinHistory := entity.CoinHistory{
		Received: receivedCoin,
		Sent:     sentCoin,
	}

	result := &entity.UserInfo{
		Coins:       user.Amount,
		Inventory:   invertoris,
		CoinHistory: coinHistory,
	}

	return result, nil
}
