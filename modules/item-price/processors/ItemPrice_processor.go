package processors

import (
	"billingdashboard/connections"
	"billingdashboard/modules/item-price/datastruct"
	"billingdashboard/modules/item-price/models"
)

func GetListItemPrice(conn *connections.Connections, req datastruct.ItemPriceRequest) ([]datastruct.ItemPriceDataStruct, error) {
	var output []datastruct.ItemPriceDataStruct
	var err error

	// grab mapping data from model
	itemPriceList, err := models.GetItemPriceFromRequest(conn, req)
	if err != nil {
		return output, err
	}

	for _, itemPrice := range itemPriceList {
		single := CreateSingleItemPriceStruct(itemPrice, req)
		output = append(output, single)
	}

	return output, err
}

func CreateSingleItemPriceStruct(itemPrice map[string]interface{}, req datastruct.ItemPriceRequest) datastruct.ItemPriceDataStruct {
	var single datastruct.ItemPriceDataStruct
	single.ItemID, _ = itemPrice["item_id"].(string)
	single.AccountID, _ = itemPrice["account_id"].(string)
	single.ServerID, _ = itemPrice["server_id"].(string)
	single.Price, _ = itemPrice["price"].(string)
	single.Category = itemPrice["category"].(string)

	if len(req.AccountID) > 0 && len(req.ServerID) > 0 && len(req.ServerID) > 0 && len(req.ItemID) == 0 {
		item := datastruct.ItemDataStruct{
			ItemName: itemPrice["item"].(map[string]interface{})["item_name"].(string),
			UOM:      itemPrice["item"].(map[string]interface{})["uom"].(string),
		}
		single.Item = item
	}

	account := datastruct.AccountDataStruct{
		AccountID:          itemPrice["account"].(map[string]interface{})["account_id"].(string),
		Name:               itemPrice["account"].(map[string]interface{})["name"].(string),
		Status:             itemPrice["account"].(map[string]interface{})["status"].(string),
		CompanyID:          itemPrice["account"].(map[string]interface{})["company_id"].(string),
		AccountType:        itemPrice["account"].(map[string]interface{})["account_type"].(string),
		BillingType:        itemPrice["account"].(map[string]interface{})["billing_type"].(string),
		Address1:           itemPrice["account"].(map[string]interface{})["address1"].(string),
		Address2:           itemPrice["account"].(map[string]interface{})["address2"].(string),
		City:               itemPrice["account"].(map[string]interface{})["city"].(string),
		Phone:              itemPrice["account"].(map[string]interface{})["phone"].(string),
		ContactPerson:      itemPrice["account"].(map[string]interface{})["contact_person"].(string),
		ContactPersonPhone: itemPrice["account"].(map[string]interface{})["contact_person_phone"].(string),
		Desc:               itemPrice["account"].(map[string]interface{})["desc"].(string),
		LastUpdateUsername: itemPrice["account"].(map[string]interface{})["last_update_username"].(string),
		LastUpdateDate:     itemPrice["account"].(map[string]interface{})["last_update_date"].(string),
	}

	server := datastruct.ServerDataStruct{
		ServerID:   itemPrice["server"].(map[string]interface{})["server_id"].(string),
		ServerName: itemPrice["server"].(map[string]interface{})["server_name"].(string),
		ServerUrl:  itemPrice["server"].(map[string]interface{})["server_url"].(string),
	}

	single.Account = account
	single.Server = server

	return single
}

func InsertItemPrice(conn *connections.Connections, req datastruct.ItemPriceRequest) (datastruct.ItemPriceDataStruct, error) {
	var output datastruct.ItemPriceDataStruct
	var err error

	err = models.InsertItemPrice(conn, req)
	if err != nil {
		return output, err
	}

	// jika tidak ada error, return single instance of single stub
	// single, err := models.GetSingleItemPrice(req.ItemPriceID, conn, req)
	// if err != nil {
	// 	return output, err
	// }

	// output = CreateSingleItemPriceStruct(single)
	return output, err
}

func UpdateItemPrice(conn *connections.Connections, req datastruct.ItemPriceRequest) (datastruct.ItemPriceDataStruct, error) {
	var output datastruct.ItemPriceDataStruct
	var err error

	err = models.UpdateItemPrice(conn, req)
	if err != nil {
		return output, err
	}

	// jika tidak ada error, return single instance of single stub
	// single, err := models.GetSingleItemPrice(req.ItemPriceID, conn, req)
	// if err != nil {
	// 	return output, err
	// }

	// output = CreateSingleItemPriceStruct(single)
	return output, err
}

func DeleteItemPrice(conn *connections.Connections, req datastruct.ItemPriceRequest) error {
	err := models.DeleteItemPrice(conn, req)
	return err
}
