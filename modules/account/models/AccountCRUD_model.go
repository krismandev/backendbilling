package models

import (
	"billingdashboard/connections"
	"billingdashboard/lib"
	"billingdashboard/modules/account/datastruct"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func GetAccountFromRequest(conn *connections.Connections, req datastruct.AccountRequest) ([]map[string]string, error) {
	var result []map[string]string
	var err error

	// -- THIS IS BASIC GET REQUEST EXAMPLE LOGIC
	var baseWhere string
	var baseParam []interface{}

	lib.AppendWhere(&baseWhere, &baseParam, "account_id = ?", req.AccountID)
	lib.AppendWhere(&baseWhere, &baseParam, "company_id = ?", req.CompanyID)
	lib.AppendWhere(&baseWhere, &baseParam, "name = ?", req.Name)
	if len(req.ListAccountID) > 0 {
		var baseIn string
		for _, prid := range req.ListAccountID {
			lib.AppendComma(&baseIn, &baseParam, "?", prid)
		}
		lib.AppendWhereRaw(&baseWhere, "account_id IN ("+baseIn+")")
	}

	runQuery := "SELECT account_id, name, status, company_id, address1, address2, account_type, billing_type,city, phone, contact_person, contact_person_phone ,account.desc, last_update_username, last_update_date FROM account "
	if len(baseWhere) > 0 {
		runQuery += "WHERE " + baseWhere
	}
	lib.AppendOrderBy(&runQuery, req.Param.OrderBy, req.Param.OrderDir)
	if req.Param.NoPagination == true {
		//DO Nothing
	} else {
		lib.AppendLimit(&runQuery, req.Param.Page, req.Param.PerPage)
	}

	result, _, err = conn.DBAppConn.SelectQueryByFieldNameSlice(runQuery, baseParam...)
	return result, err
}

func InsertAccount(conn *connections.Connections, req datastruct.AccountRequest) error {
	var err error

	// -- THIS IS BASIC INSERT EXAMPLE
	var baseIn string
	var baseParam []interface{}

	lastId, _ := conn.DBAppConn.GetFirstData("SELECT last_id FROM control_id where control_id.key=?", "account")
	// lastId, _ := conn.DBAppConn.GetFirstData("SELECT last_id FROM control_id where control_id.key=?", "account")

	intLastId, err := strconv.Atoi(lastId)
	insertId := intLastId + 1

	insertIdString := strconv.Itoa(insertId)

	lib.AppendComma(&baseIn, &baseParam, "?", insertIdString)
	lib.AppendComma(&baseIn, &baseParam, "?", req.Name)
	lib.AppendComma(&baseIn, &baseParam, "?", req.Status)
	lib.AppendComma(&baseIn, &baseParam, "?", req.CompanyID)
	lib.AppendComma(&baseIn, &baseParam, "?", req.AccountType)
	lib.AppendComma(&baseIn, &baseParam, "?", req.BillingType)
	lib.AppendComma(&baseIn, &baseParam, "?", req.Desc)
	lib.AppendComma(&baseIn, &baseParam, "?", req.Address1)
	lib.AppendComma(&baseIn, &baseParam, "?", req.Address2)
	lib.AppendComma(&baseIn, &baseParam, "?", req.City)
	lib.AppendComma(&baseIn, &baseParam, "?", req.Phone)
	lib.AppendComma(&baseIn, &baseParam, "?", req.ContactPerson)
	lib.AppendComma(&baseIn, &baseParam, "?", req.ContactPersonPhone)

	lib.AppendComma(&baseIn, &baseParam, "?", req.LastUpdateUsername)

	qry := "INSERT INTO account (account_id,  name, status, company_id, account_type, billing_type, account.desc, address1,address2,city,phone,contact_person,contact_person_phone, last_update_username) VALUES (" + baseIn + ")"
	_, err = conn.DBAppConn.InsertGetLastID(qry, baseParam...)

	log.Info("InsertParam - ", baseParam)
	_, _, err = conn.DBAppConn.Exec("UPDATE control_id set last_id=? where control_id.key=?", insertIdString, "account")

	return err
}

func UpdateAccount(conn *connections.Connections, req datastruct.AccountRequest) error {
	var err error

	var baseUp string
	var baseParam []interface{}

	lib.AppendComma(&baseUp, &baseParam, "name = ?", req.Name)
	lib.AppendComma(&baseUp, &baseParam, "status = ?", req.Status)
	lib.AppendComma(&baseUp, &baseParam, "account.desc = ?", req.Desc)
	lib.AppendComma(&baseUp, &baseParam, "account_type = ?", req.AccountType)
	lib.AppendComma(&baseUp, &baseParam, "billing_type = ?", req.BillingType)
	lib.AppendComma(&baseUp, &baseParam, "account.desc = ?", req.Desc)
	lib.AppendComma(&baseUp, &baseParam, "address1 = ?", req.Address1)
	lib.AppendComma(&baseUp, &baseParam, "address2 = ?", req.Address2)
	lib.AppendComma(&baseUp, &baseParam, "city = ?", req.City)
	lib.AppendComma(&baseUp, &baseParam, "phone = ? ", req.Phone)
	lib.AppendComma(&baseUp, &baseParam, "contact_person = ?", req.ContactPerson)
	lib.AppendComma(&baseUp, &baseParam, "contact_person_phone = ?", req.ContactPersonPhone)
	lib.AppendComma(&baseUp, &baseParam, "last_update_username = ?", req.LastUpdateUsername)

	qry := "UPDATE account SET " + baseUp + " WHERE account_id = ?"
	baseParam = append(baseParam, req.AccountID)
	_, _, err = conn.DBAppConn.Exec(qry, baseParam...)
	return err
}

func DeleteAccount(conn *connections.Connections, req datastruct.AccountRequest) error {
	var err error
	// -- THIS IS DELETE LOGIC EXAMPLE
	qry := "DELETE FROM account WHERE account_id = ?"
	_, _, err = conn.DBAppConn.Exec(qry, req.AccountID)
	return err
}
