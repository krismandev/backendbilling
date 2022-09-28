package models

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/company/datastruct"
	"errors"
	"strconv"
)

func GetSingleCompany(companyID string, conn *connections.Connections, req datastruct.CompanyRequest) (map[string]string, error) {
	var result map[string]string
	var err error

	// -- EXAMPLE
	if len(companyID) == 0 {
		companyID = req.CompanyID
	}
	query := "SELECT companyid, name, status, addr1, addr2, city, country, contactperson, contactpersonphone, phone, fax, company.desc, last_update_username, last_update_date FROM company WHERE companyid = ?"
	results, _, err := conn.DBAppConn.SelectQueryByFieldNameSlice(query, companyID)
	if err != nil {
		return result, err
	}

	// convert from []map[string]string to single map[string]string
	for _, res := range results {
		result = res
		break
	}
	return result, err
}

func CheckCompanyExists(stubID string, conn *connections.Connections) error {
	// var param []interface{}
	// qry := "SELECT COUNT(stubid) FROM stub WHERE stubid = ?"
	// param = append(param, stubID)

	// cnt, _ := conn.DBAppConn.GetFirstData(qry, param...)
	// datacount, _ := strconv.Atoi(cnt)
	// if datacount == 0 {
	// 	return errors.New("Company ID is not exists")
	// }
	return nil
}

func CheckCompanyDuplicate(exceptID string, conn *connections.Connections, req datastruct.CompanyRequest) error {
	var param []interface{}
	qry := "SELECT COUNT(company_id) FROM company WHERE company_id = ?"
	param = append(param, req.CompanyID)
	if len(exceptID) > 0 {
		qry += " AND company_id <> ?"
		param = append(param, exceptID)
	}

	cnt, _ := conn.DBAppConn.GetFirstData(qry, param...)
	datacount, _ := strconv.Atoi(cnt)
	if datacount > 0 {
		return errors.New("Company ID is already exists. Please use another Company ID")
	}
	return nil
}
