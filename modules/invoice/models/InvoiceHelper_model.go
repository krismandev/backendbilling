package models

import (
	"billingdashboard/connections"
	"billingdashboard/modules/invoice/datastruct"
	"errors"
	"strconv"
)

func GetSingleInvoice(invoiceID string, conn *connections.Connections, req datastruct.InvoiceRequest) (map[string]string, error) {
	var result map[string]string
	var err error

	// -- EXAMPLE
	// if len(invoiceID) == 0 {
	// 	invoiceID = req.InvoiceID
	// }
	// query := "SELECT invoiceid, invoicename FROM invoice WHERE invoiceid = ?"
	// results, _, err := conn.DBAppConn.SelectQueryByFieldNameSlice(query, invoiceID)
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

func CheckInvoiceExists(invoiceID string, conn *connections.Connections) error {
	// var param []interface{}
	// qry := "SELECT COUNT(invoiceid) FROM invoice WHERE invoiceid = ?"
	// param = append(param, invoiceID)

	// cnt, _ := conn.DBAppConn.GetFirstData(qry, param...)
	// datacount, _ := strconv.Atoi(cnt)
	// if datacount == 0 {
	// 	return errors.New("Invoice ID is not exists")
	// }
	return nil
}

func CheckInvoiceNoDuplicate2(invoiceNo string, conn *connections.Connections) error {
	var param []interface{}
	param = append(param, invoiceNo)
	param = append(param, "A")
	qry := "SELECT COUNT(invoice_no) FROM invoice WHERE invoice_no = ? AND invoicestatus = ?"

	cnt, _ := conn.DBAppConn.GetFirstData(qry, param...)
	datacount, _ := strconv.Atoi(cnt)
	if datacount > 0 {
		return errors.New("Invoice Number is already used. Please use another Invoice Number")
	}
	return nil
}

func CheckInvoiceNoDuplicate(invoiceNo string, conn *connections.Connections, req datastruct.InvoiceRequest) error {
	qryCheckInvoiceNo := "SELECT invoice_id, invoice_no, invoicestatus FROM invoice WHERE invoice_no = ? AND invoicestatus = ?"
	resCheck, countInvoice, errCheckInvoiceNo := conn.DBAppConn.SelectQueryByFieldName(qryCheckInvoiceNo, invoiceNo, "A")
	if errCheckInvoiceNo != nil {
		return errCheckInvoiceNo
	}

	// if countInvoice == 0 {
	// 	qryInserControlIdInvoiceDetail := "INSERT INTO control_id (control_id.key,control_id.period,control_id.last_id) VALUES (?,?,?)"
	// 	_, _, errInsert := conn.DBAppConn.Exec(qryInserControlIdInvoiceDetail, "invoice_detail", "0", "0")
	// 	if errInsert != nil {
	// 		return "", errInsert
	// 	}

	// 	return "0", nil
	// }

	invoice := resCheck[1]

	if countInvoice > 0 {
		if invoice["invoice_id"] == req.InvoiceID {
			return nil
		} else {
			return errors.New("Invoice Number is already used. Please use another Invoice Number")
		}
	}

	return nil
}

func CheckInvoiceDuplicate(exceptID string, conn *connections.Connections, req datastruct.InvoiceRequest) error {
	var param []interface{}
	qry := "SELECT COUNT(invoiceid) FROM invoice WHERE invoiceid = ?"
	param = append(param, req.InvoiceID)
	if len(exceptID) > 0 {
		qry += " AND invoiceid <> ?"
		param = append(param, exceptID)
	}

	cnt, _ := conn.DBAppConn.GetFirstData(qry, param...)
	datacount, _ := strconv.Atoi(cnt)
	if datacount > 0 {
		return errors.New("Invoice ID is already exists. Please use another Invoice ID")
	}
	return nil
}
