package datastruct

import (
	"backendbillingdashboard/core"
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

type BalanceRequest struct {
	BalanceType        string              `json:"balance_type"`
	BalanceName        string              `json:"balance_name"`
	Exponent           string              `json:"exponent"`
	BalanceCategory    string              `json:"balance_category"`
	LastUpdateUsername string              `json:"last_update_username"`
	LastUpdateDate     string              `json:"last_update_date"`
	Param              core.DataTableParam `json:"param"`
}

type BalanceDataStruct struct {
	BalanceType        string `json:"balance_type"`
	BalanceName        string `json:"balance_name"`
	Exponent           string `json:"exponent"`
	BalanceCategory    string `json:"balance_category"`
	LastUpdateUsername string `json:"last_update_username"`
	LastUpdateDate     string `json:"last_update_date"`
}
