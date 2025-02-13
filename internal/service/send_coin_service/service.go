package send_coin_service

import (
	"avito-shop/internal/common"
	"avito-shop/internal/repository/user_repository"
	"avito-shop/internal/repository/wallet"
	"context"
	"fmt"
	"log"
)

type Service struct {
	w *wallet.Repository
	r *user_repository.Repository
}

func New(w *wallet.Repository, r *user_repository.Repository) *Service {
	return &Service{w: w, r: r}
}

//service
// Получить сумму из кошелька +
// Сравнение коинов +
// Существование пользователя
// Вычитание текущей суммы у отправителя
// Добавление суммы получателю

func (s *Service) SendCoin(ctx context.Context, usernameRecipient string, amount int) error {
	tx, err := s.w.StartTransaction(ctx)
	if err != nil {
		log.Printf("%v", err)
		return err
	}

	recipientUserID := ctx.Value("userId")
	if recipientUserID == nil {
		return fmt.Errorf("userId not found in context")
	}

	sender, err := s.w.GetAmountByUserId(ctx, tx, recipientUserID.(int))

	if err != nil {
		log.Printf("%v", err)
		return err
	}

	if sender.Amount < amount {
		return common.ErrLowBalance
	}

	recipient, err := s.r.FindUserByUsername(ctx, usernameRecipient)

	if err != nil {
		log.Printf("%v", err)
		return err
	}

	if recipient == nil {
		return common.ErrUserNotFound
	}

	err = s.w.SetAmount(ctx, tx, recipient.ID, recipient.Amount+amount)
	if err != nil {
		// Если произошла ошибка, откатываем транзакцию
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			log.Printf("failed to rollback transaction: %v", rollbackErr)
		}
	}
	err = s.w.SetAmount(ctx, tx, sender.ID, sender.Amount-amount)

	defer func() {
		if err != nil {
			// Если произошла ошибка, откатываем транзакцию
			if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
				log.Printf("failed to rollback transaction: %v", rollbackErr)
			}
		} else {
			// Если все прошло успешно, фиксируем транзакцию
			if commitErr := tx.Commit(ctx); commitErr != nil {
				log.Printf("failed to commit transaction: %v", commitErr)
				err = commitErr
			}
		}
	}()

	return err
}
