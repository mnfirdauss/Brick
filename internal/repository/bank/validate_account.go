package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/mnfirdauss/Brick/internal/entity/bank"
)

func (r *bankRepository) ValidateAccount(ctx context.Context, account bank.Account) (*bank.Account, error) {
	var client = &http.Client{}

	url := fmt.Sprintf("%s/bank/account", r.baseURL)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("account_number", strconv.FormatInt(int64(account.AccountNumber), 10))
	q.Add("account_name", account.AccountName)
	q.Add("bank_name", account.BankName)
	req.URL.RawQuery = q.Encode()

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch account: %s", response.Status)
	}
	type resp struct {
		Data bank.Account `json:"data"`
	}

	var responses []resp
	err = json.NewDecoder(response.Body).Decode(&responses)
	if err != nil {
		return nil, err
	}

	return &responses[0].Data, nil
}
