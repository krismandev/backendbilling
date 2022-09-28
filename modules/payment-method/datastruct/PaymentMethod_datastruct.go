package datastruct

import (
	"backendbillingdashboard/core"
)

//LoginRequest is use for clients login
type PaymentMethodRequest struct {
	ListKey            []string            `json:"list_key"`
	Key                string              `json:"key"`
	PaymentMethodName  string              `json:"payment_method_name"`
	NeedClearingDate   string              `json:"need_clearing_date"`
	NeedCardNumber     string              `json:"need_card_number"`
	BankName           string              `json:"bank_name"`
	Branch             string              `json:"branch"`
	AccountName        string              `json:"account_name"`
	AccountNo          string              `json:"account_no"`
	Code               string              `json:"code"`
	Status             string              `json:"status"`
	PayementType       string              `json:"payment_type"`
	CurrencyCode       string              `json:"currency_code"`
	LastUpdateUsername string              `json:"last_update_username"`
	LastUpdateDate     string              `json:"last_update_date"`
	Param              core.DataTableParam `json:"param"`
}

type PaymentMethodDataStruct struct {
	Key                string `json:"key"`
	PaymentMethodName  string `json:"payment_method_name"`
	NeedClearingDate   string `json:"need_clearing_date"`
	NeedCardNumber     string `json:"need_card_number"`
	BankName           string `json:"bank_name"`
	Branch             string `json:"branch"`
	AccountName        string `json:"account_name"`
	AccountNo          string `json:"account_no"`
	Code               string `json:"code"`
	Status             string `json:"status"`
	PayementType       string `json:"payment_type"`
	CurrencyCode       string `json:"currency_code"`
	LastUpdateUsername string `json:"last_update_username"`
	LastUpdateDate     string `json:"last_update_date"`
}
