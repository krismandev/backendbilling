package models

import (
	"billingdashboard/connections"
	"billingdashboard/modules/item/datastruct"
	"errors"
	"strconv"
)

func GetSingleItem(itemID string, conn *connections.Connections, req datastruct.ItemRequest) (map[string]interface{}, error) {
	var result map[string]interface{}
	var err error

	// -- EXAMPLE
	if len(itemID) == 0 {
		itemID = req.ItemID
	}
	query := "SELECT item_id, item_name, operator, route, category_id FROM item WHERE item_id = ?"
	results, _, err := conn.DBAppConn.SelectQueryByFieldNameSlice(query, itemID)
	if err != nil {
		return result, err
	}

	for _, res := range results {
		result["item_id"] = res["item_id"]
		result["item_name"] = res["item_name"]
		result["operator"] = res["operator"]
		result["route"] = res["route"]
		result["category_id"] = res["category_id"]
		break
	}
	return result, err
}

func CheckItemExists(itemID string, conn *connections.Connections) error {
	var param []interface{}
	qry := "SELECT COUNT(item_id) FROM item WHERE item_id = ?"
	param = append(param, itemID)

	cnt, _ := conn.DBAppConn.GetFirstData(qry, param...)
	datacount, _ := strconv.Atoi(cnt)
	if datacount == 0 {
		return errors.New("Item ID is not exists")
	}
	return nil
}

func CheckItemDuplicate(exceptID string, conn *connections.Connections, req datastruct.ItemRequest) error {
	var param []interface{}
	qry := "SELECT COUNT(item_id) FROM item WHERE operator = ? AND route = ? AND category = ?"
	param = append(param, req.Operator)
	param = append(param, req.Route)
	param = append(param, req.Category)

	// if len(exceptID) > 0 {
	// 	qry += " AND item_id <> ?"
	// 	param = append(param, exceptID)
	// }

	cnt, _ := conn.DBAppConn.GetFirstData(qry, param...)
	datacount, _ := strconv.Atoi(cnt)
	if datacount > 0 {
		return errors.New("Item is already exists. Please use another route, category or operator")
	}
	return nil
}
