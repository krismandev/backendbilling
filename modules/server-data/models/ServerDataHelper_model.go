package models

import (
	"billingdashboard/connections"
	"billingdashboard/modules/server-data/datastruct"
)

func GetSingleServerData(serverDataID string, conn *connections.Connections, req datastruct.ServerDataRequest) (map[string]string, error) {
	var result map[string]string
	var err error

	// -- EXAMPLE
	// if len(server-dataID) == 0 {
	// 	server-dataID = req.ServerDataID
	// }
	// query := "SELECT server-dataid, server-dataname FROM server-data WHERE server-dataid = ?"
	// results, _, err := conn.DBAppConn.SelectQueryByFieldNameSlice(query, server-dataID)
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

func CheckServerDataExists(serverDataID string, conn *connections.Connections) error {
	// var param []interface{}
	// qry := "SELECT COUNT(server-dataid) FROM server-data WHERE server-dataid = ?"
	// param = append(param, server-dataID)

	// cnt, _ := conn.DBAppConn.GetFirstData(qry, param...)
	// datacount, _ := strconv.Atoi(cnt)
	// if datacount == 0 {
	// 	return errors.New("ServerData ID is not exists")
	// }
	return nil
}

func CheckServerDataDuplicate(exceptID string, conn *connections.Connections, req datastruct.ServerDataRequest) error {
	// var param []interface{}
	// qry := "SELECT COUNT(server-dataid) FROM server-data WHERE server-dataid = ?"
	// param = append(param, req.ServerDataID)
	// if len(exceptID) > 0 {
	// 	qry += " AND server-dataid <> ?"
	// 	param = append(param, exceptID)
	// }

	// cnt, _ := conn.DBAppConn.GetFirstData(qry, param...)
	// datacount, _ := strconv.Atoi(cnt)
	// if datacount > 0 {
	// 	return errors.New("ServerData ID is already exists. Please use another ServerData ID")
	// }
	return nil
}
