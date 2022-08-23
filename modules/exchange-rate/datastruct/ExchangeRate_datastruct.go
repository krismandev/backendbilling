package datastruct

import (
	"billingdashboard/core"
)

//LoginRequest is use for clients login
type ExchangeRateRequest struct {
	// ListExchangeRateID []string            `json:"list_stubid"`
	Date         string              `json:"date"`
	Currency     string              `json:"currency"`
	FromCurrency string              `json:"from_currency"`
	ToCurrency   string              `json:"to_currency"`
	ConvertValue string              `json:"convert_value"`
	Param        core.DataTableParam `json:"param"`
}

type ExchangeRateDataStruct struct {
	Date         string `json:"date"`
	Currency     string `json:"currency"`
	FromCurrency string `json:"from_currency"`
	ToCurrency   string `json:"to_currency"`
	ConvertValue string `json:"convert_value"`
}
