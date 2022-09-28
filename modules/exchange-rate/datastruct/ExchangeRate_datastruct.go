package datastruct

import (
	"backendbillingdashboard/core"
)

//LoginRequest is use for clients login
type ExchangeRateRequest struct {
	// ListExchangeRateID []string            `json:"list_stubid"`
	Date               string              `json:"date"`
	FromCurrency       string              `json:"from_currency"`
	ToCurrency         string              `json:"to_currency"`
	ConvertValue       string              `json:"convert_value"`
	LastUpdateUsername string              `json:"last_update_username"`
	LastUpdateDate     string              `json:"last_update_date"`
	IntervalDay        string              `json:"interval_day"`
	Param              core.DataTableParam `json:"param"`
}

type ExchangeRateDataStruct struct {
	Date               string `json:"date"`
	FromCurrency       string `json:"from_currency"`
	ToCurrency         string `json:"to_currency"`
	ConvertValue       string `json:"convert_value"`
	LastUpdateUsername string `json:"last_update_username"`
	LastUpdateDate     string `json:"last_update_date"`
}
