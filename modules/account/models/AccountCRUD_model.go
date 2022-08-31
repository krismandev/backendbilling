package models

import (
	"billingdashboard/connections"
	"billingdashboard/lib"
	"billingdashboard/modules/account/datastruct"
	"strconv"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func GetAccountFromRequest(conn *connections.Connections, req datastruct.AccountRequest) ([]map[string]string, error) {
	var result []map[string]string
	var err error

	// -- THIS IS BASIC GET REQUEST EXAMPLE LOGIC
	var baseWhere string
	var baseParam []interface{}

	lib.AppendWhere(&baseWhere, &baseParam, "account_id = ?", req.AccountID)
	lib.AppendWhere(&baseWhere, &baseParam, "company_id = ?", req.CompanyID)
	lib.AppendWhere(&baseWhere, &baseParam, "billing_type = ?", req.BillingType)
	lib.AppendWhere(&baseWhere, &baseParam, "account_type = ?", req.AccountType)
	lib.AppendWhere(&baseWhere, &baseParam, "status = ?", req.Status)
	lib.AppendWhere(&baseWhere, &baseParam, "name = ?", req.Name)
	if len(req.ListAccountID) > 0 {
		var baseIn string
		for _, prid := range req.ListAccountID {
			lib.AppendComma(&baseIn, &baseParam, "?", prid)
		}
		lib.AppendWhereRaw(&baseWhere, "account_id IN ("+baseIn+")")
	}

	runQuery := "SELECT account_id, name, status, company_id, address1, address2, account_type, billing_type,city, phone, contact_person, contact_person_phone ,account.desc, last_update_username, last_update_date, non_taxable FROM account "
	if len(baseWhere) > 0 {
		runQuery += "WHERE " + baseWhere
	}
	lib.AppendOrderBy(&runQuery, req.Param.OrderBy, req.Param.OrderDir)
	lib.AppendLimit(&runQuery, req.Param.Page, req.Param.PerPage)

	result, _, err = conn.DBAppConn.SelectQueryByFieldNameSlice(runQuery, baseParam...)
	return result, err
}

func InsertAccount(conn *connections.Connections, req datastruct.AccountRequest) error {
	var err error

	// -- THIS IS BASIC INSERT EXAMPLE
	var baseIn string
	var baseParam []interface{}

	lastId, _ := conn.DBAppConn.GetFirstData("SELECT last_id FROM control_id where control_id.key=?", "account")
	// lastId, _ := conn.DBAppConn.GetFirstData("SELECT last_id FROM control_id where control_id.key=?", "account")

	intLastId, err := strconv.Atoi(lastId)
	insertId := intLastId + 1

	insertIdString := strconv.Itoa(insertId)

	lib.AppendComma(&baseIn, &baseParam, "?", insertIdString)
	lib.AppendComma(&baseIn, &baseParam, "?", req.Name)
	lib.AppendComma(&baseIn, &baseParam, "?", req.Status)
	lib.AppendComma(&baseIn, &baseParam, "?", req.CompanyID)
	lib.AppendComma(&baseIn, &baseParam, "?", req.AccountType)
	lib.AppendComma(&baseIn, &baseParam, "?", req.BillingType)
	lib.AppendComma(&baseIn, &baseParam, "?", req.Desc)
	lib.AppendComma(&baseIn, &baseParam, "?", req.Address1)
	lib.AppendComma(&baseIn, &baseParam, "?", req.Address2)
	lib.AppendComma(&baseIn, &baseParam, "?", req.City)
	lib.AppendComma(&baseIn, &baseParam, "?", req.Phone)
	lib.AppendComma(&baseIn, &baseParam, "?", req.ContactPerson)
	lib.AppendComma(&baseIn, &baseParam, "?", req.ContactPersonPhone)

	lib.AppendComma(&baseIn, &baseParam, "?", req.LastUpdateUsername)
	lib.AppendComma(&baseIn, &baseParam, "?", req.NonTaxable)

	qry := "INSERT INTO account (account_id,  name, status, company_id, account_type, billing_type, account.desc, address1,address2,city,phone,contact_person,contact_person_phone, last_update_username,non_taxable) VALUES (" + baseIn + ")"
	_, err = conn.DBAppConn.InsertGetLastID(qry, baseParam...)

	log.Info("InsertParam - ", baseParam)
	_, _, err = conn.DBAppConn.Exec("UPDATE control_id set last_id=? where control_id.key=?", insertIdString, "account")

	return err
}

func UpdateAccount(conn *connections.Connections, req datastruct.AccountRequest) error {
	var err error

	var baseUp string
	var baseParam []interface{}

	lib.AppendComma(&baseUp, &baseParam, "name = ?", req.Name)
	lib.AppendComma(&baseUp, &baseParam, "status = ?", req.Status)
	lib.AppendComma(&baseUp, &baseParam, "account.desc = ?", req.Desc)
	lib.AppendComma(&baseUp, &baseParam, "account_type = ?", req.AccountType)
	lib.AppendComma(&baseUp, &baseParam, "billing_type = ?", req.BillingType)
	lib.AppendComma(&baseUp, &baseParam, "account.desc = ?", req.Desc)
	lib.AppendComma(&baseUp, &baseParam, "address1 = ?", req.Address1)
	lib.AppendComma(&baseUp, &baseParam, "address2 = ?", req.Address2)
	lib.AppendComma(&baseUp, &baseParam, "city = ?", req.City)
	lib.AppendComma(&baseUp, &baseParam, "phone = ? ", req.Phone)
	lib.AppendComma(&baseUp, &baseParam, "contact_person = ?", req.ContactPerson)
	lib.AppendComma(&baseUp, &baseParam, "contact_person_phone = ?", req.ContactPersonPhone)
	lib.AppendComma(&baseUp, &baseParam, "last_update_username = ?", req.LastUpdateUsername)
	lib.AppendComma(&baseUp, &baseParam, "non_taxable = ?", req.NonTaxable)

	qry := "UPDATE account SET " + baseUp + " WHERE account_id = ?"
	baseParam = append(baseParam, req.AccountID)
	_, _, err = conn.DBAppConn.Exec(qry, baseParam...)
	return err
}

func DeleteAccount(conn *connections.Connections, req datastruct.AccountRequest) error {
	var err error
	// -- THIS IS DELETE LOGIC EXAMPLE
	qry := "DELETE FROM account WHERE account_id = ?"
	_, _, err = conn.DBAppConn.Exec(qry, req.AccountID)
	return err
}

func GetRootParentAccountFromRequest(conn *connections.Connections, req datastruct.RootParentAccountRequest) ([]map[string]string, error) {
	var result []map[string]string
	var err error

	// -- THIS IS BASIC GET REQUEST EXAMPLE LOGIC
	// var baseWhere string
	// var baseParam []interface{}

	// lib.AppendWhere(&baseWhere, &baseParam, "account_id = ?", req.AccountID)
	// lib.AppendWhere(&baseWhere, &baseParam, "company_id = ?", req.CompanyID)
	// lib.AppendWhere(&baseWhere, &baseParam, "name = ?", req.Name)

	if len(req.ListAccountID) > 0 {
		// var baseIn string
		// for _, prid := range req.ListAccountID {
		// 	lib.AppendComma(&baseIn, &baseParam, "?", prid)
		// }
		// lib.AppendWhereRaw(&baseWhere, "account_id IN ("+baseIn+")")
		for _, accountId := range req.ListAccountID {
			rootParentAccount, errQry := conn.DBAppConn.GetFirstData("SELECT ocs.getrootparent(?) as root", accountId)
			if errQry != nil {
				return result, errQry
			}

			single := make(map[string]string)
			single["account_id"] = accountId
			single["root_parent_account"] = rootParentAccount
			result = append(result, single)
		}

	}

	// if len(baseWhere) > 0 {
	// 	runQuery += "WHERE " + baseWhere
	// }
	// lib.AppendOrderBy(&runQuery, req.Param.OrderBy, req.Param.OrderDir)
	// lib.AppendLimit(&runQuery, req.Param.Page, req.Param.PerPage)

	// result, _, err = conn.DBAppConn.SelectQueryByFieldNameSlice(runQuery, baseParam...)
	return result, err
}

//Get Root Account by Invoice Type Id
func GetRootAccountFromRequest(conn *connections.Connections, req datastruct.AccountRequest) ([]map[string]string, error) {
	var result []map[string]string
	var err error
	logrus.Info("Test")

	// -- THIS IS BASIC GET REQUEST EXAMPLE LOGIC
	var baseWhereInvoiceType string
	var baseParamInvoiceType []interface{}

	lib.AppendWhere(&baseWhereInvoiceType, &baseParamInvoiceType, "inv_type_id = ?", req.InvoiceTypeID)

	runQueryInvoiceType := "SELECT inv_type_id, inv_type_name, server_id, category, load_from_server FROM invoice_type "
	if len(baseWhereInvoiceType) > 0 {
		runQueryInvoiceType += "WHERE " + baseWhereInvoiceType
	}

	resultInvoiceType, _, err := conn.DBAppConn.SelectQueryByFieldNameSlice(runQueryInvoiceType, baseParamInvoiceType...)

	if len(resultInvoiceType) > 0 {
		invType := resultInvoiceType[0]
		serverId := invType["server_id"]

		var baseWhereServerAccount string
		var baseParamServerAccount []interface{}

		lib.AppendWhere(&baseWhereServerAccount, &baseParamServerAccount, "server_id = ?", serverId)

		runQueryServerAccount := "SELECT account_id, server_id, serveraccount FROM server_account "
		if len(baseWhereServerAccount) > 0 {
			runQueryServerAccount += "WHERE " + baseWhereServerAccount
		}

		resultServerAccount, _, errServerAccount := conn.DBAppConn.SelectQueryByFieldNameSlice(runQueryServerAccount, baseParamServerAccount...)
		if errServerAccount != nil {
			return result, errServerAccount
		}

		var listServerAccountId []string
		for _, eachServerAccount := range resultServerAccount {
			listServerAccountId = append(listServerAccountId, eachServerAccount["serveraccount"])
		}

		lib.UniqueSlice(&listServerAccountId)

		var listRootParentAccount []map[string]string
		if len(listServerAccountId) > 0 {
			for _, accountId := range listServerAccountId {
				accountIdInt, _ := strconv.Atoi(accountId)
				qryGetRootParent := "SELECT "
				if accountIdInt > 70000 {
					qryGetRootParent += " hutch_ocs.getrootparent(?) as root"
				} else {
					qryGetRootParent += " ocs.getrootparent(?) as root"
				}
				rootParentAccount, errQry := conn.DBAppConn.GetFirstData(qryGetRootParent, accountId)
				if errQry != nil {
					return result, errQry
				}

				single := make(map[string]string)
				single["account_id"] = accountId
				single["root_parent_account"] = rootParentAccount
				listRootParentAccount = append(listRootParentAccount, single)
			}
		}

		var listAccountIdBilling []string
		for _, eachRootParentAccount := range listRootParentAccount {
		findFromServerAccount:
			for _, eachServerAccount := range resultServerAccount {
				if eachRootParentAccount["root_parent_account"] == "NULL" && eachRootParentAccount["account_id"] == eachServerAccount["serveraccount"] {
					listAccountIdBilling = append(listAccountIdBilling, eachServerAccount["account_id"])
					break findFromServerAccount
				}
			}
		}

		var baseWhereAccountBilling string
		var baseParamAccountBilling []interface{}
		if len(listAccountIdBilling) > 0 {
			var baseIn string
			for _, prid := range listAccountIdBilling {
				lib.AppendComma(&baseIn, &baseParamAccountBilling, "?", prid)
			}
			lib.AppendWhereRaw(&baseWhereAccountBilling, "account_id IN ("+baseIn+")")
		}

		runQueryAccountBilling := "SELECT account_id, name, status, company_id, address1, address2, account_type, billing_type,city, phone, contact_person, contact_person_phone ,account.desc, last_update_username, last_update_date FROM account "
		if len(baseWhereAccountBilling) > 0 {
			runQueryAccountBilling += "WHERE " + baseWhereAccountBilling
		}
		lib.AppendOrderBy(&runQueryAccountBilling, req.Param.OrderBy, req.Param.OrderDir)
		lib.AppendLimit(&runQueryAccountBilling, req.Param.Page, req.Param.PerPage)

		if len(baseWhereAccountBilling) > 0 {
			result, _, err = conn.DBAppConn.SelectQueryByFieldNameSlice(runQueryAccountBilling, baseParamAccountBilling...)
		}
	}

	return result, err
}

// func GetRootAccountFromRequest(conn *connections.Connections, req datastruct.AccountRequest) ([]map[string]string, error) {
// 	var result []map[string]string
// 	var err error

// 	// -- THIS IS BASIC GET REQUEST EXAMPLE LOGIC
// 	var baseWhere string
// 	var baseParam []interface{}

// 	lib.AppendWhere(&baseWhere, &baseParam, "account_id = ?", req.AccountID)
// 	lib.AppendWhere(&baseWhere, &baseParam, "company_id = ?", req.CompanyID)
// 	lib.AppendWhere(&baseWhere, &baseParam, "name = ?", req.Name)
// 	if len(req.ListAccountID) > 0 {
// 		var baseIn string
// 		for _, prid := range req.ListAccountID {
// 			lib.AppendComma(&baseIn, &baseParam, "?", prid)
// 		}
// 		lib.AppendWhereRaw(&baseWhere, "account_id IN ("+baseIn+")")
// 	}

// 	runQuery := "SELECT account_id, name, status, company_id, address1, address2, account_type, billing_type,city, phone, contact_person, contact_person_phone ,account.desc, last_update_username, last_update_date FROM account "
// 	if len(baseWhere) > 0 {
// 		runQuery += "WHERE " + baseWhere
// 	}
// 	lib.AppendOrderBy(&runQuery, req.Param.OrderBy, req.Param.OrderDir)
// 	lib.AppendLimit(&runQuery, req.Param.Page, req.Param.PerPage)

// 	resultAccountBilling, _, err := conn.DBAppConn.SelectQueryByFieldNameSlice(runQuery, baseParam...)

// 	var baseWhereServerAccount string
// 	var baseParamServerAccount []interface{}

// 	lib.AppendWhere(&baseWhereServerAccount, &baseParamServerAccount, "server_id = ?", req.ServerID)

// 	if len(req.ListAccountID) > 0 {
// 		var baseInServerAccount string
// 		for _, prid := range req.ListAccountID {
// 			lib.AppendComma(&baseInServerAccount, &baseParamServerAccount, "?", prid)
// 		}
// 		lib.AppendWhereRaw(&baseWhereServerAccount, "account_id IN ("+baseInServerAccount+")")
// 	}

// 	runQueryServerAccount := "SELECT account_id, server_id, serveraccount FROM server_account "
// 	if len(baseWhereServerAccount) > 0 {
// 		runQueryServerAccount += "WHERE " + baseWhereServerAccount
// 	}

// 	resultServerAccount, _, err := conn.DBAppConn.SelectQueryByFieldNameSlice(runQueryServerAccount, baseParamServerAccount...)

// 	var listServerAccountId []string
// 	for _, eachServerAccount := range resultServerAccount {
// 		listServerAccountId = append(listServerAccountId, eachServerAccount["serveraccount"])
// 	}

// 	lib.UniqueSlice(&listServerAccountId)

// 	var listRootParentAccount []map[string]string
// 	if len(listServerAccountId) > 0 {
// 		for _, accountId := range listServerAccountId {
// 			rootParentAccount, errQry := conn.DBAppConn.GetFirstData("SELECT ocs.getrootparent(?) as root", accountId)
// 			if errQry != nil {
// 				return result, errQry
// 			}

// 			single := make(map[string]string)
// 			single["account_id"] = accountId
// 			single["root_parent_account"] = rootParentAccount
// 			listRootParentAccount = append(listRootParentAccount, single)
// 		}
// 	}

// 	for _, eachRootParentAccount := range listRootParentAccount {
// 		for _, eachServerAccount := range resultServerAccount {
// 			if eachRootParentAccount["root_parent_account"] == "NULL" && eachRootParentAccount["account_id"] == eachServerAccount["serveraccount"] {
// 				for _, acc := range resultAccountBilling {
// 					if acc["account_id"] == eachServerAccount["account_id"] {
// 						result = append(result, acc)
// 					}
// 				}
// 				// result = append(result,
// 			}
// 		}
// 	}

// 	return result, err
// }
