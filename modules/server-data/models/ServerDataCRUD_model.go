package models

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/lib"
	"backendbillingdashboard/modules/server-data/datastruct"

	log "github.com/sirupsen/logrus"
)

func GetServerDataFromRequest(conn *connections.Connections, req datastruct.ServerDataRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	var resultQuery []map[string]string
	var err error

	// -- THIS IS BASIC GET REQUEST EXAMPLE LOGIC
	var baseWhere string
	var baseParam []interface{}

	baseParam = append(baseParam, req.CurrencyCode)
	lib.AppendWhere(&baseWhere, &baseParam, "server_data_id = ?", req.ServerDataID)
	lib.AppendWhere(&baseWhere, &baseParam, "server_data.server_id = ?", req.ServerID)
	lib.AppendWhere(&baseWhere, &baseParam, "server_data.account_id = ?", req.AccountID)
	lib.AppendWhere(&baseWhere, &baseParam, "server_data.external_sender = ?", req.ExternalSender)
	lib.AppendWhere(&baseWhere, &baseParam, "server_data.external_rootparent_account = ?", req.ExternalRootParentAccount)
	lib.AppendWhere(&baseWhere, &baseParam, "DATE_FORMAT(server_data.external_transdate, '%Y%m') = ?", req.MonthUse)
	lib.AppendWhereRaw(&baseWhere, "server_data.invoice_id IS NULL")
	// lib.AppendWhere(&baseWhere, &baseParam, "item_price.currency_code = ?", req.CurrencyCode)

	if len(req.ListExternalRootParentAccount) > 0 {
		var baseIn string
		for _, extrootaccountid := range req.ListExternalRootParentAccount {
			lib.AppendComma(&baseIn, &baseParam, "?", extrootaccountid)
		}
		lib.AppendWhereRaw(&baseWhere, "external_rootparent_account IN ("+baseIn+")")
	}

	resultCurrencey, _, errGetCurrency := conn.DBOcsConn.SelectQueryByFieldNameSlice("SELECT balance_type, balance_name, exponent, balance_category FROM ocs.balance WHERE balance_category = ?", "C")
	if errGetCurrency != nil {
		return result, errGetCurrency
	}
	// log.Info("LihatResultCurrency-", resultCurrencey)

	currencyIn := ""
	for i, each := range resultCurrencey {
		currencyIn = currencyIn + "'" + each["balance_type"] + "'"
		if i != len(resultCurrencey)-1 {
			currencyIn = currencyIn + ", "
		}
	}

	log.Info("LihatCurrency-", currencyIn)
	// runQuery := "SELECT server_data_id, server_id, server_account, item_id, account_id, external_smscount,external_transdate, external_transcount, invoice_id FROM server_data "
	runQuery := `SELECT server_data.server_data_id, server_data.server_id, server_data.external_account_id, 
	server_data.external_rootparent_account, server_data.item_id, server_data.account_id, server_data.external_smscount,server_data.external_transdate, 
	server_data.external_transcount, server_data.external_balance_type, server_data.invoice_id,server_data.external_user_id, 
	server_data.external_sender, server_data.external_operatorcode, server_data.external_route, 
	item.item_id as tblitem_item_id, item.item_name, item.uom, item.category, item_price.item_id as tblitem_price_item_id, 
	IF((server_data.external_price IS NOT NULL && server_data.external_price <> 0 && server_data.external_balance_type in (` + currencyIn + `) ),
	server_data.external_price, item_price.price) as price, item_price.server_id as tblitem_price_server_id, item_price.account_id as tblitem_price_account_id, 
	server.server_name as tblserver_server_name, mapping_operator.new_route 
	FROM server_data 
	LEFT JOIN item ON server_data.item_id=item.item_id 
	LEFT JOIN server ON server.server_id = server_data.server_id 
	LEFT JOIN item_price ON item.item_id=item_price.item_id AND server_data.server_id=item_price.server_id 
		AND server_data.account_id=item_price.account_id AND item_price.currency_code = ? 
	LEFT JOIN mapping_operator ON server_data.external_operatorcode = mapping_operator.operatorcode AND server_data.external_route = mapping_operator.route `
	if len(baseWhere) > 0 {
		runQuery += " WHERE " + baseWhere
	}
	log.Info("final query:", runQuery)
	lib.AppendOrderBy(&runQuery, req.Param.OrderBy, req.Param.OrderDir)
	lib.AppendLimit(&runQuery, req.Param.Page, req.Param.PerPage)

	resultQuery, _, err = conn.DBAppConn.SelectQueryByFieldNameSlice(runQuery, baseParam...)

	for _, each := range resultQuery {
		single := make(map[string]interface{})
		single["server_data_id"] = each["server_data_id"]
		single["account_id"] = each["account_id"]
		single["server_id"] = each["server_id"]
		single["external_account_id"] = each["external_account_id"]
		single["external_rootparent_account"] = each["external_rootparent_account"]
		single["item_id"] = each["item_id"]
		single["account_id"] = each["account_id"]
		single["external_transdate"] = each["external_transdate"]
		single["external_price"] = each["price"]
		single["external_balance_type"] = each["external_balance_type"]
		single["external_user_id"] = each["external_user_id"]
		single["external_sender"] = each["external_sender"]
		single["external_operatorcode"] = each["external_operatorcode"]
		single["external_route"] = each["external_route"]
		single["external_smscount"] = each["external_smscount"]
		single["external_transcount"] = each["external_transcount"]
		single["invoice_id"] = each["invoice_id"]
		single["new_route"] = each["new_route"]
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

		server := make(map[string]interface{})
		server["server_id"] = each["server_id"]
		server["server_name"] = each["tblserver_server_name"]

		single["server"] = server

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

func GetSenderFromRequest(conn *connections.Connections, req datastruct.ServerDataRequest) ([]map[string]string, error) {
	var result []map[string]string
	var err error

	// -- THIS IS BASIC GET REQUEST EXAMPLE LOGIC
	var baseWhere string
	var baseParam []interface{}

	lib.AppendWhere(&baseWhere, &baseParam, "account_id = ?", req.AccountID)
	lib.AppendWhere(&baseWhere, &baseParam, "external_rootparent_account = ?", req.ExternalRootParentAccount)
	lib.AppendWhere(&baseWhere, &baseParam, "DATE_FORMAT(external_transdate, '%Y%m') = ?", req.MonthUse)

	runQuery := "SELECT distinct(external_sender) FROM server_data "
	if len(baseWhere) > 0 {
		runQuery += "WHERE " + baseWhere
	}
	lib.AppendOrderBy(&runQuery, req.Param.OrderBy, req.Param.OrderDir)
	lib.AppendLimit(&runQuery, req.Param.Page, req.Param.PerPage)

	log.Info("LihatQuery-", runQuery)

	result, _, err = conn.DBAppConn.SelectQueryByFieldNameSlice(runQuery, baseParam...)
	log.Info("LihatSender-", result)
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
