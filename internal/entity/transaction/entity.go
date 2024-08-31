package transaction

import (
	"time"

	"github.com/google/uuid"
	"github.com/mnfirdauss/Brick/internal/entity/bank"
)

type TransactionStatus string

const (
	TransactionPending TransactionStatus = "PENDING"
	TransactionSuccess TransactionStatus = "SUCCESS"
	TransactionFailed  TransactionStatus = "FAILED"
)

type Transaction struct {
	ID                 string            `json:"id"`
	SourceAccount      bank.Account      `json:"source_account"`
	DestinationAccount bank.Account      `json:"destination_account"`
	Amount             float64           `json:"amount"`
	Status             TransactionStatus `json:"status"`
	CreatedAt          time.Time         `json:"created_at"`
	UpdatedAt          time.Time         `json:"updated_at"`
}

func NewTransaction(sourceAccount, destinationAccount bank.Account, amount float64) Transaction {
	return Transaction{
		ID:                 uuid.NewString(),
		SourceAccount:      sourceAccount,
		DestinationAccount: destinationAccount,
		Amount:             amount,
		Status:             TransactionPending,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}
}
