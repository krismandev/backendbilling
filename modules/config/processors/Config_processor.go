package processors

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/config/datastruct"
	"backendbillingdashboard/modules/config/models"
)

func GetListConfig(conn *connections.Connections, req datastruct.ConfigRequest) ([]datastruct.ConfigDataStruct, error) {
	var output []datastruct.ConfigDataStruct
	var err error

	// grab mapping data from model
	configList, err := models.GetConfigFromRequest(conn, req)
	if err != nil {
		return output, err
	}

	for _, config := range configList {
		single := CreateSingleConfigStruct(config)
		output = append(output, single)
	}

	return output, err
}

func CreateSingleConfigStruct(config map[string]string) datastruct.ConfigDataStruct {
	var single datastruct.ConfigDataStruct
	single.Key, _ = config["key"]
	single.Type, _ = config["type"]
	single.Value, _ = config["value"]

	return single
}

// func InsertConfig(conn *connections.Connections, req datastruct.ConfigRequest) (datastruct.ConfigDataStruct, error) {
// 	var output datastruct.ConfigDataStruct
// 	var err error

// 	err = models.InsertConfig(conn, req)
// 	if err != nil {
// 		return output, err
// 	}

// 	// jika tidak ada error, return single instance of single config
// 	single, err := models.GetSingleConfig(req.Key, conn, req)
// 	if err != nil {
// 		return output, err
// 	}

// 	output = CreateSingleConfigStruct(single)
// 	return output, err
// }

// func UpdateConfig(conn *connections.Connections, req datastruct.ConfigRequest) (datastruct.ConfigDataStruct, error) {
// 	var output datastruct.ConfigDataStruct
// 	var err error

// 	err = models.UpdateConfig(conn, req)
// 	if err != nil {
// 		return output, err
// 	}

// 	// jika tidak ada error, return single instance of single config
// 	single, err := models.GetSingleConfig(req.ConfigID, conn, req)
// 	if err != nil {
// 		return output, err
// 	}

// 	output = CreateSingleConfigStruct(single)
// 	return output, err
// }

// func DeleteConfig(conn *connections.Connections, req datastruct.ConfigRequest) error {
// 	err := models.DeleteConfig(conn, req)
// 	return err
// }
