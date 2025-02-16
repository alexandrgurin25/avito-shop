package send_coin_service

import (
	"avito-shop/internal/common"
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

type authRepository interface {
	FindUserByUsername(ctx context.Context, tx pgx.Tx, username string) (*entity.User, error)
	CreateUser(ctx context.Context, username string, passwordHash string) (*entity.User, error)
}

type coinHistoryRepository interface {
	AddCoinHisory(ctx context.Context, tx pgx.Tx, fromUserId int, toUserID int, amount int) error
}

type Service struct {
	w walletRepository
	r authRepository
	c coinHistoryRepository
}

func New(w walletRepository, r authRepository, c coinHistoryRepository) *Service {
	return &Service{w: w, r: r, c: c}
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

	err = s.c.AddCoinHisory(ctx, tx, sender.ID, recipient.ID, amount)
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
