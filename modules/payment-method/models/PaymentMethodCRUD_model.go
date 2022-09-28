package models

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/lib"
	"backendbillingdashboard/modules/payment-method/datastruct"
)

func GetPaymentMethodFromRequest(conn *connections.Connections, req datastruct.PaymentMethodRequest) ([]map[string]string, error) {
	var result []map[string]string
	var err error

	// -- THIS IS BASIC GET REQUEST EXAMPLE LOGIC
	var baseWhere string
	var baseParam []interface{}

	lib.AppendWhere(&baseWhere, &baseParam, "payment_method.key = ?", req.Key)
	lib.AppendWhere(&baseWhere, &baseParam, "payment_method_name = ?", req.PaymentMethodName)
	lib.AppendWhere(&baseWhere, &baseParam, "currency_code = ?", req.CurrencyCode)
	if len(req.ListKey) > 0 {
		var baseIn string
		for _, prid := range req.ListKey {
			lib.AppendComma(&baseIn, &baseParam, "?", prid)
		}
		lib.AppendWhereRaw(&baseWhere, "payment_method.key IN ("+baseIn+")")
	}

	runQuery := "SELECT payment_method.key, payment_method_name, need_clearing_date, need_card_number, bank_name, branch, account_name, account_no, code, status, payment_type, currency_code FROM payment_method "
	if len(baseWhere) > 0 {
		runQuery += "WHERE " + baseWhere
	}
	lib.AppendOrderBy(&runQuery, req.Param.OrderBy, req.Param.OrderDir)
	lib.AppendLimit(&runQuery, req.Param.Page, req.Param.PerPage)

	result, _, err = conn.DBAppConn.SelectQueryByFieldNameSlice(runQuery, baseParam...)
	return result, err
}

func InsertPaymentMethod(conn *connections.Connections, req datastruct.PaymentMethodRequest) error {
	var err error

	// -- THIS IS BASIC INSERT EXAMPLE
	var baseIn string
	var baseParam []interface{}

	lib.AppendComma(&baseIn, &baseParam, "?", req.Key)
	lib.AppendComma(&baseIn, &baseParam, "?", req.PaymentMethodName)
	lib.AppendComma(&baseIn, &baseParam, "?", req.NeedClearingDate)
	lib.AppendComma(&baseIn, &baseParam, "?", req.NeedCardNumber)
	lib.AppendComma(&baseIn, &baseParam, "?", req.BankName)
	lib.AppendComma(&baseIn, &baseParam, "?", req.Branch)
	lib.AppendComma(&baseIn, &baseParam, "?", req.AccountName)
	lib.AppendComma(&baseIn, &baseParam, "?", req.AccountNo)
	lib.AppendComma(&baseIn, &baseParam, "?", req.Code)
	lib.AppendComma(&baseIn, &baseParam, "?", req.CurrencyCode)
	lib.AppendComma(&baseIn, &baseParam, "?", req.LastUpdateUsername)
	lib.AppendCommaRaw(&baseIn, "now()")

	qry := "INSERT INTO payment_method (payment_method.key, payment_method_name, need_clearing_date, need_card_number,bank_name,branch, account_name,account_no, code, currency_code, last_update_username, last_update_date) VALUES (" + baseIn + ")"
	_, _, err = conn.DBAppConn.Exec(qry, baseParam...)

	return err
}

func UpdatePaymentMethod(conn *connections.Connections, req datastruct.PaymentMethodRequest) error {
	var err error

	// -- THIS IS BASIC UPDATE EXAMPLE
	var baseUp string
	var baseParam []interface{}

	lib.AppendComma(&baseUp, &baseParam, "payment_method_name = ?", req.PaymentMethodName)
	lib.AppendComma(&baseUp, &baseParam, "need_clearing_date = ?", req.NeedClearingDate)
	lib.AppendComma(&baseUp, &baseParam, "need_card_number = ?", req.NeedCardNumber)
	lib.AppendComma(&baseUp, &baseParam, "bank_name = ?", req.BankName)
	lib.AppendComma(&baseUp, &baseParam, "branch = ?", req.Branch)
	lib.AppendComma(&baseUp, &baseParam, "account_name = ?", req.AccountName)
	lib.AppendComma(&baseUp, &baseParam, "account_no = ?", req.AccountNo)
	lib.AppendComma(&baseUp, &baseParam, "code = ?", req.Code)
	lib.AppendComma(&baseUp, &baseParam, "currency_code = ?", req.CurrencyCode)
	lib.AppendComma(&baseUp, &baseParam, "last_update_username = ?", req.LastUpdateUsername)
	qry := "UPDATE payment_method SET " + baseUp + " WHERE payment_method.key = ?"
	baseParam = append(baseParam, req.Key)
	_, _, err = conn.DBAppConn.Exec(qry, baseParam...)
	return err
}

func DeletePaymentMethod(conn *connections.Connections, req datastruct.PaymentMethodRequest) error {
	var err error
	// -- THIS IS DELETE LOGIC EXAMPLE
	qry := "UPDATE payment_method SET status = ? where payment_method.key = ?"
	_, _, err = conn.DBAppConn.Exec(qry, "D", req.Key)
	return err
}
