package models

import (
	"billingdashboard/connections"
	"billingdashboard/lib"
	"billingdashboard/modules/payment-method/datastruct"
)

func GetPaymentMethodFromRequest(conn *connections.Connections, req datastruct.PaymentMethodRequest) ([]map[string]string, error) {
	var result []map[string]string
	var err error

	// -- THIS IS BASIC GET REQUEST EXAMPLE LOGIC
	var baseWhere string
	var baseParam []interface{}

	lib.AppendWhere(&baseWhere, &baseParam, "payment_method.key = ?", req.Key)
	lib.AppendWhere(&baseWhere, &baseParam, "payment_method_name = ?", req.PaymentMethodName)
	if len(req.ListKey) > 0 {
		var baseIn string
		for _, prid := range req.ListKey {
			lib.AppendComma(&baseIn, &baseParam, "?", prid)
		}
		lib.AppendWhereRaw(&baseWhere, "list_key IN ("+baseIn+")")
	}

	runQuery := "SELECT payment_method.key, payment_method_name, need_clearing_date, need_card_number FROM payment_method "
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
	// var baseIn string
	// var baseParam []interface{}

	// lib.AppendComma(&baseIn, &baseParam, "?", req.PaymentMethodID)
	// lib.AppendComma(&baseIn, &baseParam, "?", req.PaymentMethodName)

	// qry := "INSERT INTO stub (stubid, stubname) VALUES (" + baseIn + ")"
	// _, _, err = conn.DBAppConn.Exec(qry, baseParam...)

	return err
}

func UpdatePaymentMethod(conn *connections.Connections, req datastruct.PaymentMethodRequest) error {
	var err error

	// -- THIS IS BASIC UPDATE EXAMPLE
	// var baseUp string
	// var baseParam []interface{}

	// lib.AppendComma(&baseUp, &baseParam, "stubname = ?", req.PaymentMethodName)
	// qry := "UPDATE stub SET " + baseUp + " WHERE stubid = ?"
	// baseParam = append(baseParam, req.PaymentMethodID)
	// _, _, err = conn.DBAppConn.Exec(qry, baseParam...)
	return err
}

func DeletePaymentMethod(conn *connections.Connections, req datastruct.PaymentMethodRequest) error {
	var err error
	// -- THIS IS DELETE LOGIC EXAMPLE
	// qry := "DELETE FROM stub WHERE stubid = ?"
	// _, _, err = conn.DBAppConn.Exec(qry, req.PaymentMethodID)
	return err
}
