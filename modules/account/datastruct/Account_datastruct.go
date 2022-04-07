package datastruct

import (
	"billingdashboard/core"
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
	Param              core.DataTableParam `json:"param"`
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
