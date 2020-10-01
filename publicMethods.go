package qiwi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// NewPersonalAPI - init new api with default http client
func NewPersonalAPI(apiKey string) *PersonalAPI {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: MaxIdleConnections,
		},
		Timeout: time.Duration(RequestTimeout) * time.Second,
	}

	return newPersonalAPIWithHTTPClient(apiKey, client)
}

// GetBalance - return accounts balances
func (qiwi *PersonalAPI) GetBalance() (*Balance, error) {
	resp, err := qiwi.sendRequest(qiwi.apiKey, "GET", "/funding-sources/v1/accounts/current", nil)

	if err != nil {
		return nil, err
	}

	var res Balance

	err = json.Unmarshal(resp, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// GetPaymentsHistory - return payments history
func (qiwi *PersonalAPI) GetPaymentsHistory(wallet string, rows int) (*History, error) {
	resp, err := qiwi.sendRequest(qiwi.apiKey, "GET", fmt.Sprintf("/payment-history/v1/persons/%s/payments?rows=%d", wallet, rows), nil)

	if err != nil {
		return nil, err
	}

	var res History

	err = json.Unmarshal(resp, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
