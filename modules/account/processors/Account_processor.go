package processors

import (
	"billingdashboard/connections"
	"billingdashboard/modules/account/datastruct"
	"billingdashboard/modules/account/models"
)

func GetListAccount(conn *connections.Connections, req datastruct.AccountRequest) ([]datastruct.AccountDataStruct, error) {
	var output []datastruct.AccountDataStruct
	var err error

	// grab mapping data from model
	accountList, err := models.GetAccountFromRequest(conn, req)
	if err != nil {
		return output, err
	}

	for _, account := range accountList {
		single := CreateSingleAccountStruct(account)
		output = append(output, single)
	}

	return output, err
}

func CreateSingleAccountStruct(account map[string]string) datastruct.AccountDataStruct {
	var single datastruct.AccountDataStruct
	single.AccountID = account["account_id"]
	single.Name = account["name"]
	single.CompanyID = account["company_id"]
	single.Status = account["status"]
	single.Desc = account["desc"]
	single.AccountType = account["account_type"]
	single.BillingType = account["billing_type"]
	single.LastUpdateUsername = account["last_update_username"]
	single.Address1 = account["address1"]
	single.Address2 = account["address2"]
	single.ContactPerson = account["contact_person"]
	single.ContactPersonPhone = account["contact_person_phone"]
	single.Phone = account["phone"]
	single.City = account["city"]
	return single
}

func InsertAccount(conn *connections.Connections, req datastruct.AccountRequest) (datastruct.AccountDataStruct, error) {
	var output datastruct.AccountDataStruct
	var err error

	err = models.InsertAccount(conn, req)
	if err != nil {
		return output, err
	}

	// jika tidak ada error, return single instance of single stub
	single, err := models.GetSingleAccount(req.AccountID, conn, req)
	if err != nil {
		return output, err
	}

	output = CreateSingleAccountStruct(single)
	return output, err
}

func UpdateAccount(conn *connections.Connections, req datastruct.AccountRequest) (datastruct.AccountDataStruct, error) {
	var output datastruct.AccountDataStruct
	var err error

	err = models.UpdateAccount(conn, req)
	if err != nil {
		return output, err
	}

	// jika tidak ada error, return single instance of single stub
	single, err := models.GetSingleAccount(req.AccountID, conn, req)
	if err != nil {
		return output, err
	}

	output = CreateSingleAccountStruct(single)
	return output, err
}

func DeleteAccount(conn *connections.Connections, req datastruct.AccountRequest) error {
	err := models.DeleteAccount(conn, req)
	return err
}
