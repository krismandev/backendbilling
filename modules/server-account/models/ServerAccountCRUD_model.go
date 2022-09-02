package models

import (
	"billingdashboard/connections"
	"billingdashboard/lib"
	"billingdashboard/modules/server-account/datastruct"
	"strings"

	"github.com/sirupsen/logrus"
)

func GetServerAccountFromRequest(conn *connections.Connections, req datastruct.ServerAccountRequest) ([]map[string]string, error) {
	var result []map[string]string
	var err error

	// -- THIS IS BASIC GET REQUEST EXAMPLE LOGIC
	var baseWhere string
	var baseParam []interface{}

	lib.AppendWhere(&baseWhere, &baseParam, "account_id = ?", req.AccountID)
	lib.AppendWhere(&baseWhere, &baseParam, "server_id = ?", req.ServerID)
	lib.AppendWhere(&baseWhere, &baseParam, "external_account_id = ?", req.ExternalAccountID)
	if len(req.ListAccountID) > 0 {
		var baseIn string
		for _, accId := range req.ListAccountID {
			lib.AppendComma(&baseIn, &baseParam, "?", accId)
		}
		lib.AppendWhereRaw(&baseWhere, "account_id IN ("+baseIn+")")
	}

	runQuery := "SELECT server_id, account_id ,external_account_id, last_update_username, last_update_date FROM server_account "
	if len(baseWhere) > 0 {
		runQuery += "WHERE " + baseWhere
	}
	lib.AppendOrderBy(&runQuery, req.Param.OrderBy, req.Param.OrderDir)
	lib.AppendLimit(&runQuery, req.Param.Page, req.Param.PerPage)

	result, _, err = conn.DBAppConn.SelectQueryByFieldNameSlice(runQuery, baseParam...)
	return result, err
}

func InsertServerAccount(conn *connections.Connections, req datastruct.ServerAccountRequest) error {
	var err error

	// -- THIS IS BASIC INSERT EXAMPLE
	var baseIn string
	var baseParam []interface{}

	if len(req.ListServerAccount) == 0 {
		lib.AppendComma(&baseIn, &baseParam, "?", req.AccountID)
		lib.AppendComma(&baseIn, &baseParam, "?", req.ServerID)
		lib.AppendComma(&baseIn, &baseParam, "?", req.ExternalAccountID)

		qry := "INSERT INTO server_account (account_id, server_id ,external_account_id) VALUES (" + baseIn + ")"
		_, _, err = conn.DBAppConn.Exec(qry, baseParam...)
	} else if len(req.ListServerAccount) > 0 {
		bulkInsertQuery := "INSERT INTO server_account (account_id, server_id, external_account_id) VALUES "
		var paramsBulkInsert []interface{}
		var stringGroup []string
		for _, each := range req.ListServerAccount {
			partquery := "(?, ?, ?)"
			paramsBulkInsert = append(paramsBulkInsert, each.AccountID)
			paramsBulkInsert = append(paramsBulkInsert, each.ServerID)
			paramsBulkInsert = append(paramsBulkInsert, each.ExternalAccountID)
			stringGroup = append(stringGroup, partquery)
		}

		final_query := bulkInsertQuery + strings.Join(stringGroup, ", ")
		_, _, errInsert := conn.DBAppConn.Exec(final_query, paramsBulkInsert...)
		if errInsert != nil {
			logrus.Error("Error: " + errInsert.Error())
			return errInsert
		}

	}

	return err
}

func UpdateServerAccount(conn *connections.Connections, req datastruct.ServerAccountRequest) error {
	var err error

	// -- THIS IS BASIC UPDATE EXAMPLE
	// var baseUp string
	// var baseParam []interface{}

	// lib.AppendComma(&baseUp, &baseParam, "ServerAccountname = ?", req.ServerAccountName)
	// qry := "UPDATE ServerAccount SET " + baseUp + " WHERE ServerAccountid = ?"
	// baseParam = append(baseParam, req.ServerAccountID)
	// _, _, err = conn.DBAppConn.Exec(qry, baseParam...)
	return err
}

func DeleteServerAccount(conn *connections.Connections, req datastruct.ServerAccountRequest) error {
	var err error
	// -- THIS IS DELETE LOGIC EXAMPLE
	qry := "DELETE FROM server_account WHERE account_id = ? and server_id = ?"
	_, _, err = conn.DBAppConn.Exec(qry, req.AccountID, req.ServerID)
	return err
}
