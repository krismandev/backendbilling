package datastruct

import (
	"backendbillingdashboard/core"
	dtServer "backendbillingdashboard/modules/server/datastruct"
)

//LoginRequest is use for clients login
type ServerDataRequest struct {
	ListServerDataID []string `json:"list_server_data_id"`
	MonthUse         string   `json:"month_use"`

	ServerDataID              string              `json:"server_data_id"`
	ServerID                  string              `json:"server_id"`
	ExternalAccountID         string              `json:"external_account_id"`
	ItemID                    string              `json:"item_id"`
	AccountID                 string              `json:"account_id"`
	ExternalRootParentAccount string              `json:"external_rootparent_account"`
	ExternalTransdate         string              `json:"external_transdate"`
	ExternalUserID            string              `json:"external_user_id"`
	ExternalSender            string              `json:"external_sender"`
	ExternalOperatorCode      string              `json:"external_operator_code"`
	ExternalRoute             string              `json:"external_route"`
	ExternalSMSCount          string              `json:"external_smscount"`
	ExternalTransCount        string              `json:"external_transcount"`
	ExternalPrice             string              `json:"external_price"`
	ExternalBalanceType       string              `json:"external_balance_type"`
	InvoiceID                 string              `json:"invoice_id"`
	CurrencyCode              string              `json:"currency_code"`
	Param                     core.DataTableParam `json:"param"`
}

type ServerDataDataStruct struct {
	ServerDataID              string `json:"server_data_id"`
	ServerID                  string `json:"server_id"`
	ExternalAccountID         string `json:"external_account_id"`
	ItemID                    string `json:"item_id"`
	AccountID                 string `json:"account_id"`
	ExternalRootParentAccount string `json:"external_rootparent_account"`
	ExternalTransdate         string `json:"external_transdate"`
	ExternalUserID            string `json:"external_user_id"`
	ExternalSender            string `json:"external_sender"`
	ExternalOperatorCode      string `json:"external_operator_code"`
	ExternalRoute             string `json:"external_route"`
	ExternalSMSCount          string `json:"external_smscount"`
	ExternalTransCount        string `json:"external_transcount"`
	ExternalPrice             string `json:"external_price"`
	ExternalBalanceType       string `json:"external_balance_type"`
	InvoiceID                 string `json:"invoice_id"`

	Item   ItemDataStruct            `json:"item"`
	Server dtServer.ServerDataStruct `json:"server"`
}

type ItemDataStruct struct {
	ItemID   string `json:"item_id"`
	ItemName string `json:"item_name"`
	Operator string `json:"operator"`
	Route    string `json:"route"`
	Category string `json:"category"`
	UOM      string `json:"uom"`

	ItemPrice ItemPriceDataStruct `json:"item_price"`
}

type ItemPriceDataStruct struct {
	ItemID    string `json:"item_id"`
	AccountID string `json:"account_id"`
	ServerID  string `json:"server_id"`
	Price     string `json:"price"`
	Category  string `json:"category"`
}
