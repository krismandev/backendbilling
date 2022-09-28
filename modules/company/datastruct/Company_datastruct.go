package datastruct

import (
	"backendbillingdashboard/core"
)

//LoginRequest is use for clients login
type CompanyRequest struct {
	ListCompanyID      []string            `json:"list_company_id"`
	CompanyID          string              `json:"company_id"`
	Name               string              `json:"name"`
	Status             string              `json:"status"`
	Addr1              string              `json:"addr1"`
	Addr2              string              `json:"addr2"`
	City               string              `json:"city"`
	Country            string              `json:"country"`
	ContactPerson      string              `json:"contact_person"`
	ContactPersonPhone string              `json:"contact_person_phone"`
	Phone              string              `json:"phone"`
	Fax                string              `json:"fax"`
	Desc               string              `json:"desc"`
	LastUpdateUsername string              `json:"last_update_username"`
	Param              core.DataTableParam `json:"param"`
}

type CompanyDataStruct struct {
	CompanyID          string `json:"company_id"`
	Name               string `json:"name"`
	Status             string `json:"status"`
	Addr1              string `json:"addr1"`
	Addr2              string `json:"addr2"`
	City               string `json:"city"`
	Country            string `json:"country"`
	ContactPerson      string `json:"contact_person"`
	ContactPersonPhone string `json:"contact_person_phone"`
	Phone              string `json:"phone"`
	Fax                string `json:"fax"`
	Desc               string `json:"desc"`
	LastUpdateUsername string `json:"last_update_username"`
	LastUpdateDate     string `json:"last_update_date"`
}
