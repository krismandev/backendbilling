package models

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/currency/datastruct"
)

func GetSingleCurrency(currencyID string, conn *connections.Connections, req datastruct.CurrencyRequest) (map[string]string, error) {
	var result map[string]string
	var err error

	// -- EXAMPLE
	// if len(currencyID) == 0 {
	// 	currencyID = req.CurrencyID
	// }
	// query := "SELECT currencyid, currencyname FROM currency WHERE currencyid = ?"
	// results, _, err := conn.DBAppConn.SelectQueryByFieldNameSlice(query, currencyID)
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

func CheckCurrencyExists(currencyID string, conn *connections.Connections) error {
	// var param []interface{}
	// qry := "SELECT COUNT(currencyid) FROM currency WHERE currencyid = ?"
	// param = append(param, currencyID)

	// cnt, _ := conn.DBAppConn.GetFirstData(qry, param...)
	// datacount, _ := strconv.Atoi(cnt)
	// if datacount == 0 {
	// 	return errors.New("Currency ID is not exists")
	// }
	return nil
}

func CheckCurrencyDuplicate(exceptID string, conn *connections.Connections, req datastruct.CurrencyRequest) error {
	// var param []interface{}
	// qry := "SELECT COUNT(currencyid) FROM currency WHERE currencyid = ?"
	// param = append(param, req.CurrencyID)
	// if len(exceptID) > 0 {
	// 	qry += " AND currencyid <> ?"
	// 	param = append(param, exceptID)
	// }

	// cnt, _ := conn.DBAppConn.GetFirstData(qry, param...)
	// datacount, _ := strconv.Atoi(cnt)
	// if datacount > 0 {
	// 	return errors.New("Currency ID is already exists. Please use another Currency ID")
	// }
	return nil
}
