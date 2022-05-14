package models

import (
	"billingdashboard/connections"
	"billingdashboard/lib"
	"billingdashboard/modules/item-price/datastruct"
	"strings"

	"github.com/sirupsen/logrus"
)

func GetItemPriceFromRequest(conn *connections.Connections, req datastruct.ItemPriceRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	var err error

	// -- THIS IS BASIC GET REQUEST EXAMPLE LOGIC
	var baseWhere string
	var baseParam []interface{}

	lib.AppendWhere(&baseWhere, &baseParam, "item_price.item_id = ?", req.ItemID)
	lib.AppendWhere(&baseWhere, &baseParam, "item_price.account_id = ?", req.AccountID)
	lib.AppendWhere(&baseWhere, &baseParam, "item_price.server_id = ?", req.ServerID)
	lib.AppendWhere(&baseWhere, &baseParam, "item.category = ?", req.Category)

	if len(req.ListAccountID) > 0 {
		var listAccountParam string
		for _, prid := range req.ListAccountID {
			lib.AppendComma(&listAccountParam, &baseParam, "?", prid)
		}
		lib.AppendWhereRaw(&baseWhere, "account_id IN ("+listAccountParam+")")
	}

	if len(req.ListServerID) > 0 {
		var listServerParam string
		for _, prid := range req.ListServerID {
			lib.AppendComma(&listServerParam, &baseParam, "?", prid)
		}
		lib.AppendWhereRaw(&baseWhere, "server_id IN ("+listServerParam+")")
	}

	if len(req.ListItemID) > 0 {
		var listItemParam string
		for _, prid := range req.ListItemID {
			lib.AppendComma(&listItemParam, &baseParam, "?", prid)
		}
		lib.AppendWhereRaw(&baseWhere, "item_price.item_id IN ("+listItemParam+")")
	}

	// SELECT distinct item_price.account_id, item_price.server_id, item.category from item_price INNER JOIN item on item.item_id=item_price.item_id;

	runQuery := "SELECT distinct item_price.account_id, item_price.server_id, item.category from item_price INNER JOIN item on item.item_id=item_price.item_id"
	if len(baseWhere) > 0 {
		runQuery = "SELECT item_price.item_id, item_price.account_id, item_price.server_id, item_price.price from item_price INNER JOIN item ON item.item_id=item_price.item_id"
		if len(req.AccountID) > 0 && len(req.ServerID) > 0 && len(req.ServerID) > 0 && len(req.ItemID) == 0 {
			runQuery = "SELECT item_price.item_id, item_price.account_id, item_price.server_id, item_price.price, item.item_name, item.uom from item_price INNER JOIN item ON item.item_id=item_price.item_id"
		}
		runQuery += " WHERE " + baseWhere
	}

	logrus.Info("FinalQuery", runQuery)
	// lib.AppendOrderBy(&runQuery, "item_price.item_id", req.Param.OrderDir)
	lib.AppendLimit(&runQuery, req.Param.Page, req.Param.PerPage)

	resultSelect, _, errSelect := conn.DBAppConn.SelectQueryByFieldNameSlice(runQuery, baseParam...)
	if errSelect != nil {
		return nil, errSelect
	}
	for _, each := range resultSelect {
		single := make(map[string]interface{})

		single["item_id"] = each["item_id"]
		single["price"] = each["price"]
		single["account_id"] = each["account_id"]
		single["server_id"] = each["server_id"]
		single["category"] = each["category"]

		if len(req.AccountID) > 0 && len(req.ServerID) > 0 && len(req.ServerID) > 0 && len(req.ItemID) == 0 {
			item := make(map[string]interface{})
			item["item_name"] = each["item_name"]
			item["uom"] = each["uom"]
			single["item"] = item
		}
		findAccount, _, errFindAccount := conn.DBAppConn.SelectQueryByFieldNameSlice("SELECT account_id, name, status, company_id, address1, address2, account_type, billing_type,city, phone, contact_person, contact_person_phone ,account.desc, last_update_username, last_update_date FROM account WHERE account_id = ?", single["account_id"])
		if errFindAccount != nil {
			return nil, errFindAccount
		}

		account := make(map[string]interface{})
		account["account_id"] = findAccount[0]["account_id"]
		account["name"] = findAccount[0]["name"]
		account["status"] = findAccount[0]["status"]
		account["company_id"] = findAccount[0]["company_id"]
		account["account_type"] = findAccount[0]["account_type"]
		account["billing_type"] = findAccount[0]["billing_type"]
		account["address1"] = findAccount[0]["address1"]
		account["address2"] = findAccount[0]["address2"]
		account["city"] = findAccount[0]["city"]
		account["phone"] = findAccount[0]["phone"]
		account["contact_person"] = findAccount[0]["contact_person"]
		account["contact_person_phone"] = findAccount[0]["contact_person_phone"]
		account["desc"] = findAccount[0]["desc"]
		account["last_update_username"] = findAccount[0]["last_update_username"]
		account["last_update_date"] = findAccount[0]["last_update_date"]
		single["account"] = account

		findServer, _, errFindServer := conn.DBAppConn.SelectQueryByFieldNameSlice("SELECT server_id, server_name, server_url FROM server WHERE server_id = ?", single["server_id"])
		if errFindServer != nil {
			return nil, errFindServer
		}

		server := make(map[string]interface{})
		server["server_id"] = findServer[0]["server_id"]
		server["server_name"] = findServer[0]["server_name"]
		server["server_url"] = findServer[0]["server_url"]

		single["server"] = server

		result = append(result, single)
	}

	return result, err
}

// func GetItemPriceFromRequest(conn *connections.Connections, req datastruct.ItemPriceRequest) ([]map[string]string, error) {
// 	var result []map[string]string
// 	var err error

// 	// -- THIS IS BASIC GET REQUEST EXAMPLE LOGIC
// 	var baseWhere string
// 	var baseParam []interface{}

// 	lib.AppendWhere(&baseWhere, &baseParam, "item_id = ?", req.ItemID)
// 	lib.AppendWhere(&baseWhere, &baseParam, "account_id = ?", req.AccountID)
// 	lib.AppendWhere(&baseWhere, &baseParam, "server_id = ?", req.ServerID)

// 	runQuery := "SELECT item_id, account_id, price, server_id FROM item_price "
// 	if len(baseWhere) > 0 {
// 		runQuery += "WHERE " + baseWhere
// 	}
// 	lib.AppendOrderBy(&runQuery, req.Param.OrderBy, req.Param.OrderDir)
// 	lib.AppendLimit(&runQuery, req.Param.Page, req.Param.PerPage)

// 	result, _, err = conn.DBAppConn.SelectQueryByFieldNameSlice(runQuery, baseParam...)
// 	return result, err
// }

func InsertItemPrice(conn *connections.Connections, req datastruct.ItemPriceRequest) error {
	var err error

	// -- THIS IS BASIC INSERT EXAMPLE
	var baseIn string
	var baseParam []interface{}

	if len(req.ListItemPrice) == 0 {
		lib.AppendComma(&baseIn, &baseParam, "?", req.ItemID)
		lib.AppendComma(&baseIn, &baseParam, "?", req.AccountID)
		lib.AppendComma(&baseIn, &baseParam, "?", req.Price)
		lib.AppendComma(&baseIn, &baseParam, "?", req.ServerID)
		lib.AppendComma(&baseIn, &baseParam, "?", "0")

		qry := "INSERT INTO item_price (item_id, account_id,price,server_id,tiering) VALUES (" + baseIn + ")"
		_, _, err = conn.DBAppConn.Exec(qry, baseParam...)
	} else if len(req.ListItemPrice) != 0 {

		for _, each := range req.ListItemPrice {
			var baseInputForList string
			var baseParamForList []interface{}

			// lastId, _ := conn.DBAppConn.GetFirstData("SELECT last_id FROM control_id where control_id.key=?", "item_price")
			// // lastId, _ := conn.DBAppConn.GetFirstData("SELECT last_id FROM control_id where control_id.key=?", "account")

			// intLastId, errLastId := strconv.Atoi(lastId)
			// if errLastId != nil {
			// 	return errLastId
			// }
			// insertId := intLastId + 1

			// insertIdString := strconv.Itoa(insertId)

			// lib.AppendComma(&baseInputForList, &baseParamForList, "?", insertIdString)
			lib.AppendComma(&baseInputForList, &baseParamForList, "?", each.ItemID)
			lib.AppendComma(&baseInputForList, &baseParamForList, "?", req.AccountID)
			lib.AppendComma(&baseInputForList, &baseParamForList, "?", each.Price)
			lib.AppendComma(&baseInputForList, &baseParamForList, "?", req.ServerID)
			lib.AppendComma(&baseInputForList, &baseParamForList, "?", "0")

			errCheck := CheckItemPriceDuplicate("", conn, each)
			if errCheck != nil {
				errUpdate := UpdateItemPriceHelper(conn, each)
				if errUpdate != nil {
					return errUpdate
				}
			} else {
				qry := "INSERT INTO item_price (item_id, account_id,price,server_id,tiering) VALUES (" + baseInputForList + ")"
				_, _, errInsert := conn.DBAppConn.Exec(qry, baseParamForList...)
				if errInsert != nil {
					return errInsert
				}

			}

		}

	}

	return err
}

func UpdateItemPrice(conn *connections.Connections, req datastruct.ItemPriceRequest) error {
	var err error

	// -- THIS IS BASIC INSERT EXAMPLE
	var baseIn string
	var baseParam []interface{}

	if len(req.ListItemPrice) == 0 {
		lib.AppendComma(&baseIn, &baseParam, "?", req.ItemID)
		lib.AppendComma(&baseIn, &baseParam, "?", req.AccountID)
		lib.AppendComma(&baseIn, &baseParam, "?", req.Price)
		lib.AppendComma(&baseIn, &baseParam, "?", req.ServerID)

		qry := "INSERT INTO item_price (item_id, account_id,price,server_id) VALUES (" + baseIn + ")"
		_, _, err = conn.DBAppConn.Exec(qry, baseParam...)
	} else if len(req.ListItemPrice) != 0 {

		for _, each := range req.ListItemPrice {
			var baseInputForList string
			var baseParamForList []interface{}

			// lastId, _ := conn.DBAppConn.GetFirstData("SELECT last_id FROM control_id where control_id.key=?", "item_price")
			// // lastId, _ := conn.DBAppConn.GetFirstData("SELECT last_id FROM control_id where control_id.key=?", "account")

			// intLastId, errLastId := strconv.Atoi(lastId)
			// if errLastId != nil {
			// 	return errLastId
			// }
			// insertId := intLastId + 1

			// insertIdString := strconv.Itoa(insertId)

			// lib.AppendComma(&baseInputForList, &baseParamForList, "?", insertIdString)
			lib.AppendComma(&baseInputForList, &baseParamForList, "?", each.ItemID)
			lib.AppendComma(&baseInputForList, &baseParamForList, "?", req.AccountID)
			lib.AppendComma(&baseInputForList, &baseParamForList, "?", each.Price)
			lib.AppendComma(&baseInputForList, &baseParamForList, "?", req.ServerID)

			errCheck := CheckItemPriceDuplicate("", conn, each)
			if errCheck != nil {
				errUpdate := UpdateItemPriceHelper(conn, each)
				if errUpdate != nil {
					return errUpdate
				}
			} else {
				qry := "INSERT INTO item_price (item_id, account_id,price,server_id) VALUES (" + baseInputForList + ")"
				_, _, errInsert := conn.DBAppConn.Exec(qry, baseParamForList...)
				if errInsert != nil {
					return errInsert
				}

			}

		}

	}

	return err
}

func BulkUpdateItemPrice(conn *connections.Connections, req datastruct.ItemPriceRequest) error {
	var err error

	// -- THIS IS BASIC INSERT EXAMPLE
	var baseIn string
	var baseParam []interface{}

	if len(req.ListItemPrice) == 0 {
		lib.AppendComma(&baseIn, &baseParam, "?", req.ItemID)
		lib.AppendComma(&baseIn, &baseParam, "?", req.AccountID)
		lib.AppendComma(&baseIn, &baseParam, "?", req.Price)
		lib.AppendComma(&baseIn, &baseParam, "?", req.ServerID)

		qry := "INSERT INTO item_price (item_id, account_id,price,server_id) VALUES (" + baseIn + ")"
		_, _, err = conn.DBAppConn.Exec(qry, baseParam...)
	} else if len(req.ListItemPrice) != 0 {

		bulkUpdateQuery := "INSERT INTO item_price (item_price.item_id, item_price.account_id,item_price.price,item_price.server_id,item_price.tiering, item_price.last_update_username) VALUES"
		var paramsBulkUpdate []interface{}
		var stringGroup []string

		for _, each := range req.ListItemPrice {

			partquery := "(?, ?, ?, ?, ?, ?)"
			paramsBulkUpdate = append(paramsBulkUpdate, each.ItemID)
			paramsBulkUpdate = append(paramsBulkUpdate, each.AccountID)
			paramsBulkUpdate = append(paramsBulkUpdate, each.Price)
			paramsBulkUpdate = append(paramsBulkUpdate, each.ServerID)
			paramsBulkUpdate = append(paramsBulkUpdate, "0")
			paramsBulkUpdate = append(paramsBulkUpdate, req.LastUpdateUsername)
			stringGroup = append(stringGroup, partquery)

		}
		final_query := bulkUpdateQuery + strings.Join(stringGroup, ", ") + " ON DUPLICATE KEY UPDATE price = VALUES(item_price.price), last_update_username = VALUES(item_price.last_update_username)"
		logrus.Info("FinalQuery", final_query)
		_, _, errInsert := conn.DBAppConn.Exec(final_query, paramsBulkUpdate...)
		if errInsert != nil {
			logrus.Error("Error BulkUpdateItemPrice Sender SMSC : " + errInsert.Error())
			return errInsert
		}

	}

	return err
}

func DeleteItemPrice(conn *connections.Connections, req datastruct.ItemPriceRequest) error {
	var err error
	// -- THIS IS DELETE LOGIC EXAMPLE
	qry := "DELETE FROM item_price WHERE item_id = ? AND account_id = ? AND server_id = ?"
	_, _, err = conn.DBAppConn.Exec(qry, req.ItemID, req.AccountID, req.ServerID)
	return err
}
