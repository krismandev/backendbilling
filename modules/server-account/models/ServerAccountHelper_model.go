package models

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/server-account/datastruct"
)

func GetSingleServerAccount(ServerAccountID string, conn *connections.Connections, req datastruct.ServerAccountRequest) (map[string]string, error) {
	var result map[string]string
	var err error

	// -- EXAMPLE
	// if len(ServerAccountID) == 0 {
	// 	ServerAccountID = req.ServerAccountID
	// }
	// query := "SELECT ServerAccountid, ServerAccountname FROM ServerAccount WHERE ServerAccountid = ?"
	// results, _, err := conn.DBAppConn.SelectQueryByFieldNameSlice(query, ServerAccountID)
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

func CheckServerAccountExists(ServerAccountID string, conn *connections.Connections) error {
	// var param []interface{}
	// qry := "SELECT COUNT(ServerAccountid) FROM ServerAccount WHERE ServerAccountid = ?"
	// param = append(param, ServerAccountID)

	// cnt, _ := conn.DBAppConn.GetFirstData(qry, param...)
	// datacount, _ := strconv.Atoi(cnt)
	// if datacount == 0 {
	// 	return errors.New("ServerAccount ID is not exists")
	// }
	return nil
}

func CheckServerAccountDuplicate(exceptID string, conn *connections.Connections, req datastruct.ServerAccountRequest) error {
	// var param []interface{}
	// qry := "SELECT COUNT(ServerAccountid) FROM ServerAccount WHERE ServerAccountid = ?"
	// param = append(param, req.ServerAccountID)
	// if len(exceptID) > 0 {
	// 	qry += " AND ServerAccountid <> ?"
	// 	param = append(param, exceptID)
	// }

	// cnt, _ := conn.DBAppConn.GetFirstData(qry, param...)
	// datacount, _ := strconv.Atoi(cnt)
	// if datacount > 0 {
	// 	return errors.New("ServerAccount ID is already exists. Please use another ServerAccount ID")
	// }
	return nil
}
