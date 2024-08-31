package repository

import (
	"context"

	"github.com/mnfirdauss/Brick/internal/entity/bank"
	"github.com/mnfirdauss/Brick/internal/entity/transaction"
)

type BankRepository interface {
	ValidateAccount(ctx context.Context, account bank.Account) (*bank.Account, error)
	Transfer(ctx context.Context, transaction *transaction.Transaction) error
}

type bankRepository struct {
	baseURL string
}

func NewIAccountRepository(baseURL string) BankRepository {
	return &bankRepository{baseURL: baseURL}
}
