package repository

import (
	"context"
	"encoding/json"

	"github.com/mnfirdauss/Brick/internal/entity/transaction"
)

func (r *transactiontRepository) SaveTransaction(ctx context.Context, transaction *transaction.Transaction) (*transaction.Transaction, error) {
	query := `
		insert into transactions
			(id, source_account, destination_account, amount, status, created_at, updated_at)
		values
			($1, $2, $3, $4, $5, $6, $7)
	`

	sourceAccountJSON, err := json.Marshal(transaction.SourceAccount)
	if err != nil {
		return nil, err
	}

	destinationAccountJSON, err := json.Marshal(transaction.DestinationAccount)
	if err != nil {
		return nil, err
	}

	_, err = r.db.Exec(
		ctx,
		query,
		transaction.ID,
		sourceAccountJSON,
		destinationAccountJSON,
		transaction.Amount,
		transaction.Status,
		transaction.CreatedAt,
		transaction.UpdatedAt,
	)

	return r.GetTransactionByID(ctx, transaction.ID)
}
