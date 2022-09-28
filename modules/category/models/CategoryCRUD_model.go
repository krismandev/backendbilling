package models

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/lib"
	"backendbillingdashboard/modules/category/datastruct"
)

func GetCategoryFromRequest(conn *connections.Connections, req datastruct.CategoryRequest) ([]map[string]string, error) {
	var result []map[string]string
	var err error

	// -- THIS IS BASIC GET REQUEST EXAMPLE LOGIC
	var baseWhere string
	var baseParam []interface{}

	lib.AppendWhere(&baseWhere, &baseParam, "category_id = ?", req.CategoryID)

	runQuery := "SELECT category_id, name FROM category "
	if len(baseWhere) > 0 {
		runQuery += "WHERE " + baseWhere
	}
	lib.AppendOrderBy(&runQuery, req.Param.OrderBy, req.Param.OrderDir)
	lib.AppendLimit(&runQuery, req.Param.Page, req.Param.PerPage)

	result, _, err = conn.DBAppConn.SelectQueryByFieldNameSlice(runQuery, baseParam...)
	return result, err
}
