package processors

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/item/datastruct"
	"backendbillingdashboard/modules/item/models"
)

func GetListItem(conn *connections.Connections, req datastruct.ItemRequest) ([]datastruct.ItemDataStruct, error) {
	var output []datastruct.ItemDataStruct
	var err error

	// grab mapping data from model
	itemList, err := models.GetItemFromRequest(conn, req)
	if err != nil {
		return output, err
	}

	for _, item := range itemList {
		single := CreateSingleItemStruct(item)
		output = append(output, single)
	}

	return output, err
}

func CreateSingleItemStruct(item map[string]interface{}) datastruct.ItemDataStruct {
	var single datastruct.ItemDataStruct
	single.ItemID, _ = item["item_id"].(string)
	single.ItemName, _ = item["item_name"].(string)
	single.Operator, _ = item["operator"].(string)
	single.Route, _ = item["route"].(string)
	single.Category, _ = item["category"].(string)
	single.UOM, _ = item["uom"].(string)
	// category := datastruct.Category{
	// 	CategoryId:         item["category"].(map[string]interface{})["category_id"].(string),
	// 	CategoryName:       item["category"].(map[string]interface{})["name"].(string),
	// 	LastUpdateUsername: item["category"].(map[string]interface{})["last_update_username"].(string),
	// 	LastUdpateDate:     item["category"].(map[string]interface{})["last_update_date"].(string),
	// }

	// logrus.Info("TestCoba", category)
	// single.Category = category
	// category.CategoryId = item["category"]["category_id"]
	// single.Category =

	return single
}

func InsertItem(conn *connections.Connections, req datastruct.ItemRequest) error {
	var err error

	err = models.InsertItem(conn, req)
	if err != nil {
		return err
	}

	return err
}

func UpdateItem(conn *connections.Connections, req datastruct.ItemRequest) (datastruct.ItemDataStruct, error) {
	var output datastruct.ItemDataStruct
	var err error

	err = models.UpdateItem(conn, req)
	if err != nil {
		return output, err
	}

	// jika tidak ada error, return single instance of single item
	// single, err := models.GetSingleItem(req.ItemID, conn, req)
	// if err != nil {
	// 	return output, err
	// }

	// output = CreateSingleItemStruct(single)
	return output, err
}

func DeleteItem(conn *connections.Connections, req datastruct.ItemRequest) error {
	err := models.DeleteItem(conn, req)
	return err
}
