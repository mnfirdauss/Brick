package usecase

import (
	"context"
	"log"

	transactionEntity "github.com/mnfirdauss/Brick/internal/entity/transaction"
)

func (u *transactionUseCase) Transfer(ctx context.Context, transaction transactionEntity.Transaction) (transactionEntity.Transaction, error) {
	transaction = transactionEntity.NewTransaction(transaction.SourceAccount, transaction.DestinationAccount, transaction.Amount)

	newTransaction, err := u.transactionRepo.SaveTransaction(ctx, &transaction)
	if err != nil {
		return *newTransaction, err
	}

	goCtx := context.WithoutCancel(ctx)
	go func(ctx context.Context, transaction transactionEntity.Transaction) {
		err := u.bankRepo.Transfer(ctx, &transaction)
		if err != nil {
			log.Println("error transfer")
		}
	}(goCtx, *newTransaction)

	return *newTransaction, nil
}
