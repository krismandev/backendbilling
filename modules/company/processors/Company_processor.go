package processors

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/company/datastruct"
	"backendbillingdashboard/modules/company/models"
)

func GetListCompany(conn *connections.Connections, req datastruct.CompanyRequest) ([]datastruct.CompanyDataStruct, error) {
	var output []datastruct.CompanyDataStruct
	var err error

	// grab mapping data from model
	companyList, err := models.GetCompanyFromRequest(conn, req)
	if err != nil {
		return output, err
	}

	for _, company := range companyList {
		single := CreateSingleCompanyStruct(company)
		output = append(output, single)
	}

	return output, err
}

func CreateSingleCompanyStruct(stub map[string]string) datastruct.CompanyDataStruct {
	var single datastruct.CompanyDataStruct
	single.CompanyID = stub["companyid"]
	single.Name = stub["name"]
	single.Status = stub["status"]
	single.Addr1 = stub["addr1"]
	single.Addr2 = stub["addr2"]
	single.City = stub["city"]
	single.Country = stub["country"]
	single.ContactPerson = stub["contactperson"]
	single.ContactPersonPhone = stub["contactpersonphone"]
	single.Phone = stub["phone"]
	single.Fax = stub["fax"]
	single.Desc = stub["desc"]
	single.LastUpdateUsername = stub["last_update_username"]
	single.LastUpdateDate = stub["last_update_date"]

	return single
}

func InsertCompany(conn *connections.Connections, req datastruct.CompanyRequest) (datastruct.CompanyDataStruct, error) {
	var output datastruct.CompanyDataStruct
	var err error

	err = models.InsertCompany(conn, req)
	if err != nil {
		return output, err
	}

	// jika tidak ada error, return single instance of single stub
	single, err := models.GetSingleCompany(req.CompanyID, conn, req)
	if err != nil {
		return output, err
	}

	output = CreateSingleCompanyStruct(single)
	return output, err
}

func UpdateCompany(conn *connections.Connections, req datastruct.CompanyRequest) (datastruct.CompanyDataStruct, error) {
	var output datastruct.CompanyDataStruct
	var err error

	err = models.UpdateCompany(conn, req)
	if err != nil {
		return output, err
	}

	// jika tidak ada error, return single instance of single stub
	single, err := models.GetSingleCompany(req.CompanyID, conn, req)
	if err != nil {
		return output, err
	}

	output = CreateSingleCompanyStruct(single)
	return output, err
}

func DeleteCompany(conn *connections.Connections, req datastruct.CompanyRequest) error {
	err := models.DeleteCompany(conn, req)
	return err
}
