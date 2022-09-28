package models

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/category/datastruct"
)

func GetSingleCategory(categoryID string, conn *connections.Connections, req datastruct.CategoryRequest) (map[string]string, error) {
	var result map[string]string
	var err error

	// -- EXAMPLE
	// if len(categoryID) == 0 {
	// 	categoryID = req.CategoryID
	// }
	// query := "SELECT categoryid, categoryname FROM category WHERE categoryid = ?"
	// results, _, err := conn.DBAppConn.SelectQueryByFieldNameSlice(query, categoryID)
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

func CheckCategoryExists(categoryID string, conn *connections.Connections) error {
	// var param []interface{}
	// qry := "SELECT COUNT(categoryid) FROM category WHERE categoryid = ?"
	// param = append(param, categoryID)

	// cnt, _ := conn.DBAppConn.GetFirstData(qry, param...)
	// datacount, _ := strconv.Atoi(cnt)
	// if datacount == 0 {
	// 	return errors.New("Category ID is not exists")
	// }
	return nil
}

func CheckCategoryDuplicate(exceptID string, conn *connections.Connections, req datastruct.CategoryRequest) error {
	// var param []interface{}
	// qry := "SELECT COUNT(categoryid) FROM category WHERE categoryid = ?"
	// param = append(param, req.CategoryID)
	// if len(exceptID) > 0 {
	// 	qry += " AND categoryid <> ?"
	// 	param = append(param, exceptID)
	// }

	// cnt, _ := conn.DBAppConn.GetFirstData(qry, param...)
	// datacount, _ := strconv.Atoi(cnt)
	// if datacount > 0 {
	// 	return errors.New("Category ID is already exists. Please use another Category ID")
	// }
	return nil
}
