package models

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/lib"
	"backendbillingdashboard/modules/currency/datastruct"
)

func GetCurrencyFromRequest(conn *connections.Connections, req datastruct.CurrencyRequest) ([]map[string]string, error) {
	var result []map[string]string
	var err error

	// -- THIS IS BASIC GET REQUEST EXAMPLE LOGIC
	var baseWhere string
	var baseParam []interface{}

	lib.AppendWhere(&baseWhere, &baseParam, "currency_code = ?", req.CurrencyCode)
	lib.AppendWhere(&baseWhere, &baseParam, "currency_name = ?", req.CurrencyName)
	if len(req.ListCurrencyCode) > 0 {
		var baseIn string
		for _, prid := range req.ListCurrencyCode {
			lib.AppendComma(&baseIn, &baseParam, "?", prid)
		}
		lib.AppendWhereRaw(&baseWhere, "currency_code IN ("+baseIn+")")
	}

	runQuery := "SELECT currency_code, currency_name, currency.default, last_update_username, last_update_date FROM currency "
	if len(baseWhere) > 0 {
		runQuery += "WHERE " + baseWhere
	}
	lib.AppendOrderBy(&runQuery, req.Param.OrderBy, req.Param.OrderDir)
	lib.AppendLimit(&runQuery, req.Param.Page, req.Param.PerPage)

	result, _, err = conn.DBAppConn.SelectQueryByFieldNameSlice(runQuery, baseParam...)
	return result, err
}

func InsertCurrency(conn *connections.Connections, req datastruct.CurrencyRequest) error {
	var err error

	// -- THIS IS BASIC INSERT EXAMPLE
	// var baseIn string
	// var baseParam []interface{}

	// lib.AppendComma(&baseIn, &baseParam, "?", req.CurrencyID)
	// lib.AppendComma(&baseIn, &baseParam, "?", req.CurrencyName)

	// qry := "INSERT INTO currency (currencyid, currencyname) VALUES (" + baseIn + ")"
	// _, _, err = conn.DBAppConn.Exec(qry, baseParam...)

	return err
}

func UpdateCurrency(conn *connections.Connections, req datastruct.CurrencyRequest) error {
	var err error

	// -- THIS IS BASIC UPDATE EXAMPLE
	// var baseUp string
	// var baseParam []interface{}

	// lib.AppendComma(&baseUp, &baseParam, "currencyname = ?", req.CurrencyName)
	// qry := "UPDATE currency SET " + baseUp + " WHERE currencyid = ?"
	// baseParam = append(baseParam, req.CurrencyID)
	// _, _, err = conn.DBAppConn.Exec(qry, baseParam...)
	return err
}

func DeleteCurrency(conn *connections.Connections, req datastruct.CurrencyRequest) error {
	var err error
	// -- THIS IS DELETE LOGIC EXAMPLE
	// qry := "DELETE FROM currency WHERE currencyid = ?"
	// _, _, err = conn.DBAppConn.Exec(qry, req.CurrencyID)
	return err
}

func GetBalanceFromRequest(conn *connections.Connections, req datastruct.BalanceRequest) ([]map[string]string, error) {
	var result []map[string]string
	var err error

	// -- THIS IS BASIC GET REQUEST EXAMPLE LOGIC
	var baseWhere string
	var baseParam []interface{}

	lib.AppendWhere(&baseWhere, &baseParam, "balance_type = ?", req.BalanceType)
	lib.AppendWhere(&baseWhere, &baseParam, "balance_name = ?", req.BalanceName)

	runQuery := "SELECT balance_type, balance_name, exponent, balance_category ,last_update_username, last_update_date FROM balance "
	if len(baseWhere) > 0 {
		runQuery += "WHERE " + baseWhere
	}
	lib.AppendOrderBy(&runQuery, req.Param.OrderBy, req.Param.OrderDir)
	lib.AppendLimit(&runQuery, req.Param.Page, req.Param.PerPage)

	result, _, err = conn.DBOcsConn.SelectQueryByFieldNameSlice(runQuery, baseParam...)
	return result, err
}
