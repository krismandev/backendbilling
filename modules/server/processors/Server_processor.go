package processors

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/server/datastruct"
	"backendbillingdashboard/modules/server/models"
)

func GetListServer(conn *connections.Connections, req datastruct.ServerRequest) ([]datastruct.ServerDataStruct, error) {
	var output []datastruct.ServerDataStruct
	var err error

	// grab mapping data from model
	serverList, err := models.GetServerFromRequest(conn, req)
	if err != nil {
		return output, err
	}

	for _, server := range serverList {
		single := CreateSingleServerStruct(server)
		output = append(output, single)
	}

	return output, err
}

func CreateSingleServerStruct(server map[string]string) datastruct.ServerDataStruct {
	var single datastruct.ServerDataStruct
	single.ServerID, _ = server["server_id"]
	single.ServerName, _ = server["server_name"]
	single.ServerUrl, _ = server["server_url"]

	return single
}

func InsertServer(conn *connections.Connections, req datastruct.ServerRequest) error {
	var err error

	err = models.InsertServer(conn, req)
	if err != nil {
		return err
	}

	return err
}

func UpdateServer(conn *connections.Connections, req datastruct.ServerRequest) error {
	var err error

	err = models.UpdateServer(conn, req)
	if err != nil {
		return err
	}

	// jika tidak ada error, return single instance of single server

	return err
}

func DeleteServer(conn *connections.Connections, req datastruct.ServerRequest) error {
	err := models.DeleteServer(conn, req)
	return err
}

func GetServerAccount(conn *connections.Connections, req datastruct.ServerAccountRequest) ([]datastruct.ServerAccountDataStruct, error) {
	var output []datastruct.ServerAccountDataStruct
	var err error

	// grab mapping data from model
	serverList, err := models.GetServerAccountFromRequest(conn, req)
	if err != nil {
		return output, err
	}

	for _, server := range serverList {
		single := CreateSingleServerAccountStruct(server)
		output = append(output, single)
	}

	return output, err
}

func CreateSingleServerAccountStruct(server map[string]string) datastruct.ServerAccountDataStruct {
	var single datastruct.ServerAccountDataStruct
	single.ServerID, _ = server["server_id"]
	single.AccountID, _ = server["account_id"]
	single.ExternalAccountID, _ = server["external_account_id"]
	single.LastUpdateUsername, _ = server["last_update_username"]
	single.LastUpdateDate, _ = server["last_update_date"]

	return single
}
