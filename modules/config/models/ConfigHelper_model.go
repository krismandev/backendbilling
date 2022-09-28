package models

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/config/datastruct"
)

func GetSingleConfig(configID string, conn *connections.Connections, req datastruct.ConfigRequest) (map[string]string, error) {
	var result map[string]string
	var err error

	// -- EXAMPLE
	// if len(configID) == 0 {
	// 	configID = req.ConfigID
	// }
	// query := "SELECT configid, configname FROM config WHERE configid = ?"
	// results, _, err := conn.DBAppConn.SelectQueryByFieldNameSlice(query, configID)
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

func CheckConfigExists(configID string, conn *connections.Connections) error {
	// var param []interface{}
	// qry := "SELECT COUNT(configid) FROM config WHERE configid = ?"
	// param = append(param, configID)

	// cnt, _ := conn.DBAppConn.GetFirstData(qry, param...)
	// datacount, _ := strconv.Atoi(cnt)
	// if datacount == 0 {
	// 	return errors.New("Config ID is not exists")
	// }
	return nil
}

func CheckConfigDuplicate(exceptID string, conn *connections.Connections, req datastruct.ConfigRequest) error {
	// var param []interface{}
	// qry := "SELECT COUNT(configid) FROM config WHERE configid = ?"
	// param = append(param, req.ConfigID)
	// if len(exceptID) > 0 {
	// 	qry += " AND configid <> ?"
	// 	param = append(param, exceptID)
	// }

	// cnt, _ := conn.DBAppConn.GetFirstData(qry, param...)
	// datacount, _ := strconv.Atoi(cnt)
	// if datacount > 0 {
	// 	return errors.New("Config ID is already exists. Please use another Config ID")
	// }
	return nil
}
