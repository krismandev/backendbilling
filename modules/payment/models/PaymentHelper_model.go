package models

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/payment/datastruct"
	"errors"
)

func GetSinglePayment(paymentID string, conn *connections.Connections, req datastruct.PaymentRequest) (map[string]string, error) {
	var result map[string]string
	var err error

	// -- EXAMPLE
	// if len(paymentID) == 0 {
	// 	paymentID = req.PaymentID
	// }
	// query := "SELECT paymentid, paymentname FROM payment WHERE paymentid = ?"
	// results, _, err := conn.DBAppConn.SelectQueryByFieldNameSlice(query, paymentID)
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

func CheckPaymentExists(paymentID string, conn *connections.Connections) error {
	// var param []interface{}
	// qry := "SELECT COUNT(paymentid) FROM payment WHERE paymentid = ?"
	// param = append(param, paymentID)

	// cnt, _ := conn.DBAppConn.GetFirstData(qry, param...)
	// datacount, _ := strconv.Atoi(cnt)
	// if datacount == 0 {
	// 	return errors.New("Payment ID is not exists")
	// }
	return nil
}

func CheckPaymentDuplicate(exceptID string, conn *connections.Connections, req datastruct.PaymentRequest) error {
	// var param []interface{}
	// qry := "SELECT COUNT(paymentid) FROM payment WHERE paymentid = ?"
	// param = append(param, req.PaymentID)
	// if len(exceptID) > 0 {
	// 	qry += " AND paymentid <> ?"
	// 	param = append(param, exceptID)
	// }

	// cnt, _ := conn.DBAppConn.GetFirstData(qry, param...)
	// datacount, _ := strconv.Atoi(cnt)
	// if datacount > 0 {
	// 	return errors.New("Payment ID is already exists. Please use another Payment ID")
	// }
	return nil
}

func CheckIsItPaid(invoiceID string, conn *connections.Connections, req datastruct.PaymentRequest) error {
	//return error if invoice already paid
	qryGetPaymentStatus := "SELECT paid FROM invoice where invoice_id=?"
	paymentStatus, _ := conn.DBAppConn.GetFirstData(qryGetPaymentStatus, invoiceID)
	if paymentStatus == "1" {
		return errors.New("This invoice already paid")
	}
	// qryGetInvoice := "SELECT"
	return nil
}

func UpdateControlId(conn *connections.Connections, newId string, key string) error {
	var err error
	_, _, err = conn.DBAppConn.Exec("UPDATE control_id set last_id = ? WHERE control_id.key=?", newId, key)
	return err
}
