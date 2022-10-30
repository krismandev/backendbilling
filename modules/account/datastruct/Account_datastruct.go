package datastruct

import (
	"backendbillingdashboard/core"
)

//LoginRequest is use for clients login
type AccountRequest struct {
	ListAccountID      []string            `json:"list_account_id"`
	AccountID          string              `json:"account_id"`
	Name               string              `json:"name"`
	Status             string              `json:"status"`
	CompanyID          string              `json:"company_id"`
	AccountType        string              `json:"account_type"`
	BillingType        string              `json:"billing_type"`
	Desc               string              `json:"desc"`
	Address1           string              `json:"address1"`
	Address2           string              `json:"address2"`
	City               string              `json:"city"`
	Phone              string              `json:"phone"`
	ContactPerson      string              `json:"contact_person"`
	ContactPersonPhone string              `json:"contact_person_phone"`
	LastUpdateUsername string              `json:"last_update_username"`
	InvoiceTypeID      string              `json:"invoice_type_id"`
	NonTaxable         string              `json:"non_taxable"`
	TermOfPayment      string              `json:"term_of_payment"`
	Param              core.DataTableParam `json:"param"`
}

type AccountDataStruct struct {
	AccountID          string                     `json:"account_id"`
	Name               string                     `json:"name"`
	Status             string                     `json:"status"`
	CompanyID          string                     `json:"company_id"`
	AccountType        string                     `json:"account_type"`
	BillingType        string                     `json:"billing_type"`
	Desc               string                     `json:"desc"`
	Address1           string                     `json:"address1"`
	Address2           string                     `json:"address2"`
	City               string                     `json:"city"`
	Phone              string                     `json:"phone"`
	ContactPerson      string                     `json:"contact_person"`
	NonTaxable         string                     `json:"non_taxable"`
	TermOfPayment      string                     `json:"term_of_payment"`
	BalanceList        []AccountBalanceDataStruct `json:"balance_list"`
	ContactPersonPhone string                     `json:"contact_person_phone"`
	LastUpdateUsername string                     `json:"last_update_username"`
}

type RootParentAccountRequest struct {
	ListAccountID     []string            `json:"list_account_id"`
	AccountID         string              `json:"account_id"`
	RootParentAccount string              `json:"root_parent_account"`
	Param             core.DataTableParam `json:"param"`
}

type RootParentAccountDataStruct struct {
	AccountID         string `json:"account_id"`
	RootParentAccount string `json:"root_parent_account"`
}

type AccountBalanceDataStruct struct {
	AccountId      string `json:"accountid"`
	BalanceType    string `json:"balance_type"`
	CurrentBalance string `json:"balance"`
	Limit          string `json:"limit"`
	ExpiryDate     string `json:"expiry_date"`
}

type AccountJsonResponse struct {
	TransId      string                   `json:"trans_id"`
	ResponseCode int                      `json:"response_code"`
	ResponseDesc string                   `json:"response_desc"`
	IPAddr       string                   `json:"ip_addr"`
	ErrorDetail  string                   `json:"error_detail"`
	ResponseData AccountDataTableResponse `json:"data,omitempty"`
}

type AccountDataTableResponse struct {
	TotalData int64               `json:"total_data"`
	PerPage   int                 `json:"per_page"`
	Page      int                 `json:"page"`
	TotalPage int                 `json:"total_page"`
	Lists     []AccountDataStruct `json:"list,omitempty"`
}
