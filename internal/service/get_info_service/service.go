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

func New(repo *info_repository.Repository) *Service {
	return &Service{repo: repo}
}

//Создать таблицу coinHistory
// id | fromUser_id | toUser_id | amount

func (s Service) GetInfoByUser(ctx context.Context, userID int) (*entity.UserInfo, error) {
	tx, err := s.repo.StartTransaction(ctx)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	//Получить кол-во денег на балансе
	user, err := s.wallet.GetAmountByUserId(ctx, tx, userID)
	if err != nil {
		// Если произошла ошибка, откатываем транзакцию
		if err := tx.Rollback(ctx); err != nil {
			log.Printf("failed to rollback transaction GetAmountByUserId: %v", err)
		}
		return nil, err
	}


}
