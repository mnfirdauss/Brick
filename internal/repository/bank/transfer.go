package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mnfirdauss/Brick/internal/entity/transaction"
)

func (r *bankRepository) Transfer(ctx context.Context, transaction *transaction.Transaction) error {
	var client = &http.Client{}

	url := fmt.Sprintf("%s/bank/transfer", r.baseURL)

	reqBytes, err := json.Marshal(transaction)
	if err != nil {
		return err
	}

	bodyReq := bytes.NewBuffer(reqBytes)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bodyReq)
	if err != nil {
		return err
	}

	response, err := client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to fetch account: %s", response.Status)
	}

	return nil
}
