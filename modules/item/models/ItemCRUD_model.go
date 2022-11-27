package models

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/lib"
	"backendbillingdashboard/modules/item/datastruct"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func GetItemFromRequest(conn *connections.Connections, req datastruct.ItemRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	var resultQuery []map[string]string
	var err error

	// -- THIS IS BASIC GET REQUEST EXAMPLE LOGIC
	var baseWhere string
	var baseParam []interface{}

	lib.AppendWhere(&baseWhere, &baseParam, "item_id = ?", req.ItemID)
	lib.AppendWhere(&baseWhere, &baseParam, "item_name = ?", req.ItemName)
	lib.AppendWhere(&baseWhere, &baseParam, "category = ?", req.Category)
	lib.AppendWhere(&baseWhere, &baseParam, "operator = ?", req.Operator)
	lib.AppendWhere(&baseWhere, &baseParam, "route = ?", req.Route)
	if len(req.ListItem) > 0 {
		var baseIn string
		// var whereInCol string
		for index, item := range req.ListItem {
			var eachElement string
			// baseIn += "("
			eachElement = "(?,?)"
			baseParam = append(baseParam, item.Operator)
			baseParam = append(baseParam, item.Route)

			// lib.AppendComma(&eachElement, &baseParam, "?", item.Operator)
			// lib.AppendComma(&eachElement, &baseParam, "?", item.Route)
			if index != len(req.ListItem)-1 {
				eachElement += ","
			}

			baseIn += eachElement
		}
		lib.AppendWhereRaw(&baseWhere, "(operator, route) IN ("+baseIn+")")
	}

	runQuery := "SELECT item_id, item_name, operator, route, uom, category FROM item "
	if len(baseWhere) > 0 {
		runQuery += "WHERE " + baseWhere
	}
	lib.AppendOrderBy(&runQuery, req.Param.OrderBy, req.Param.OrderDir)
	lib.AppendLimit(&runQuery, req.Param.Page, req.Param.PerPage)

	resultQuery, _, err = conn.DBAppConn.SelectQueryByFieldNameSlice(runQuery, baseParam...)
	for _, each := range resultQuery {
		single := make(map[string]interface{})
		single["item_id"] = each["item_id"]
		single["item_name"] = each["item_name"]
		single["operator"] = each["operator"]
		single["route"] = each["route"]
		single["category"] = each["category"]
		single["uom"] = each["uom"]
		// var category
		result = append(result, single)
	}
	return result, err
}

func InsertItem(conn *connections.Connections, req datastruct.ItemRequest) error {
	var err error

	// -- THIS IS BASIC INSERT EXAMPLE
	var baseIn string
	var baseParam []interface{}

	lastId, _ := conn.DBAppConn.GetFirstData("SELECT last_id FROM control_id where control_id.key=?", "item")
	// lastId, _ := conn.DBAppConn.GetFirstData("SELECT last_id FROM control_id where control_id.key=?", "account")

	intLastId, err := strconv.Atoi(lastId)
	insertId := intLastId + 1

	insertIdString := strconv.Itoa(insertId)

	lib.AppendComma(&baseIn, &baseParam, "?", insertIdString)
	lib.AppendComma(&baseIn, &baseParam, "?", req.ItemName)
	lib.AppendComma(&baseIn, &baseParam, "?", req.Operator)
	lib.AppendComma(&baseIn, &baseParam, "?", req.Route)
	lib.AppendComma(&baseIn, &baseParam, "?", req.Category)

	qry := "INSERT INTO item (item_id, item_name,operator, route, category) VALUES (" + baseIn + ")"
	_, _, errInsert := conn.DBAppConn.Exec(qry, baseParam...)

	if errInsert != nil {
		return errInsert
	}

	_, _, err = conn.DBAppConn.Exec("UPDATE control_id set last_id=? where control_id.key=?", insertIdString, "item")

	return err
}

func UpdateItem(conn *connections.Connections, req datastruct.ItemRequest) error {
	var err error

	// -- THIS IS BASIC UPDATE EXAMPLE
	var baseUp string
	var baseParam []interface{}

	lib.AppendComma(&baseUp, &baseParam, "item_name = ?", req.ItemName)
	lib.AppendComma(&baseUp, &baseParam, "route = ?", req.Route)
	lib.AppendComma(&baseUp, &baseParam, "operator = ?", req.Operator)
	lib.AppendComma(&baseUp, &baseParam, "operator = ?", req.Operator)
	lib.AppendComma(&baseUp, &baseParam, "category = ?", req.Category)
	qry := "UPDATE item SET " + baseUp + " WHERE item_id = ?"
	baseParam = append(baseParam, req.ItemID)
	_, _, err = conn.DBAppConn.Exec(qry, baseParam...)
	log.Info("UpdateParam - ", baseParam)
	return err
}

func DeleteItem(conn *connections.Connections, req datastruct.ItemRequest) error {
	var err error
	// -- THIS IS DELETE LOGIC EXAMPLE
	qry := "UPDATE item SET status = A WHERE item_id = ?"
	_, _, err = conn.DBAppConn.Exec(qry, req.ItemID)
	return err
}
