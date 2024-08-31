package repository

import (
	"context"
	"encoding/json"

	"github.com/mnfirdauss/Brick/internal/entity/transaction"
)

func (r *transactiontRepository) GetTransactionByID(ctx context.Context, transactionID string) (*transaction.Transaction, error) {
	query := `SELECT id, source_account, destination_account, amount, status, created_at, updated_at FROM transactions WHERE id = $1`

	var transfer transaction.Transaction
	var sourceAccountJSON, destinationAccountJSON []byte

	err := r.db.QueryRow(ctx, query, transactionID).Scan(
		&transfer.ID,
		&sourceAccountJSON,
		&destinationAccountJSON,
		&transfer.Amount,
		&transfer.Status,
		&transfer.CreatedAt,
		&transfer.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(sourceAccountJSON, &transfer.SourceAccount)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(destinationAccountJSON, &transfer.DestinationAccount)
	if err != nil {
		return nil, err
	}

	return &transfer, nil
}
