package datastruct

import (
	"billingdashboard/core"
)

//LoginRequest is use for clients login
type PaymentRequest struct {
	ListPaymentID      []string                     `json:"list_paymentid"`
	PaymentID          string                       `json:"payment_id"`
	InvoiceID          string                       `json:"invoice_id"`
	PaymentDate        string                       `json:"payment_date"`
	Total              string                       `json:"total"`
	Note               string                       `json:"note"`
	CreatedBy          string                       `json:"created_by"`
	UserName           string                       `json:"username"`
	PaymentType        string                       `json:"payment_type"`
	ClearingDate       string                       `json:"clearing_date"`
	CardNumber         string                       `json:"card_number"`
	PaymentMethod      string                       `json:"payment_method"`
	LastUpdateUsername string                       `json:"last_update_username"`
	AccountID          string                       `json:"account_id"`
	PaymentDeduction   []PaymentDeductionDataStruct `json:"payment_deductions"`
	Param              core.DataTableParam          `json:"param"`
}

type PaymentDataStruct struct {
	PaymentID        string                       `json:"payment_id"`
	InvoiceID        string                       `json:"invoice_id"`
	PaymentDate      string                       `json:"payment_date"`
	Total            string                       `json:"total"`
	Note             string                       `json:"note"`
	CreatedBy        string                       `json:"created_by"`
	UserName         string                       `json:"username"`
	PaymentType      string                       `json:"payment_type"`
	ClearingDate     string                       `json:"clearing_date"`
	CardNumber       string                       `json:"card_number"`
	PaymentMethod    string                       `json:"payment_method"`
	Status           string                       `json:"status"`
	OcsStatus        string                       `json:"ocs_status"`
	PaymentDeduction []PaymentDeductionDataStruct `json:"payment_deductions"`

	Invoice InvoiceDataStruct `json:"invoice"`
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

	Item ItemDataStruct `json:"item"`
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

type PaymentDeductionDataStruct struct {
	PaymentDeductionTypeID string `json:"payment_deduction_type_id"`
	PaymentID              string `json:"payment_id"`
	Amount                 string `json:"amount"`
	Description            string `json:"description"`
}

type PaymentDeductionTypeRequest struct {
	PaymentDeductionTypeID     string              `json:"payment_deduction_type_id"`
	Description                string              `json:"description"`
	Category                   string              `json:"category"`
	Amount                     string              `json:"amount"`
	LastUpdateUsername         string              `json:"last_update_username"`
	LastUpdateDate             string              `json:"last_update_date"`
	ListPaymentDeductionTypeID []string            `json:"list_payment_deduction_type_id"`
	Param                      core.DataTableParam `json:"param"`
}

type PaymentDeductionTypeDataStruct struct {
	PaymentDeductionTypeID string `json:"payment_deduction_type_id"`
	Description            string `json:"description"`
	Category               string `json:"category"`
	Amount                 string `json:"amount"`
	LastUpdateUsername     string `json:"last_update_username"`
	LastUpdateDate         string `json:"last_update_date"`
}
