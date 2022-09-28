package models

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/exchange-rate/datastruct"
)

func GetSingleExchangeRate(stubID string, conn *connections.Connections, req datastruct.ExchangeRateRequest) (map[string]string, error) {
	var result map[string]string
	var err error

	// -- EXAMPLE
	// if len(stubID) == 0 {
	// 	stubID = req.ExchangeRateID
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

func CheckExchangeRateExists(stubID string, conn *connections.Connections) error {
	// var param []interface{}
	// qry := "SELECT COUNT(stubid) FROM stub WHERE stubid = ?"
	// param = append(param, stubID)

	// cnt, _ := conn.DBAppConn.GetFirstData(qry, param...)
	// datacount, _ := strconv.Atoi(cnt)
	// if datacount == 0 {
	// 	return errors.New("ExchangeRate ID is not exists")
	// }
	return nil
}

func CheckExchangeRateDuplicate(exceptID string, conn *connections.Connections, req datastruct.ExchangeRateRequest) error {
	// var param []interface{}
	// qry := "SELECT COUNT(stubid) FROM stub WHERE stubid = ?"
	// param = append(param, req.ExchangeRateID)
	// if len(exceptID) > 0 {
	// 	qry += " AND stubid <> ?"
	// 	param = append(param, exceptID)
	// }

	// cnt, _ := conn.DBAppConn.GetFirstData(qry, param...)
	// datacount, _ := strconv.Atoi(cnt)
	// if datacount > 0 {
	// 	return errors.New("ExchangeRate ID is already exists. Please use another ExchangeRate ID")
	// }
	return nil
}
