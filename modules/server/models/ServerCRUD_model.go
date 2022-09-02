package models

import (
	"billingdashboard/connections"
	"billingdashboard/lib"
	"billingdashboard/modules/server/datastruct"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func GetServerFromRequest(conn *connections.Connections, req datastruct.ServerRequest) ([]map[string]string, error) {
	var result []map[string]string
	var err error

	// -- THIS IS BASIC GET REQUEST EXAMPLE LOGIC
	var baseWhere string
	var baseParam []interface{}

	lib.AppendWhere(&baseWhere, &baseParam, "server_id = ?", req.ServerID)
	lib.AppendWhere(&baseWhere, &baseParam, "server_name = ?", req.ServerName)
	if len(req.ListServerID) > 0 {
		var baseIn string
		for _, prid := range req.ListServerID {
			lib.AppendComma(&baseIn, &baseParam, "?", prid)
		}
		lib.AppendWhereRaw(&baseWhere, "server_id IN ("+baseIn+")")
	}

	runQuery := "SELECT server_id, server_name, server_url FROM server "
	if len(baseWhere) > 0 {
		runQuery += "WHERE " + baseWhere
	}
	lib.AppendOrderBy(&runQuery, req.Param.OrderBy, req.Param.OrderDir)
	lib.AppendLimit(&runQuery, req.Param.Page, req.Param.PerPage)

	result, _, err = conn.DBAppConn.SelectQueryByFieldNameSlice(runQuery, baseParam...)
	return result, err
}

func InsertServer(conn *connections.Connections, req datastruct.ServerRequest) error {
	var err error

	// -- THIS IS BASIC INSERT EXAMPLE
	var baseIn string
	var baseParam []interface{}

	lastId, _ := conn.DBAppConn.GetFirstData("SELECT last_id FROM control_id where control_id.key=?", "server")
	// lastId, _ := conn.DBAppConn.GetFirstData("SELECT last_id FROM control_id where control_id.key=?", "account")

	intLastId, err := strconv.Atoi(lastId)
	insertId := intLastId + 1

	insertIdString := strconv.Itoa(insertId)

	lib.AppendComma(&baseIn, &baseParam, "?", insertIdString)
	lib.AppendComma(&baseIn, &baseParam, "?", req.ServerName)
	lib.AppendComma(&baseIn, &baseParam, "?", req.ServerUrl)

	log.Info("InsertParam - ", baseParam)
	qry := "INSERT INTO server (server_id, server_name, server_url) VALUES (" + baseIn + ")"
	_, _, err = conn.DBAppConn.Exec(qry, baseParam...)

	_, _, err = conn.DBAppConn.Exec("UPDATE control_id set last_id=? where control_id.key=?", insertIdString, "server")

	return err
}

func UpdateServer(conn *connections.Connections, req datastruct.ServerRequest) error {
	var err error

	// -- THIS IS BASIC UPDATE EXAMPLE
	var baseUp string
	var baseParam []interface{}

	lib.AppendComma(&baseUp, &baseParam, "server_name = ?", req.ServerName)
	lib.AppendComma(&baseUp, &baseParam, "server_url = ?", req.ServerUrl)
	qry := "UPDATE server SET " + baseUp + " WHERE server_id = ?"
	baseParam = append(baseParam, req.ServerID)
	_, _, err = conn.DBAppConn.Exec(qry, baseParam...)
	return err
}

func DeleteServer(conn *connections.Connections, req datastruct.ServerRequest) error {
	var err error
	// -- THIS IS DELETE LOGIC EXAMPLE
	qry := "DELETE FROM server WHERE server_id = ?"
	_, _, err = conn.DBAppConn.Exec(qry, req.ServerID)
	return err
}

func GetServerAccountFromRequest(conn *connections.Connections, req datastruct.ServerAccountRequest) ([]map[string]string, error) {
	var result []map[string]string
	var err error

	// -- THIS IS BASIC GET REQUEST EXAMPLE LOGIC
	var baseWhere string
	var baseParam []interface{}

	lib.AppendWhere(&baseWhere, &baseParam, "server_id = ?", req.ServerID)
	lib.AppendWhere(&baseWhere, &baseParam, "account_id = ?", req.AccountID)
	if len(req.ListAccountID) > 0 {
		var baseIn string
		for _, accId := range req.ListAccountID {
			lib.AppendComma(&baseIn, &baseParam, "?", accId)
		}
		lib.AppendWhereRaw(&baseWhere, "account_id IN ("+baseIn+")")
	}

	runQuery := "SELECT server_id, account_id ,external_account_id, last_update_username, last_update_date FROM server_account "
	if len(baseWhere) > 0 {
		runQuery += "WHERE " + baseWhere
	}
	lib.AppendOrderBy(&runQuery, req.Param.OrderBy, req.Param.OrderDir)
	lib.AppendLimit(&runQuery, req.Param.Page, req.Param.PerPage)

	result, _, err = conn.DBAppConn.SelectQueryByFieldNameSlice(runQuery, baseParam...)
	return result, err
}
