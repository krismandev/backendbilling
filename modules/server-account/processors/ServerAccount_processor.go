package processors

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/server-account/datastruct"
	"backendbillingdashboard/modules/server-account/models"
)

func GetListServerAccount(conn *connections.Connections, req datastruct.ServerAccountRequest) ([]datastruct.ServerAccountDataStruct, error) {
	var output []datastruct.ServerAccountDataStruct
	var err error

	// grab mapping data from model
	ServerAccountList, err := models.GetServerAccountFromRequest(conn, req)
	if err != nil {
		return output, err
	}

	for _, ServerAccount := range ServerAccountList {
		single := CreateSingleServerAccountStruct(ServerAccount)
		output = append(output, single)
	}

	return output, err
}

func CreateSingleServerAccountStruct(ServerAccount map[string]string) datastruct.ServerAccountDataStruct {
	var single datastruct.ServerAccountDataStruct
	single.AccountID, _ = ServerAccount["account_id"]
	single.ServerID, _ = ServerAccount["server_id"]
	single.ExternalAccountID, _ = ServerAccount["external_account_id"]

	return single
}

func InsertServerAccount(conn *connections.Connections, req datastruct.ServerAccountRequest) (datastruct.ServerAccountDataStruct, error) {
	var output datastruct.ServerAccountDataStruct
	var err error

	err = models.InsertServerAccount(conn, req)
	if err != nil {
		return output, err
	}

	// jika tidak ada error, return single instance of single ServerAccount
	single, err := models.GetSingleServerAccount(req.AccountID, conn, req)
	if err != nil {
		return output, err
	}

	output = CreateSingleServerAccountStruct(single)
	return output, err
}

func UpdateServerAccount(conn *connections.Connections, req datastruct.ServerAccountRequest) (datastruct.ServerAccountDataStruct, error) {
	var output datastruct.ServerAccountDataStruct
	var err error

	err = models.UpdateServerAccount(conn, req)
	if err != nil {
		return output, err
	}

	// jika tidak ada error, return single instance of single ServerAccount
	single, err := models.GetSingleServerAccount(req.AccountID, conn, req)
	if err != nil {
		return output, err
	}

	output = CreateSingleServerAccountStruct(single)
	return output, err
}

func DeleteServerAccount(conn *connections.Connections, req datastruct.ServerAccountRequest) error {
	err := models.DeleteServerAccount(conn, req)
	return err
}
