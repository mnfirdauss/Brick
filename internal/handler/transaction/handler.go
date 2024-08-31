package http

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	transactionEntity "github.com/mnfirdauss/Brick/internal/entity/transaction"
	transactionUsecase "github.com/mnfirdauss/Brick/internal/usecase/transaction"
)

type transactionHandler struct {
	transactionUsecase transactionUsecase.TransactionUseCase
}

func NewAccountHandler(r *mux.Router, useCase transactionUsecase.TransactionUseCase) {
	handler := &transactionHandler{transactionUsecase: useCase}
	r.HandleFunc("/transaction/{id}", handler.GetTransactionByID).Methods("GET")
	r.HandleFunc("/transaction/transfer", handler.Transfer).Methods("POST")
	r.HandleFunc("/transaction/callback", handler.Callback).Methods("POST")
}

func (h *transactionHandler) GetTransactionByID(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	vars := mux.Vars(r)
	transactionID := vars["id"]

	transaction, err := h.transactionUsecase.GetTransferByID(ctx, transactionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(transaction)
}

func (h *transactionHandler) Transfer(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var req transactionEntity.Transaction

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	transaction, err := h.transactionUsecase.Transfer(ctx, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(transaction)
}

func (h *transactionHandler) Callback(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var req transactionEntity.Transaction

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	transaction, err := h.transactionUsecase.Callback(ctx, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(transaction)
}
