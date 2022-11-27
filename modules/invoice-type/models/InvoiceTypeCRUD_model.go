package models

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/lib"
	"backendbillingdashboard/modules/invoice-type/datastruct"
)

func GetInvoiceTypeFromRequest(conn *connections.Connections, req datastruct.InvoiceTypeRequest) ([]map[string]string, error) {
	var result []map[string]string
	var err error

	// -- THIS IS BASIC GET REQUEST EXAMPLE LOGIC
	var baseWhere string
	var baseParam []interface{}

	lib.AppendWhere(&baseWhere, &baseParam, "inv_type_id = ?", req.InvoiceTypeID)
	lib.AppendWhere(&baseWhere, &baseParam, "inv_type_name = ?", req.InvoiceTypeName)
	if len(req.ListInvoiceTypeID) > 0 {
		var baseIn string
		for _, prid := range req.ListInvoiceTypeID {
			lib.AppendComma(&baseIn, &baseParam, "?", prid)
		}
		lib.AppendWhereRaw(&baseWhere, "inv_type_id IN ("+baseIn+")")
	}

	runQuery := "SELECT inv_type_id, inv_type_name, server_id, category, load_from_server, last_update_username, last_update_date, currency_code FROM invoice_type "
	if len(baseWhere) > 0 {
		runQuery += "WHERE " + baseWhere
	}
	lib.AppendOrderBy(&runQuery, req.Param.OrderBy, req.Param.OrderDir)
	lib.AppendLimit(&runQuery, req.Param.Page, req.Param.PerPage)

	result, _, err = conn.DBAppConn.SelectQueryByFieldNameSlice(runQuery, baseParam...)
	return result, err
}

func InsertInvoiceType(conn *connections.Connections, req datastruct.InvoiceTypeRequest) error {
	var err error

	// -- THIS IS BASIC INSERT EXAMPLE
	// var baseIn string
	// var baseParam []interface{}

	// lib.AppendComma(&baseIn, &baseParam, "?", req.InvoiceTypeID)
	// lib.AppendComma(&baseIn, &baseParam, "?", req.InvoiceTypeName)

	// qry := "INSERT INTO invoice-type (invoice-typeid, invoice-typename) VALUES (" + baseIn + ")"
	// _, _, err = conn.DBAppConn.Exec(qry, baseParam...)

	return err
}

func UpdateInvoiceType(conn *connections.Connections, req datastruct.InvoiceTypeRequest) error {
	var err error

	// -- THIS IS BASIC UPDATE EXAMPLE
	// var baseUp string
	// var baseParam []interface{}

	// lib.AppendComma(&baseUp, &baseParam, "invoice-typename = ?", req.InvoiceTypeName)
	// qry := "UPDATE invoice-type SET " + baseUp + " WHERE invoice-typeid = ?"
	// baseParam = append(baseParam, req.InvoiceTypeID)
	// _, _, err = conn.DBAppConn.Exec(qry, baseParam...)
	return err
}

func DeleteInvoiceType(conn *connections.Connections, req datastruct.InvoiceTypeRequest) error {
	var err error
	// -- THIS IS DELETE LOGIC EXAMPLE
	// qry := "DELETE FROM invoice-type WHERE invoice-typeid = ?"
	// _, _, err = conn.DBAppConn.Exec(qry, req.InvoiceTypeID)
	return err
}
