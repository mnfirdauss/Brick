package usecase

import (
	"context"

	transactionEntity "github.com/mnfirdauss/Brick/internal/entity/transaction"
)

func (u *transactionUseCase) Callback(ctx context.Context, transaction transactionEntity.Transaction) (transactionEntity.Transaction, error) {
	existTransaction, err := u.transactionRepo.GetTransactionByID(ctx, transaction.ID)
	if err != nil {
		return transactionEntity.Transaction{}, err
	}

	switch transaction.Status {
	case transactionEntity.TransactionSuccess:
		existTransaction.Status = transactionEntity.TransactionSuccess
	default:
		existTransaction.Status = transactionEntity.TransactionFailed
	}

	existTransaction, err = u.transactionRepo.UpdateTransaction(ctx, existTransaction)
	if err != nil {
		return *existTransaction, err
	}

	return *existTransaction, nil
}
