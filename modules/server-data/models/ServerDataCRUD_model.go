package models

import (
	"billingdashboard/connections"
	"billingdashboard/lib"
	"billingdashboard/modules/server-data/datastruct"

	"github.com/sirupsen/logrus"
)

func GetServerDataFromRequest(conn *connections.Connections, req datastruct.ServerDataRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	var resultQuery []map[string]string
	var err error

	// -- THIS IS BASIC GET REQUEST EXAMPLE LOGIC
	var baseWhere string
	var baseParam []interface{}

	lib.AppendWhere(&baseWhere, &baseParam, "server_data_id = ?", req.ServerDataID)
	lib.AppendWhere(&baseWhere, &baseParam, "server_data.server_id = ?", req.ServerID)
	lib.AppendWhere(&baseWhere, &baseParam, "server_data.account_id = ?", req.AccountID)
	lib.AppendWhere(&baseWhere, &baseParam, "DATE_FORMAT(server_data.external_transdate, '%Y%m') = ?", req.MonthUse)

	if len(req.ListServerDataID) > 0 {
		var baseIn string
		for _, prid := range req.ListServerDataID {
			lib.AppendComma(&baseIn, &baseParam, "?", prid)
		}
		lib.AppendWhereRaw(&baseWhere, "server_data_id IN ("+baseIn+")")
	}

	// runQuery := "SELECT server_data_id, server_id, server_account, item_id, account_id, external_smscount,external_transdate, external_transcount, invoice_id FROM server_data "
	runQuery := "SELECT server_data.server_data_id, server_data.server_id, server_data.server_account, server_data.item_id, server_data.account_id, server_data.external_smscount,server_data.external_transdate, server_data.external_transcount, server_data.invoice_id, item.item_id as tblitem_item_id, item.item_name, item.uom, item.category, item_price.item_id as tblitem_price_item_id, item_price.price, item_price.server_id as tblitem_price_server_id, item_price.account_id as tblitem_price_account_id  FROM server_data JOIN item ON server_data.item_id=item.item_id JOIN item_price ON item.item_id=item_price.item_id AND server_data.server_id=item_price.server_id AND server_data.account_id=item_price.account_id"
	if len(baseWhere) > 0 {
		runQuery += " WHERE " + baseWhere
	}
	logrus.Info("LihatQuery-", runQuery)
	// lib.AppendOrderBy(&runQuery, req.Param.OrderBy, req.Param.OrderDir)
	// lib.AppendLimit(&runQuery, req.Param.Page, req.Param.PerPage)

	resultQuery, _, err = conn.DBAppConn.SelectQueryByFieldNameSlice(runQuery, baseParam...)

	for _, each := range resultQuery {
		single := make(map[string]interface{})
		single["server_data_id"] = each["server_data_id"]
		single["account_id"] = each["account_id"]
		single["server_id"] = each["server_id"]
		single["server_account"] = each["server_account"]
		single["item_id"] = each["item_id"]
		single["account_id"] = each["account_id"]
		single["external_transdate"] = each["external_transdate"]
		// single["external_user_id"] = each["external_user_id"]
		// single["external_sender"] = each["external_sender"]
		// single["external_operatorcode"] = each["external_operatorcode"]
		// single["external_route"] = each["external_route"]
		single["external_smscount"] = each["external_smscount"]
		single["external_transcount"] = each["external_transcount"]
		single["invoice_id"] = each["invoice_id"]
		// single["created_date"] = each["created_date"]

		// findItem, _, errFindItem := conn.DBAppConn.SelectQueryByFieldNameSlice("SELECT item_id, item_name, operator, route, category, uom FROM item WHERE item_id = ?", single["item_id"])
		// if errFindItem != nil {
		// 	return nil, errFindItem
		// }
		item := make(map[string]interface{})
		item["item_id"] = each["tblitem_item_id"]
		item["item_name"] = each["item_name"]
		item["category"] = each["category"]
		item["uom"] = each["uom"]

		// findItemPrice, _, errFindItemPrice := conn.DBAppConn.SelectQueryByFieldNameSlice("SELECT item_id,account_id,server_id,price FROM item_price WHERE item_id = ? AND account_id = ? AND server_id = ?", item["item_id"], req.AccountID, req.ServerID)
		// if errFindItemPrice != nil {
		// 	return nil, errFindItemPrice
		// }

		itemPrice := make(map[string]interface{})
		itemPrice["price"] = each["price"]
		itemPrice["item_id"] = each["tblitem_price_item_id"]
		itemPrice["server_id"] = each["tblitem_price_server_id"]
		itemPrice["account_id"] = each["tblitem_price_account_id"]
		item["item_price"] = itemPrice

		single["item"] = item
		// var category
		result = append(result, single)
	}

	return result, err
}

func InsertServerData(conn *connections.Connections, req datastruct.ServerDataRequest) error {
	var err error

	// -- THIS IS BASIC INSERT EXAMPLE
	// var baseIn string
	// var baseParam []interface{}

	// lib.AppendComma(&baseIn, &baseParam, "?", req.ServerDataID)
	// lib.AppendComma(&baseIn, &baseParam, "?", req.ServerDataName)

	// qry := "INSERT INTO server (serverid, servername) VALUES (" + baseIn + ")"
	// _, _, err = conn.DBAppConn.Exec(qry, baseParam...)

	return err
}

func UpdateServerData(conn *connections.Connections, req datastruct.ServerDataRequest) error {
	var err error

	// -- THIS IS BASIC UPDATE EXAMPLE
	// var baseUp string
	// var baseParam []interface{}

	// lib.AppendComma(&baseUp, &baseParam, "servername = ?", req.ServerDataName)
	// qry := "UPDATE server SET " + baseUp + " WHERE serverid = ?"
	// baseParam = append(baseParam, req.ServerDataID)
	// _, _, err = conn.DBAppConn.Exec(qry, baseParam...)
	return err
}

func DeleteServerData(conn *connections.Connections, req datastruct.ServerDataRequest) error {
	var err error
	// -- THIS IS DELETE LOGIC EXAMPLE
	// qry := "DELETE FROM server WHERE serverid = ?"
	// _, _, err = conn.DBAppConn.Exec(qry, req.ServerDataID)
	return err
}
