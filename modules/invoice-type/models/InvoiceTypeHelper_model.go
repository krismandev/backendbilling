package models

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/invoice-type/datastruct"
)

func GetSingleInvoiceType(invoiceTypeID string, conn *connections.Connections, req datastruct.InvoiceTypeRequest) (map[string]string, error) {
	var result map[string]string
	var err error

	// -- EXAMPLE
	// if len(invoice-typeID) == 0 {
	// 	invoice-typeID = req.InvoiceTypeID
	// }
	// query := "SELECT invoice-typeid, invoice-typename FROM invoice-type WHERE invoice-typeid = ?"
	// results, _, err := conn.DBAppConn.SelectQueryByFieldNameSlice(query, invoice-typeID)
	// if err != nil {
	// 	return result, err
	// }

	// // convert from []map[string]string to single map[string]string
	// for _, res := range results {
	// 	result = res
	// 	break
	// }
	return result, err
}

func CheckInvoiceTypeExists(invoiceTypeID string, conn *connections.Connections) error {
	// var param []interface{}
	// qry := "SELECT COUNT(invoice-typeid) FROM invoice-type WHERE invoice-typeid = ?"
	// param = append(param, invoice-typeID)

	// cnt, _ := conn.DBAppConn.GetFirstData(qry, param...)
	// datacount, _ := strconv.Atoi(cnt)
	// if datacount == 0 {
	// 	return errors.New("InvoiceType ID is not exists")
	// }
	return nil
}

func CheckInvoiceTypeDuplicate(exceptID string, conn *connections.Connections, req datastruct.InvoiceTypeRequest) error {
	// var param []interface{}
	// qry := "SELECT COUNT(invoice-typeid) FROM invoice-type WHERE invoice-typeid = ?"
	// param = append(param, req.InvoiceTypeID)
	// if len(exceptID) > 0 {
	// 	qry += " AND invoice-typeid <> ?"
	// 	param = append(param, exceptID)
	// }

	// cnt, _ := conn.DBAppConn.GetFirstData(qry, param...)
	// datacount, _ := strconv.Atoi(cnt)
	// if datacount > 0 {
	// 	return errors.New("InvoiceType ID is already exists. Please use another InvoiceType ID")
	// }
	return nil
}
