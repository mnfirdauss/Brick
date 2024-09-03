package repository

import (
	"context"
	"time"

	"github.com/mnfirdauss/Brick/internal/entity/transaction"
)

func (r *transactiontRepository) UpdateTransaction(ctx context.Context, transaction *transaction.Transaction) (*transaction.Transaction, error) {
	query := `
        UPDATE transactions 
        SET status = $1, updated_at = $2
        WHERE id = $3;
    `
	_, err := r.db.Exec(
		ctx,
		query,
		transaction.Status,
		time.Now(),
		transaction.ID,
	)
	if err != nil {
		return nil, err
	}
	return r.GetTransactionByID(ctx, transaction.ID)
}
