package models

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/payment-method/datastruct"
)

func GetSinglePaymentMethod(stubID string, conn *connections.Connections, req datastruct.PaymentMethodRequest) (map[string]string, error) {
	var result map[string]string
	var err error

	// -- EXAMPLE
	// if len(stubID) == 0 {
	// 	stubID = req.PaymentMethodID
	// }
	// query := "SELECT stubid, stubname FROM stub WHERE stubid = ?"
	// results, _, err := conn.DBAppConn.SelectQueryByFieldNameSlice(query, stubID)
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

func CheckPaymentMethodExists(stubID string, conn *connections.Connections) error {
	// var param []interface{}
	// qry := "SELECT COUNT(stubid) FROM stub WHERE stubid = ?"
	// param = append(param, stubID)

	// cnt, _ := conn.DBAppConn.GetFirstData(qry, param...)
	// datacount, _ := strconv.Atoi(cnt)
	// if datacount == 0 {
	// 	return errors.New("PaymentMethod ID is not exists")
	// }
	return nil
}

func CheckPaymentMethodDuplicate(exceptID string, conn *connections.Connections, req datastruct.PaymentMethodRequest) error {
	// var param []interface{}
	// qry := "SELECT COUNT(stubid) FROM stub WHERE stubid = ?"
	// param = append(param, req.PaymentMethodID)
	// if len(exceptID) > 0 {
	// 	qry += " AND stubid <> ?"
	// 	param = append(param, exceptID)
	// }

	// cnt, _ := conn.DBAppConn.GetFirstData(qry, param...)
	// datacount, _ := strconv.Atoi(cnt)
	// if datacount > 0 {
	// 	return errors.New("PaymentMethod ID is already exists. Please use another PaymentMethod ID")
	// }
	return nil
}
