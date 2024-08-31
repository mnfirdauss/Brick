package http

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	bankEntity "github.com/mnfirdauss/Brick/internal/entity/bank"
	bankUsecase "github.com/mnfirdauss/Brick/internal/usecase/bank"
)

type BankHandler struct {
	bankUsecase bankUsecase.BankUseCase
}

func NewAccountHandler(r *mux.Router, useCase bankUsecase.BankUseCase) {
	handler := &BankHandler{bankUsecase: useCase}
	r.HandleFunc("/validate-account", handler.ValidateAccount).Methods("POST")
}

func (h *BankHandler) ValidateAccount(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var req bankEntity.Account

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	account, err := h.bankUsecase.ValidateAccount(ctx, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(account)
}
