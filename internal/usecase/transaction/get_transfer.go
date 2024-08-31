package usecase

import (
	"context"

	transactionEntity "github.com/mnfirdauss/Brick/internal/entity/transaction"
)

func (u *transactionUseCase) GetTransferByID(ctx context.Context, transactionID string) (transactionEntity.Transaction, error) {
	transaction, err := u.transactionRepo.GetTransactionByID(ctx, transactionID)
	if err != nil {
		return transactionEntity.Transaction{}, err
	}

	return *transaction, nil
}
