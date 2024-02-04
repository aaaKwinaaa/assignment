package integration

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"

	"net/http"
	"os"
	dto "test/integration/transactionDto"
)

type TransactionIntegration interface {
	BoardCastTransaction(ctx context.Context, transaction dto.TransactionRequestDto) (*dto.TransactionResponseDto, error)
	UtilizeTransaction(ctx context.Context, transactionId string) (*dto.UtilizeResponseDto, error)
}

type transactionIntegration struct{}

func NewTransactionIntegration() TransactionIntegration {
	return &transactionIntegration{}
}

func (integration *transactionIntegration) BoardCastTransaction(ctx context.Context, transaction dto.TransactionRequestDto) (*dto.TransactionResponseDto, error) {
	// return &data, err
	jsonData, err := json.Marshal(transaction)
	if err != nil {
		return nil, err
	}

	url := os.Getenv("TRANSACTION_INTEGRATION_URL") + "/broadcast"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response dto.TransactionResponseDto
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil

}

func (integration *transactionIntegration) UtilizeTransaction(ctx context.Context, transactionId string) (*dto.UtilizeResponseDto, error) {
	url := os.Getenv("TRANSACTION_INTEGRATION_URL") + fmt.Sprintf("/check/%s", transactionId)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response dto.UtilizeResponseDto
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil

}
