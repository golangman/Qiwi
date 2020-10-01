package qiwi

import (
	"log"
)

// Api client settings
const (
	MaxIdleConnections int    = 20000                   // MaxIdleConnections - connection limit
	RequestTimeout     int    = 25                      // RequestTimeout - requests timeout
	APIURL             string = "https://edge.qiwi.com" // APIURL - main api url
)

// Currency code constants
const (
	CurrencyRUB Currency = 643 // CurrencyRUB - RUB
	CurrencyUSD Currency = 840 // CurrencyUSD - USD
	CurrencyEUR Currency = 978 // CurrencyEUR - EUR
)

// Payment status constants
const (
	PayStatusWaiting Status = "WAITING" // PayStatusWaiting - waiting status
	PayStatusSuccess Status = "SUCCESS" // PayStatusSuccess - success status
	PayStatusError   Status = "ERROR"   // PayStatusError - error status
)

// Provider code constants
const (
	ProviderQiwiWallet    Provider = "99"    // ProviderQiwiWallet - qiwi wallet code
	ProviderVisa          Provider = "1963"  // ProviderVisa - visa  code
	ProviderMastercard    Provider = "21013" // ProviderMastercard - mastercard code
	ProviderVisaCis       Provider = "1960"  // ProviderVisaCis - visa cis code
	ProviderMastercardCis Provider = "21012" // ProviderMastercardCis - mastercard cis code
	ProviderMir           Provider = "31652" // ProviderMir - mir code
	ProviderTinkoff       Provider = "466"   // ProviderTinkoff - tinkoff code
	ProviderAlfabank      Provider = "464"   // ProviderAlfabank - alfabank code
)

// Payment type constants
const (
	OperationAll         HistoryOperation = "ALL"       // HistoryOperationAll - ....
	OperationIn          HistoryOperation = "IN"        // HistoryOperationIn - ....
	OperationOut         HistoryOperation = "OUT"       // HistoryOperationOut - ....
	OperationAllQiwiCard HistoryOperation = "QIWI_CARD" // HistoryOperationAllQiwiCard - ....
)

func init() {
	log.Println("QIWI: DEVELOPED BY BOOST BOTS ( vk.com/boostbots )")
}
