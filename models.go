package qiwi

import (
	"net/http"
	"time"
)

// Currency - money type
type Currency int

// Status - ...
type Status string

// HistoryOperation - ...
type HistoryOperation string

// PersonalAPI - apiClient
type PersonalAPI struct {
	httpClient *http.Client
	apiKey     string
}

// Provider - payment provider
type Provider string

// BalanceAccount - balance
type BalanceAccount struct {
	Alias   string
	FsAlias string
	Title   string
	Type    struct {
		ID    string
		Title string
	}
	HasBalance bool
	Balance    PaymentAmount
	currency   Currency
}

// Balance - list account balances
type Balance struct {
	Accounts []BalanceAccount
}

// PaymentAmount - currency balance
type PaymentAmount struct {
	Amount   float32
	Currency Currency
}

// TransferError - model transfer error
type TransferError struct {
	Code    string
	Message string
}

// HistoryData - ...
type HistoryData struct {
	TxnID                 int
	PersonID              int
	Date                  time.Time
	ErrorCode             int
	Error                 string
	Status                Status
	Type                  HistoryOperation
	StatusText            string
	TrmTxnID              string
	Account               string
	Sum                   PaymentAmount
	Commission            PaymentAmount
	Total                 PaymentAmount
	Provider              HistoryProvider
	Comment               string
	CurrencyRate          float64
	Extras                interface{}
	ChequeReady           bool
	BankDocumentAvailable bool
	BankDocumentReady     bool
	RepeatPaymentEnabled  bool
}

// HistoryProvider - ...
type HistoryProvider struct {
	ID          int
	ShortName   string
	LongName    string
	LogoURL     string
	Description string
	Keys        string
	SiteURL     string
}

// History - ...
type History struct {
	Data        []HistoryData
	NextTxnID   int
	NextTxnDate time.Time
}
