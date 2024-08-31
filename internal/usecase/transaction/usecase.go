package usecase

import (
	"context"

	transactionEntity "github.com/mnfirdauss/Brick/internal/entity/transaction"
	bankRepo "github.com/mnfirdauss/Brick/internal/repository/bank"
	transactionRepo "github.com/mnfirdauss/Brick/internal/repository/transaction"
)

type TransactionUseCase interface {
	GetTransferByID(ctx context.Context, transactionID string) (transactionEntity.Transaction, error)
	Transfer(ctx context.Context, transaction transactionEntity.Transaction) (transactionEntity.Transaction, error)
	Callback(ctx context.Context, transaction transactionEntity.Transaction) (transactionEntity.Transaction, error)
}

type transactionUseCase struct {
	transactionRepo transactionRepo.TransactionRepository
	bankRepo        bankRepo.BankRepository
}

func NewBanktUseCase(transactionRepo transactionRepo.TransactionRepository, bankRepo bankRepo.BankRepository) TransactionUseCase {
	return &transactionUseCase{
		transactionRepo: transactionRepo,
		bankRepo:        bankRepo,
	}
}
