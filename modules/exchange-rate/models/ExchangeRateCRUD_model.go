package models

import (
	"billingdashboard/connections"
	"billingdashboard/lib"
	"billingdashboard/modules/exchange-rate/datastruct"
)

func GetExchangeRateFromRequest(conn *connections.Connections, req datastruct.ExchangeRateRequest) ([]map[string]string, error) {
	var result []map[string]string
	var err error

	// -- THIS IS BASIC GET REQUEST EXAMPLE LOGIC
	var baseWhere string
	var baseParam []interface{}

	lib.AppendWhere(&baseWhere, &baseParam, "currency = ?", req.Currency)
	lib.AppendWhere(&baseWhere, &baseParam, "date = ?", req.Date)

	runQuery := "SELECT date, currency, from_currency, to_currency, convert_value FROM exchange_rate "
	if len(baseWhere) > 0 {
		runQuery += "WHERE " + baseWhere
	}
	lib.AppendOrderBy(&runQuery, req.Param.OrderBy, req.Param.OrderDir)
	lib.AppendLimit(&runQuery, req.Param.Page, req.Param.PerPage)

	result, _, err = conn.DBAppConn.SelectQueryByFieldNameSlice(runQuery, baseParam...)
	return result, err
}

func InsertExchangeRate(conn *connections.Connections, req datastruct.ExchangeRateRequest) error {
	var err error

	// -- THIS IS BASIC INSERT EXAMPLE
	var baseIn string
	var baseParam []interface{}

	lib.AppendComma(&baseIn, &baseParam, "?", req.Date)
	lib.AppendComma(&baseIn, &baseParam, "?", req.Currency)
	lib.AppendComma(&baseIn, &baseParam, "?", req.FromCurrency)
	lib.AppendComma(&baseIn, &baseParam, "?", req.ToCurrency)
	lib.AppendComma(&baseIn, &baseParam, "?", req.ConvertValue)

	qry := "INSERT INTO exchange_rate (exchange_rate.date, currency, from_currency, to_currency, convert_value) VALUES (" + baseIn + ")"
	_, _, err = conn.DBAppConn.Exec(qry, baseParam...)

	return err
}

func UpdateExchangeRate(conn *connections.Connections, req datastruct.ExchangeRateRequest) error {
	var err error

	// -- THIS IS BASIC UPDATE EXAMPLE
	// var baseUp string
	// var baseParam []interface{}

	// lib.AppendComma(&baseUp, &baseParam, "stubname = ?", req.ExchangeRateName)
	// qry := "UPDATE stub SET " + baseUp + " WHERE stubid = ?"
	// baseParam = append(baseParam, req.ExchangeRateID)
	// _, _, err = conn.DBAppConn.Exec(qry, baseParam...)
	return err
}

func DeleteExchangeRate(conn *connections.Connections, req datastruct.ExchangeRateRequest) error {
	var err error
	// -- THIS IS DELETE LOGIC EXAMPLE
	// qry := "DELETE FROM stub WHERE stubid = ?"
	// _, _, err = conn.DBAppConn.Exec(qry, req.ExchangeRateID)
	return err
}