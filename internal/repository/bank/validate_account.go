package repository

import (
	"context"
	"encoding/json"
	"errors"
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

	var responses []bank.Account
	err = json.NewDecoder(response.Body).Decode(&responses)
	if err != nil {
		return nil, err
	}

	if len(responses) == 0 {
		return nil, errors.New("account not valid")
	}

	return &responses[0], nil
}
