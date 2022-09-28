package processors

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/stub/datastruct"
	"backendbillingdashboard/modules/stub/models"
)

func GetListStub(conn *connections.Connections, req datastruct.StubRequest) ([]datastruct.StubDataStruct, error) {
	var output []datastruct.StubDataStruct
	var err error

	// grab mapping data from model
	stubList, err := models.GetStubFromRequest(conn, req)
	if err != nil {
		return output, err
	}

	for _, stub := range stubList {
		single := CreateSingleStubStruct(stub)
		output = append(output, single)
	}

	return output, err
}

func CreateSingleStubStruct(stub map[string]string) datastruct.StubDataStruct {
	var single datastruct.StubDataStruct
	single.StubID, _ = stub["stubid"]
	single.StubName, _ = stub["stubname"]

	return single
}

func InsertStub(conn *connections.Connections, req datastruct.StubRequest) (datastruct.StubDataStruct, error) {
	var output datastruct.StubDataStruct
	var err error

	err = models.InsertStub(conn, req)
	if err != nil {
		return output, err
	}

	// jika tidak ada error, return single instance of single stub
	single, err := models.GetSingleStub(req.StubID, conn, req)
	if err != nil {
		return output, err
	}

	output = CreateSingleStubStruct(single)
	return output, err
}

func UpdateStub(conn *connections.Connections, req datastruct.StubRequest) (datastruct.StubDataStruct, error) {
	var output datastruct.StubDataStruct
	var err error

	err = models.UpdateStub(conn, req)
	if err != nil {
		return output, err
	}

	// jika tidak ada error, return single instance of single stub
	single, err := models.GetSingleStub(req.StubID, conn, req)
	if err != nil {
		return output, err
	}

	output = CreateSingleStubStruct(single)
	return output, err
}

func DeleteStub(conn *connections.Connections, req datastruct.StubRequest) error {
	err := models.DeleteStub(conn, req)
	return err
}
