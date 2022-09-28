package processors

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/category/datastruct"
	"backendbillingdashboard/modules/category/models"
)

func GetListCategory(conn *connections.Connections, req datastruct.CategoryRequest) ([]datastruct.CategoryDataStruct, error) {
	var output []datastruct.CategoryDataStruct
	var err error

	// grab mapping data from model
	categoryList, err := models.GetCategoryFromRequest(conn, req)
	if err != nil {
		return output, err
	}

	for _, category := range categoryList {
		single := CreateSingleCategoryStruct(category)
		output = append(output, single)
	}

	return output, err
}

func CreateSingleCategoryStruct(category map[string]string) datastruct.CategoryDataStruct {
	var single datastruct.CategoryDataStruct
	single.CategoryID = category["category_id"]
	single.CategoryName = category["name"]

	return single
}
