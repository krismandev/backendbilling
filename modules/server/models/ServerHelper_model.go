package models

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/server/datastruct"
	"errors"
	"strconv"
)

func GetSingleServer(serverID string, conn *connections.Connections, req datastruct.ServerRequest) (map[string]string, error) {
	var result map[string]string
	var err error

	// -- EXAMPLE
	// if len(serverID) == 0 {
	// 	serverID = req.ServerID
	// }
	// query := "SELECT serverid, servername FROM server WHERE serverid = ?"
	// results, _, err := conn.DBAppConn.SelectQueryByFieldNameSlice(query, serverID)
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

func CheckServerExists(serverID string, conn *connections.Connections) error {
	var param []interface{}
	qry := "SELECT COUNT(server_id) FROM server WHERE server_id = ?"
	param = append(param, serverID)

	cnt, _ := conn.DBAppConn.GetFirstData(qry, param...)
	datacount, _ := strconv.Atoi(cnt)
	if datacount == 0 {
		return errors.New("Server ID is not exists")
	}
	return nil
}

func CheckServerDuplicate(exceptID string, conn *connections.Connections, req datastruct.ServerRequest) error {
	// var param []interface{}
	// qry := "SELECT COUNT(serverid) FROM server WHERE serverid = ?"
	// param = append(param, req.ServerID)
	// if len(exceptID) > 0 {
	// 	qry += " AND serverid <> ?"
	// 	param = append(param, exceptID)
	// }

	// cnt, _ := conn.DBAppConn.GetFirstData(qry, param...)
	// datacount, _ := strconv.Atoi(cnt)
	// if datacount > 0 {
	// 	return errors.New("Server ID is already exists. Please use another Server ID")
	// }
	return nil
}
