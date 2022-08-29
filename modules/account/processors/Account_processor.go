package processors

import (
	"billingdashboard/connections"
	"billingdashboard/modules/account/datastruct"
	"billingdashboard/modules/account/models"

	"github.com/sirupsen/logrus"
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
	single.NonTaxable = account["non_taxable"]
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

func GetListRootParentAccount(conn *connections.Connections, req datastruct.RootParentAccountRequest) ([]datastruct.RootParentAccountDataStruct, error) {
	var output []datastruct.RootParentAccountDataStruct
	var err error

	// grab mapping data from model
	rootParentAccountList, err := models.GetRootParentAccountFromRequest(conn, req)
	if err != nil {
		return output, err
	}

	for _, rootParentAccount := range rootParentAccountList {
		single := CreateSingleRootParentAccountStruct(rootParentAccount)
		output = append(output, single)
	}

	return output, err
}

func CreateSingleRootParentAccountStruct(rootParentAccount map[string]string) datastruct.RootParentAccountDataStruct {
	var single datastruct.RootParentAccountDataStruct
	single.AccountID = rootParentAccount["account_id"]
	var rootParentAccountId string
	if rootParentAccount["root_parent_account"] == "NULL" {
		rootParentAccountId = ""
	} else {
		rootParentAccountId = rootParentAccount["root_parent_account"]
	}
	single.RootParentAccount = rootParentAccountId
	return single
}

func GetListRootAccount(conn *connections.Connections, req datastruct.AccountRequest) ([]datastruct.AccountDataStruct, error) {
	var output []datastruct.AccountDataStruct
	var err error

	// grab mapping data from model
	accountList, err := models.GetRootAccountFromRequest(conn, req)
	logrus.Info("LihatProcessor-", accountList)
	if err != nil {
		return output, err
	}

	for _, account := range accountList {
		single := CreateSingleAccountStruct(account)
		output = append(output, single)
	}

	return output, err
}
