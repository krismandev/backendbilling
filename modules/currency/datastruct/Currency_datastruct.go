package datastruct

import (
	"billingdashboard/core"
)

//LoginRequest is use for clients login
type CurrencyRequest struct {
	ListCurrencyCode   []string            `json:"list_currency_code"`
	CurrencyCode       string              `json:"currency_code"`
	CurrencyName       string              `json:"currency_name"`
	Default            string              `json:"default"`
	LastUpdateUsername string              `json:"last_update_username"`
	LastUpdateDate     string              `json:"last_update_date"`
	Param              core.DataTableParam `json:"param"`
}

type CurrencyDataStruct struct {
	CurrencyCode       string `json:"currency_code"`
	CurrencyName       string `json:"currency_name"`
	Default            string `json:"default"`
	LastUpdateUsername string `json:"last_update_username"`
	LastUpdateDate     string `json:"last_update_date"`
}
