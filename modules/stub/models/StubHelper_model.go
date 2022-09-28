package models

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/stub/datastruct"
)

func GetSingleStub(stubID string, conn *connections.Connections, req datastruct.StubRequest) (map[string]string, error) {
	var result map[string]string
	var err error

	// -- EXAMPLE
	// if len(stubID) == 0 {
	// 	stubID = req.StubID
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

func CheckStubExists(stubID string, conn *connections.Connections) error {
	// var param []interface{}
	// qry := "SELECT COUNT(stubid) FROM stub WHERE stubid = ?"
	// param = append(param, stubID)

	// cnt, _ := conn.DBAppConn.GetFirstData(qry, param...)
	// datacount, _ := strconv.Atoi(cnt)
	// if datacount == 0 {
	// 	return errors.New("Stub ID is not exists")
	// }
	return nil
}

func CheckStubDuplicate(exceptID string, conn *connections.Connections, req datastruct.StubRequest) error {
	// var param []interface{}
	// qry := "SELECT COUNT(stubid) FROM stub WHERE stubid = ?"
	// param = append(param, req.StubID)
	// if len(exceptID) > 0 {
	// 	qry += " AND stubid <> ?"
	// 	param = append(param, exceptID)
	// }

	// cnt, _ := conn.DBAppConn.GetFirstData(qry, param...)
	// datacount, _ := strconv.Atoi(cnt)
	// if datacount > 0 {
	// 	return errors.New("Stub ID is already exists. Please use another Stub ID")
	// }
	return nil
}
