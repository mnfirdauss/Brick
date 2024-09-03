package usecase

import (
	"context"

	"github.com/mnfirdauss/Brick/internal/entity/bank"
)

func (uc *bankUseCase) AddAccount(ctx context.Context, payload bank.Account) (*bank.Account, error) {
	return uc.repository.SaveAccount(ctx, payload)
}
