package models

import (
	"billingdashboard/connections"
	"billingdashboard/lib"
	"billingdashboard/modules/item-price/datastruct"
	"errors"
	"strconv"
)

func GetSingleItemPrice(stubID string, conn *connections.Connections, req datastruct.ItemPriceRequest) (map[string]string, error) {
	var result map[string]string
	var err error

	// -- EXAMPLE
	// if len(stubID) == 0 {
	// 	stubID = req.ItemPriceID
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

func CheckItemPriceExists(stubID string, conn *connections.Connections) error {
	// var param []interface{}
	// qry := "SELECT COUNT(stubid) FROM stub WHERE stubid = ?"
	// param = append(param, stubID)

	// cnt, _ := conn.DBAppConn.GetFirstData(qry, param...)
	// datacount, _ := strconv.Atoi(cnt)
	// if datacount == 0 {
	// 	return errors.New("ItemPrice ID is not exists")
	// }
	return nil
}

func CheckItemPriceDuplicate(exceptID string, conn *connections.Connections, req datastruct.ItemPriceDataStruct) error {
	var param []interface{}
	qry := "SELECT COUNT(item_id) FROM item_price WHERE item_id = ? AND account_id = ? AND server_id = ?"
	param = append(param, req.ItemID)
	param = append(param, req.AccountID)
	param = append(param, req.ServerID)

	cnt, _ := conn.DBAppConn.GetFirstData(qry, param...)

	datacount, _ := strconv.Atoi(cnt)
	if datacount > 0 {
		return errors.New("ItemPrice is already exists. Please use another server, account, or item")
	}
	return nil
}

func UpdateItemPriceHelper(conn *connections.Connections, req datastruct.ItemPriceDataStruct) error {
	var err error

	// -- THIS IS BASIC UPDATE EXAMPLE
	var baseUp string
	var baseParam []interface{}

	lib.AppendComma(&baseUp, &baseParam, "price = ?", req.Price)
	// lib.AppendComma(&baseUp, &baseParam, "account_id = ?", req.AccountID)
	// lib.AppendComma(&baseUp, &baseParam, "server_id = ?", req.AccountID)
	qry := "UPDATE item_price SET " + baseUp + " WHERE item_id = ? AND account_id = ? AND server_id = ?"
	baseParam = append(baseParam, req.ItemID)
	baseParam = append(baseParam, req.AccountID)
	baseParam = append(baseParam, req.ServerID)
	_, _, err = conn.DBAppConn.Exec(qry, baseParam...)
	return err
}
