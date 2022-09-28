package models

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/account/datastruct"
	"errors"
	"strconv"
)

func GetSingleAccount(accountID string, conn *connections.Connections, req datastruct.AccountRequest) (map[string]string, error) {
	var result map[string]string
	var err error

	// -- EXAMPLE
	if len(accountID) == 0 {
		accountID = req.AccountID
	}
	query := "SELECT account_id, name, status, company_id, account_type, billing_type, account.desc, last_update_username, last_update_date FROM account WHERE account_id = ?"
	results, _, err := conn.DBAppConn.SelectQueryByFieldNameSlice(query, accountID)
	if err != nil {
		return result, err
	}

	// convert from []map[string]string to single map[string]string
	for _, res := range results {
		result = res
		break
	}
	return result, err
}

func CheckAccountExists(stubID string, conn *connections.Connections) error {
	// var param []interface{}
	// qry := "SELECT COUNT(stubid) FROM stub WHERE stubid = ?"
	// param = append(param, stubID)

	// cnt, _ := conn.DBAppConn.GetFirstData(qry, param...)
	// datacount, _ := strconv.Atoi(cnt)
	// if datacount == 0 {
	// 	return errors.New("Company ID is not exists")
	// }
	return nil
}

func CheckAccountDuplicate(exceptID string, conn *connections.Connections, req datastruct.AccountRequest) error {
	var param []interface{}
	qry := "SELECT COUNT(account_id) FROM account WHERE account_id = ?"
	param = append(param, req.AccountID)
	if len(exceptID) > 0 {
		qry += " AND account_id <> ?"
		param = append(param, exceptID)
	}

	cnt, _ := conn.DBAppConn.GetFirstData(qry, param...)
	datacount, _ := strconv.Atoi(cnt)
	if datacount > 0 {
		return errors.New("Account ID is already exists. Please use another Account ID")
	}
	return nil
}
