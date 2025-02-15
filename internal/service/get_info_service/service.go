package get_info_service

import (
	"avito-shop/internal/entity"
	"avito-shop/internal/repository/info_repository"
	"avito-shop/internal/repository/wallet_repository"
	"context"
	"log"
)

type Service struct {
	repo   *info_repository.Repository
	wallet *wallet_repository.Repository
}

func New(repo *info_repository.Repository, wallet *wallet_repository.Repository) *Service {
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

	coinHistory :=  entity.CoinHistory{
		Received: receivedCoin,
		Sent: sentCoin,
	}


	result := &entity.UserInfo{
		Coins:     user.Amount,
		Inventory: invertoris,
		CoinHistory: coinHistory,
	}

	return result, nil
}
