package usecase

import (
	"context"

	"github.com/mnfirdauss/Brick/internal/entity/bank"
	bankRepo "github.com/mnfirdauss/Brick/internal/repository/bank"
)

type BankUseCase interface {
	ValidateAccount(ctx context.Context, payload bank.Account) (*bank.Account, error)
	AddAccount(ctx context.Context, payload bank.Account) (*bank.Account, error)
}

type bankUseCase struct {
	repository bankRepo.BankRepository
}

func NewBankUseCase(repo bankRepo.BankRepository) BankUseCase {
	return &bankUseCase{
		repository: repo,
	}
}
