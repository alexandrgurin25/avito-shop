package send_coin_service

import (
	"avito-shop/internal/common"
	"avito-shop/internal/repository/coin_hisory_repository"
	"avito-shop/internal/repository/user_repository"
	"avito-shop/internal/repository/wallet_repository"
	"context"
	"log"
)

type Service struct {
	w *wallet_repository.Repository
	r *user_repository.Repository
	c *coin_hisory_repository.Repository
}

func New(w *wallet_repository.Repository, r *user_repository.Repository) *Service {
	return &Service{w: w, r: r}
}

func (s *Service) SendCoin(ctx context.Context, usernameRecipient string, senderUserID int, amount int) error {
	tx, err := s.w.StartTransaction(ctx)
	if err != nil {
		log.Printf("%v", err)
		return err
	}

	sender, err := s.w.GetAmountByUserId(ctx, tx, senderUserID)

	if err != nil {
		// Если произошла ошибка, откатываем транзакцию
		if err := tx.Rollback(ctx); err != nil {
			log.Printf("failed to rollback transaction: %v", err)
		}
		return err
	}

	if sender.Amount < amount {
		return common.ErrLowBalance
	}

	recipient, err := s.r.FindUserByUsername(ctx, tx, usernameRecipient)

	if err != nil {
		// Если произошла ошибка, откатываем транзакцию
		if err := tx.Rollback(ctx); err != nil {
			log.Printf("failed to rollback transaction: %v", err)
		}
		return err
	}

	if recipient == nil {
		return common.ErrUserNotFound
	}

	err = s.w.SetAmount(ctx, tx, recipient.ID, recipient.Amount+amount)
	if err != nil {
		// Если произошла ошибка, откатываем транзакцию
		if err := tx.Rollback(ctx); err != nil {
			log.Printf("failed to rollback transaction: %v", err)
		}
		return err
	}

	err = s.w.SetAmount(ctx, tx, sender.ID, sender.Amount-amount)
	if err != nil {
		// Если произошла ошибка, откатываем транзакцию
		if err := tx.Rollback(ctx); err != nil {
			log.Printf("failed to rollback transaction: %v", err)
		}
		return err
	}

	err = s.c.AddCoinHisory(ctx, tx, recipient.ID, sender.ID, amount)
	if err != nil {
		// Если произошла ошибка, откатываем транзакцию
		if err := tx.Rollback(ctx); err != nil {
			log.Printf("failed to rollback transaction: %v", err)
		}
		return err
	}

	// Если все прошло успешно, фиксируем транзакцию
	if err := tx.Commit(ctx); err != nil {
		log.Printf("failed to commit transaction: %v", err)
		return err
	}

	return nil
}
