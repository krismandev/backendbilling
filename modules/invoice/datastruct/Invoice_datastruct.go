package datastruct

import (
	"backendbillingdashboard/core"
	dtServer "backendbillingdashboard/modules/server/datastruct"
)

//LoginRequest is use for clients login
type InvoiceRequest struct {
	ListInvoiceID              []string `json:"list_invoiceid"`
	InvoiceID                  string   `json:"invoice_id"`
	InvoiceNo                  string   `json:"invoice_no"`
	InvoiceDate                string   `json:"invoice_date"`
	InvoiceStatus              string   `json:"invoice_status"`
	AccountID                  string   `json:"account_id"`
	MonthUse                   string   `json:"month_use"`
	InvoiceTypeID              string   `json:"invoice_type_id"`
	PrintCounter               string   `json:"print_counter"`
	Note                       string   `json:"note"`
	CancelDesc                 string   `json:"cancel_desc"`
	LastPrintUsername          string   `json:"last_print_username"`
	LastPrintDate              string   `json:"last_print_date"`
	CreatedAt                  string   `json:"created_at"`
	CreatedBy                  string   `json:"created_by"`
	LastUpdateUsername         string   `json:"last_update_username"`
	LastUpdateDate             string   `json:"last_update_date"`
	DiscountType               string   `json:"discount_type"`
	Discount                   string   `json:"discount"`
	PPN                        string   `json:"ppn"`
	Paid                       string   `json:"paid"`
	PaymentMethod              string   `json:"payment_method"`
	ExchangeRateDate           string   `json:"exchange_rate_date"`
	DueDate                    string   `json:"due_date"`
	ApproachingDueDateInterval string   `json:"approaching_due_date_interval"`
	GrandTotal                 string   `json:"grand_total"`
	PPNAmount                  string   `json:"ppn_amount"`

	ServerID          string                `json:"server_id"`
	ListInvoiceDetail []InvoiceDetailStruct `json:"list_invoice_detail"`

	Param core.DataTableParam `json:"param"`
}

type InvoiceDataStruct struct {
	InvoiceID          string `json:"invoice_id"`
	InvoiceNo          string `json:"invoice_no"`
	InvoiceDate        string `json:"invoice_date"`
	InvoiceStatus      string `json:"invoice_status"`
	AccountID          string `json:"account_id"`
	MonthUse           string `json:"month_use"`
	InvoiceTypeID      string `json:"invoice_type_id"`
	PrintCounter       string `json:"print_counter"`
	Note               string `json:"note"`
	CancelDesc         string `json:"cancel_desc"`
	LastPrintUsername  string `json:"last_print_username"`
	LastPrintDate      string `json:"last_print_date"`
	CreatedAt          string `json:"created_at"`
	CreatedBy          string `json:"created_by"`
	LastUpdateUsername string `json:"last_update_username"`
	LastUpdateDate     string `json:"last_update_date"`
	DiscountType       string `json:"discount_type"`
	Discount           string `json:"discount"`
	PPN                string `json:"ppn"`
	Paid               string `json:"paid"`
	PaymentMethod      string `json:"payment_method"`
	ExchangeRateDate   string `json:"exchange_rate_date"`
	DueDate            string `json:"due_date"`
	GrandTotal         string `json:"grand_total"`
	PPNAmount          string `json:"ppn_amount"`

	InvoiceType InvoiceTypeDataStruct `json:"invoice_type"`
	Account     AccountDataStruct     `json:"account"`

	ListInvoiceDetail []InvoiceDetailStruct `json:"list_invoice_detail"`
}

type InvoiceDetailStruct struct {
	InvoiceDetailID string `json:"invoice_detail_id"`
	InvoiceID       string `json:"invoice_id"`
	ItemID          string `json:"item_id"`
	Qty             string `json:"qty"`
	Uom             string `json:"uom"`
	ItemPrice       string `json:"item_price"`
	Note            string `json:"note"`
	BalanceType     string `json:"balance_type"`
	ServerID        string `json:"server_id"`

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
}

type InvoiceTypeDataStruct struct {
	InvoiceTypeID      string `json:"invoice_type_id"`
	InvoiceTypeName    string `json:"invoice_type_name"`
	ServerID           string `json:"server_id"`
	Category           string `json:"category"`
	LoadFromServer     string `json:"load_from_server"`
	LastUpdateUsername string `json:"last_update_username"`
	LastUpdateDate     string `json:"last_update_date"`
	CurrencyCode       string `json:"currency_code"`
}

type AccountDataStruct struct {
	AccountID          string `json:"account_id"`
	Name               string `json:"name"`
	Status             string `json:"status"`
	CompanyID          string `json:"company_id"`
	AccountType        string `json:"account_type"`
	BillingType        string `json:"billing_type"`
	Desc               string `json:"desc"`
	Address1           string `json:"address1"`
	Address2           string `json:"address2"`
	City               string `json:"city"`
	Phone              string `json:"phone"`
	ContactPerson      string `json:"contact_person"`
	ContactPersonPhone string `json:"contact_person_phone"`
	LastUpdateUsername string `json:"last_update_username"`
}
