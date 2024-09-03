package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mnfirdauss/Brick/internal/entity/bank"
)

func (r *bankRepository) SaveAccount(ctx context.Context, account bank.Account) (*bank.Account, error) {
	var client = &http.Client{}

	url := fmt.Sprintf("%s/bank/account", r.baseURL)

	body, _ := json.Marshal(account)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("failed to fetch account: %s", response.Status)
	}

	var newAccount bank.Account
	err = json.NewDecoder(response.Body).Decode(&newAccount)
	if err != nil {
		return nil, err
	}

	return &newAccount, nil
}
