package processors

import (
	"billingdashboard/connections"
	"billingdashboard/modules/server-data/datastruct"
	"billingdashboard/modules/server-data/models"
)

func GetListServerData(conn *connections.Connections, req datastruct.ServerDataRequest) ([]datastruct.ServerDataDataStruct, error) {
	var output []datastruct.ServerDataDataStruct
	var err error

	// grab mapping data from model
	serverDataList, err := models.GetServerDataFromRequest(conn, req)
	if err != nil {
		return output, err
	}

	for _, serverData := range serverDataList {
		single := CreateSingleServerDataStruct(serverData)
		output = append(output, single)
	}

	return output, err
}

func CreateSingleServerDataStruct(serverData map[string]interface{}) datastruct.ServerDataDataStruct {
	var single datastruct.ServerDataDataStruct
	single.ServerDataID, _ = serverData["server_data_id"].(string)
	single.ServerID, _ = serverData["server_id"].(string)
	single.ExternalAccountID, _ = serverData["external_account_id"].(string)
	single.ItemID, _ = serverData["item_id"].(string)
	single.AccountID, _ = serverData["account_id"].(string)
	single.ExternalTransdate, _ = serverData["external_transdate"].(string)
	single.ExternalRootParentAccount, _ = serverData["external_rootparent_account"].(string)
	single.ExternalPrice, _ = serverData["external_price"].(string)
	// single.ExternalUserID, _ = serverData["external_user_id"].(string)
	// single.ExternalSender, _ = serverData["external_sender"].(string)
	// single.ExternalOperatorCode, _ = serverData["external_operatorcode"].(string)
	// single.ExternalRoute, _ = serverData["external_route"].(string)
	single.ExternalSMSCount, _ = serverData["external_smscount"].(string)
	single.ExternalTransCount, _ = serverData["external_transcount"].(string)
	single.ExternalBalanceType, _ = serverData["external_balance_type"].(string)
	single.InvoiceID, _ = serverData["invoice_id"].(string)

	itemPrice := datastruct.ItemPriceDataStruct{
		ItemID:    serverData["item"].(map[string]interface{})["item_price"].(map[string]interface{})["item_id"].(string),
		AccountID: serverData["item"].(map[string]interface{})["item_price"].(map[string]interface{})["account_id"].(string),
		ServerID:  serverData["item"].(map[string]interface{})["item_price"].(map[string]interface{})["server_id"].(string),
		Price:     serverData["item"].(map[string]interface{})["item_price"].(map[string]interface{})["price"].(string),
	}

	item := datastruct.ItemDataStruct{
		ItemID:    serverData["item"].(map[string]interface{})["item_id"].(string),
		ItemName:  serverData["item"].(map[string]interface{})["item_name"].(string),
		UOM:       serverData["item"].(map[string]interface{})["uom"].(string),
		Category:  serverData["item"].(map[string]interface{})["category"].(string),
		ItemPrice: itemPrice,
	}

	single.Item = item
	return single
}

func InsertServerData(conn *connections.Connections, req datastruct.ServerDataRequest) (datastruct.ServerDataDataStruct, error) {
	var output datastruct.ServerDataDataStruct
	var err error

	err = models.InsertServerData(conn, req)
	if err != nil {
		return output, err
	}

	// jika tidak ada error, return single instance of single server-data
	// single, err := models.GetSingleServerData(req.ServerDataID, conn, req)
	// if err != nil {
	// 	return output, err
	// }

	// output = CreateSingleServerDataStruct(single)
	return output, err
}

func UpdateServerData(conn *connections.Connections, req datastruct.ServerDataRequest) (datastruct.ServerDataDataStruct, error) {
	var output datastruct.ServerDataDataStruct
	var err error

	err = models.UpdateServerData(conn, req)
	if err != nil {
		return output, err
	}

	// jika tidak ada error, return single instance of single server-data
	// single, err := models.GetSingleServerData(req.ServerDataID, conn, req)
	// if err != nil {
	// 	return output, err
	// }

	// output = CreateSingleServerDataStruct(single)
	return output, err
}

func DeleteServerData(conn *connections.Connections, req datastruct.ServerDataRequest) error {
	err := models.DeleteServerData(conn, req)
	return err
}
