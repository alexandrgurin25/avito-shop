package buy_service

import (
	"avito-shop/internal/common"
	"avito-shop/internal/repository/item_repository"
	"avito-shop/internal/repository/wallet_repository"
	"context"
	"log"
)

type Service struct {
	repo   *item_repository.Repository
	wallet *wallet_repository.Repository
}

func New(repo *item_repository.Repository, wallet *wallet_repository.Repository) *Service {
	return &Service{repo: repo, wallet: wallet}
}

func (s *Service) BuyItem(ctx context.Context, userID int, itemName string) error {

	// Проверить что пользователь хочет купить существущую вещь
	item, err := s.repo.GetItemByName(ctx, nil, itemName)
	if err != nil {
		log.Printf("%v", err)
		return err
	}

	tx, err := s.wallet.StartTransaction(ctx)
	if err != nil {
		log.Printf("%v", err)
		return err
	}
	//Проверить, что хватает денег у пользователя
	user, err := s.wallet.GetAmountByUserId(ctx, tx, userID)
	if err != nil {
		// Если произошла ошибка, откатываем транзакцию
		if err := tx.Rollback(ctx); err != nil {
			log.Printf("failed to rollback transaction GetAmountByUserId: %v", err)
		}
		return err
	}

	if user.Amount < item.Amount {
		return common.ErrLowBalance
	}

	err = s.repo.AddItemByUserId(ctx, tx, userID, item.ID)
	if err != nil {
		log.Printf("%v", err)
		return err
	}

	err = s.wallet.SetAmount(ctx, tx, userID, user.Amount-item.Amount)

	if err != nil {
		// Если произошла ошибка, откатываем транзакцию
		if err := tx.Rollback(ctx); err != nil {
			log.Printf("failed to rollback transaction SetAmount: %v", err)
		}
		return err
	}

	// Если все прошло успешно, фиксируем транзакцию
	if err := tx.Commit(ctx); err != nil {
		log.Printf("failed to commit transaction: %v", err)
	}

	return nil
}
