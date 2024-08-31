package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mnfirdauss/Brick/internal/entity/transaction"
)

type TransactionRepository interface {
	SaveTransaction(ctx context.Context, transaction *transaction.Transaction) (*transaction.Transaction, error)
	UpdateTransaction(ctx context.Context, transaction *transaction.Transaction) (*transaction.Transaction, error)
	GetTransactionByID(ctx context.Context, transactionID string) (*transaction.Transaction, error)
}

type transactiontRepository struct {
	db *pgxpool.Pool
}

func NewITransactionRepository(db *pgxpool.Pool) TransactionRepository {
	return &transactiontRepository{db: db}
}
