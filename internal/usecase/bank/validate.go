package usecase

import (
	"context"

	"github.com/mnfirdauss/Brick/internal/entity/bank"
)

func (uc *bankUseCase) ValidateAccount(ctx context.Context, payload bank.Account) (*bank.Account, error) {
	return uc.repository.ValidateAccount(ctx, payload)
}
