package models

import (
	"billingdashboard/connections"
	"billingdashboard/lib"
	"billingdashboard/modules/company/datastruct"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func GetCompanyFromRequest(conn *connections.Connections, req datastruct.CompanyRequest) ([]map[string]string, error) {
	var result []map[string]string
	var err error

	// -- THIS IS BASIC GET REQUEST EXAMPLE LOGIC
	var baseWhere string
	var baseParam []interface{}

	lib.AppendWhere(&baseWhere, &baseParam, "companyid = ?", req.CompanyID)
	lib.AppendWhere(&baseWhere, &baseParam, "companyid = ?", req.Name)
	lib.AppendWhere(&baseWhere, &baseParam, "status = ?", req.Status)
	lib.AppendWhere(&baseWhere, &baseParam, "city = ?", req.City)
	if len(req.ListCompanyID) > 0 {
		var baseIn string
		for _, prid := range req.ListCompanyID {
			lib.AppendComma(&baseIn, &baseParam, "?", prid)
		}
		lib.AppendWhereRaw(&baseWhere, "companyid IN ("+baseIn+")")
	}

	runQuery := "SELECT companyid, name, status, addr1, addr2, city, country, contactperson, contactpersonphone, phone, fax, company.desc, last_update_username, last_update_date FROM company "
	if len(baseWhere) > 0 {
		runQuery += "WHERE " + baseWhere
	}
	lib.AppendOrderBy(&runQuery, req.Param.OrderBy, req.Param.OrderDir)
	lib.AppendLimit(&runQuery, req.Param.Page, req.Param.PerPage)

	result, _, err = conn.DBAppConn.SelectQueryByFieldNameSlice(runQuery, baseParam...)
	return result, err
}

func InsertCompany(conn *connections.Connections, req datastruct.CompanyRequest) error {
	var err error

	// -- THIS IS BASIC INSERT EXAMPLE
	var baseIn string
	var baseParam []interface{}

	lastId, _ := conn.DBAppConn.GetFirstData("SELECT max(companyid) FROM company ")

	intLastId, err := strconv.Atoi(lastId)
	insertId := intLastId + 1

	log.Info("HasilID", lastId)

	insertIdString := strconv.Itoa(insertId)
	lib.AppendComma(&baseIn, &baseParam, "?", insertIdString)
	lib.AppendComma(&baseIn, &baseParam, "?", req.Name)
	lib.AppendComma(&baseIn, &baseParam, "?", req.Status)
	lib.AppendComma(&baseIn, &baseParam, "?", req.Addr1)
	lib.AppendComma(&baseIn, &baseParam, "?", req.Addr2)
	lib.AppendComma(&baseIn, &baseParam, "?", req.City)
	lib.AppendComma(&baseIn, &baseParam, "?", req.Country)
	lib.AppendComma(&baseIn, &baseParam, "?", req.ContactPerson)
	lib.AppendComma(&baseIn, &baseParam, "?", req.ContactPersonPhone)
	lib.AppendComma(&baseIn, &baseParam, "?", req.Phone)
	lib.AppendComma(&baseIn, &baseParam, "?", req.Fax)
	lib.AppendComma(&baseIn, &baseParam, "?", req.Desc)
	lib.AppendComma(&baseIn, &baseParam, "?", req.LastUpdateUsername)

	qry := "INSERT INTO company (companyid, name, status, addr1, addr2, city, country, contactperson, contactpersonphone, phone, fax, company.desc, last_update_username) VALUES (" + baseIn + ")"
	newId, err := conn.DBAppConn.InsertGetLastID(qry, baseParam...)
	log.Info("NewID - ", newId)
	_, _, err = conn.DBAppConn.Exec("UPDATE control_id set last_id=? where control_id.key=?", insertIdString, "company")
	return err
}

func UpdateCompany(conn *connections.Connections, req datastruct.CompanyRequest) error {
	var err error

	var baseUp string
	var baseParam []interface{}

	lib.AppendComma(&baseUp, &baseParam, "name = ?", req.Name)
	lib.AppendComma(&baseUp, &baseParam, "status = ?", req.Status)
	lib.AppendComma(&baseUp, &baseParam, "addr1 = ?", req.Addr1)
	lib.AppendComma(&baseUp, &baseParam, "addr2 = ?", req.Addr2)
	lib.AppendComma(&baseUp, &baseParam, "city = ?", req.City)
	lib.AppendComma(&baseUp, &baseParam, "country = ?", req.Country)
	lib.AppendComma(&baseUp, &baseParam, "contactperson = ?", req.ContactPerson)
	lib.AppendComma(&baseUp, &baseParam, "contactpersonphone = ?", req.ContactPersonPhone)
	lib.AppendComma(&baseUp, &baseParam, "phone = ?", req.Phone)
	lib.AppendComma(&baseUp, &baseParam, "fax = ?", req.Fax)
	lib.AppendComma(&baseUp, &baseParam, "company.desc = ?", req.Desc)
	lib.AppendComma(&baseUp, &baseParam, "last_update_username = ?", req.LastUpdateUsername)
	qry := "UPDATE company SET " + baseUp + " WHERE companyid = ?"
	baseParam = append(baseParam, req.CompanyID)
	_, _, err = conn.DBAppConn.Exec(qry, baseParam...)
	return err
}

func DeleteCompany(conn *connections.Connections, req datastruct.CompanyRequest) error {
	var err error
	// -- THIS IS DELETE LOGIC EXAMPLE
	qry := "DELETE FROM company WHERE companyid = ?"
	_, _, err = conn.DBAppConn.Exec(qry, req.CompanyID)
	return err
}
