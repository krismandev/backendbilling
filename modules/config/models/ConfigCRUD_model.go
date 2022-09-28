package models

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/lib"
	"backendbillingdashboard/modules/config/datastruct"
)

func GetConfigFromRequest(conn *connections.Connections, req datastruct.ConfigRequest) ([]map[string]string, error) {
	var result []map[string]string
	var err error

	// -- THIS IS BASIC GET REQUEST EXAMPLE LOGIC
	var baseWhere string
	var baseParam []interface{}

	lib.AppendWhere(&baseWhere, &baseParam, "key = ?", req.Key)
	lib.AppendWhere(&baseWhere, &baseParam, "type = ?", req.Type)
	// if len(req.ListConfigID) > 0 {
	// 	var baseIn string
	// 	for _, prid := range req.ListConfigID {
	// 		lib.AppendComma(&baseIn, &baseParam, "?", prid)
	// 	}
	// 	lib.AppendWhereRaw(&baseWhere, "configid IN ("+baseIn+")")
	// }

	runQuery := "SELECT config.key, config.type, config.value FROM config "
	if len(baseWhere) > 0 {
		runQuery += "WHERE " + baseWhere
	}
	lib.AppendOrderBy(&runQuery, req.Param.OrderBy, req.Param.OrderDir)
	lib.AppendLimit(&runQuery, req.Param.Page, req.Param.PerPage)

	result, _, err = conn.DBAppConn.SelectQueryByFieldNameSlice(runQuery, baseParam...)
	return result, err
}

// func InsertConfig(conn *connections.Connections, req datastruct.ConfigRequest) error {
// 	var err error

// 	// -- THIS IS BASIC INSERT EXAMPLE
// 	// var baseIn string
// 	// var baseParam []interface{}

// 	// lib.AppendComma(&baseIn, &baseParam, "?", req.ConfigID)
// 	// lib.AppendComma(&baseIn, &baseParam, "?", req.ConfigName)

// 	// qry := "INSERT INTO config (configid, configname) VALUES (" + baseIn + ")"
// 	// _, _, err = conn.DBAppConn.Exec(qry, baseParam...)

// 	return err
// }

// func UpdateConfig(conn *connections.Connections, req datastruct.ConfigRequest) error {
// 	var err error

// 	// -- THIS IS BASIC UPDATE EXAMPLE
// 	// var baseUp string
// 	// var baseParam []interface{}

// 	// lib.AppendComma(&baseUp, &baseParam, "configname = ?", req.ConfigName)
// 	// qry := "UPDATE config SET " + baseUp + " WHERE configid = ?"
// 	// baseParam = append(baseParam, req.ConfigID)
// 	// _, _, err = conn.DBAppConn.Exec(qry, baseParam...)
// 	return err
// }

// func DeleteConfig(conn *connections.Connections, req datastruct.ConfigRequest) error {
// 	var err error
// 	// -- THIS IS DELETE LOGIC EXAMPLE
// 	// qry := "DELETE FROM config WHERE configid = ?"
// 	// _, _, err = conn.DBAppConn.Exec(qry, req.ConfigID)
// 	return err
// }
