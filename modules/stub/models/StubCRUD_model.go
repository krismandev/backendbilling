package models

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/stub/datastruct"
)

func GetStubFromRequest(conn *connections.Connections, req datastruct.StubRequest) ([]map[string]string, error) {
	var result []map[string]string
	var err error

	// -- THIS IS BASIC GET REQUEST EXAMPLE LOGIC
	// var baseWhere string
	// var baseParam []interface{}

	// lib.AppendWhere(&baseWhere, &baseParam, "stubid = ?", req.StubID)
	// lib.AppendWhere(&baseWhere, &baseParam, "stubname = ?", req.StubName)
	// if len(req.ListStubID) > 0 {
	// 	var baseIn string
	// 	for _, prid := range req.ListStubID {
	// 		lib.AppendComma(&baseIn, &baseParam, "?", prid)
	// 	}
	// 	lib.AppendWhereRaw(&baseWhere, "stubid IN ("+baseIn+")")
	// }

	// runQuery := "SELECT stubid, stubname FROM stub "
	// if len(baseWhere) > 0 {
	// 	runQuery += "WHERE " + baseWhere
	// }
	// lib.AppendOrderBy(&runQuery, req.Param.OrderBy, req.Param.OrderDir)
	// lib.AppendLimit(&runQuery, req.Param.Page, req.Param.PerPage)

	// result, _, err = conn.DBAppConn.SelectQueryByFieldNameSlice(runQuery, baseParam...)
	return result, err
}

func InsertStub(conn *connections.Connections, req datastruct.StubRequest) error {
	var err error

	// -- THIS IS BASIC INSERT EXAMPLE
	// var baseIn string
	// var baseParam []interface{}

	// lib.AppendComma(&baseIn, &baseParam, "?", req.StubID)
	// lib.AppendComma(&baseIn, &baseParam, "?", req.StubName)

	// qry := "INSERT INTO stub (stubid, stubname) VALUES (" + baseIn + ")"
	// _, _, err = conn.DBAppConn.Exec(qry, baseParam...)

	return err
}

func UpdateStub(conn *connections.Connections, req datastruct.StubRequest) error {
	var err error

	// -- THIS IS BASIC UPDATE EXAMPLE
	// var baseUp string
	// var baseParam []interface{}

	// lib.AppendComma(&baseUp, &baseParam, "stubname = ?", req.StubName)
	// qry := "UPDATE stub SET " + baseUp + " WHERE stubid = ?"
	// baseParam = append(baseParam, req.StubID)
	// _, _, err = conn.DBAppConn.Exec(qry, baseParam...)
	return err
}

func DeleteStub(conn *connections.Connections, req datastruct.StubRequest) error {
	var err error
	// -- THIS IS DELETE LOGIC EXAMPLE
	// qry := "DELETE FROM stub WHERE stubid = ?"
	// _, _, err = conn.DBAppConn.Exec(qry, req.StubID)
	return err
}
