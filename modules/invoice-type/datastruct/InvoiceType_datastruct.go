package datastruct

import (
	"billingdashboard/core"
)

//LoginRequest is use for clients login
type InvoiceTypeRequest struct {
	ListInvoiceTypeID  []string            `json:"list_invoice-typeid"`
	InvoiceTypeID      string              `json:"invoice_type_id"`
	InvoiceTypeName    string              `json:"invoice_type_name"`
	ServerID           string              `json:"server_id"`
	Category           string              `json:"category"`
	LoadFromServer     string              `json:"load_from_server"`
	LastUpdateUsername string              `json:"last_update_username"`
	LastUpdateDate     string              `json:"last_update_date"`
	Param              core.DataTableParam `json:"param"`
}

type InvoiceTypeDataStruct struct {
	InvoiceTypeID      string `json:"invoice_type_id"`
	InvoiceTypeName    string `json:"invoice_type_name"`
	ServerID           string `json:"server_id"`
	Category           string `json:"category"`
	LoadFromServer     string `json:"load_from_server"`
	LastUpdateUsername string `json:"last_update_username"`
	LastUpdateDate     string `json:"last_update_date"`
}
