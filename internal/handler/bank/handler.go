package http

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mnfirdauss/Brick/api"
	bankEntity "github.com/mnfirdauss/Brick/internal/entity/bank"
	bankUsecase "github.com/mnfirdauss/Brick/internal/usecase/bank"
)

type BankHandler struct {
	bankUsecase bankUsecase.BankUseCase
}

func NewAccountHandler(r *mux.Router, useCase bankUsecase.BankUseCase) {
	handler := &BankHandler{bankUsecase: useCase}
	r.HandleFunc("/validate-account", handler.ValidateAccount).Methods("POST")
	r.HandleFunc("/add-account", handler.AddAccount).Methods("POST")
}

func (h *BankHandler) ValidateAccount(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	w.Header().Set("Content-Type", "application/json")

	var req bankEntity.Account

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		api.WriteError(w, http.StatusBadRequest, err)
		return
	}

	account, err := h.bankUsecase.ValidateAccount(ctx, req)
	if err != nil {
		api.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.NewEncoder(w).Encode(account)
}

func (h *BankHandler) AddAccount(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	w.Header().Set("Content-Type", "application/json")

	var req bankEntity.Account

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		api.WriteError(w, http.StatusBadRequest, err)
		return
	}

	account, err := h.bankUsecase.AddAccount(ctx, req)
	if err != nil {
		api.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.NewEncoder(w).Encode(account)
}
