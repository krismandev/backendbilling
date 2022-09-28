package models

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/lib"
	"backendbillingdashboard/modules/payment/datastruct"
	"errors"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func GetPaymentFromRequest(conn *connections.Connections, req datastruct.PaymentRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	var err error

	// -- THIS IS BASIC GET REQUEST EXAMPLE LOGIC
	var baseWhere string
	var baseParam []interface{}

	lib.AppendWhere(&baseWhere, &baseParam, "payment_id = ?", req.PaymentID)
	lib.AppendWhere(&baseWhere, &baseParam, "payment.invoice_id = ?", req.InvoiceID)
	lib.AppendWhere(&baseWhere, &baseParam, "account.account_id = ?", req.AccountID)
	lib.AppendWhere(&baseWhere, &baseParam, "payment.payment_type = ?", req.PaymentType)
	lib.AppendWhere(&baseWhere, &baseParam, "payment.status = ?", req.Status)
	lib.AppendWhere(&baseWhere, &baseParam, "payment.web_user_id = ?", req.WebUserID)
	lib.AppendWhere(&baseWhere, &baseParam, "invoice.invoice_no = ?", req.InvoiceNo)
	if len(req.ListPaymentID) > 0 {
		var baseIn string
		for _, prid := range req.ListPaymentID {
			lib.AppendComma(&baseIn, &baseParam, "?", prid)
		}
		lib.AppendWhereRaw(&baseWhere, "payment_id IN ("+baseIn+")")
	}

	var runQuery string
	if len(req.AccountID) > 0 {
		runQuery = `SELECT payment_id, payment.invoice_id, payment.payment_date, payment.total, payment.note, payment.last_update_username,  
		payment.payment_method as payment_method, payment.card_number, payment.payment_type, payment.payment_method, payment.clearing_date,
		payment.card_number, payment.status, (select count(payment_deduction.payment_id) from payment_deduction where payment_id = payment.payment_id) as payment_deduction_counter,
		payment.note, 
		invoice.account_id, invoice.invoice_no, account.name as account_name,
		pm.key, pm.payment_method_name, pm.payment_type as pm_payment_type, pm.currency_code, pm.bank_name  
		FROM payment JOIN invoice ON invoice.invoice_id = payment.invoice_id JOIN account ON account.account_id = invoice.account_id 
		LEFT JOIN payment_method as pm ON payment.payment_method = pm.key `
	} else {
		runQuery = `SELECT payment_id, payment.invoice_id, payment.payment_date, payment.total, payment.note, payment.last_update_username,  
		payment.payment_method as payment_method, payment.card_number, payment.payment_type,  payment.payment_method, payment.clearing_date, 
		payment.card_number, payment.status, (select count(payment_deduction.payment_id) from payment_deduction where payment_id = payment.payment_id) as payment_deduction_counter,
		payment.note,
		invoice.account_id, invoice.invoice_no, account.name as account_name,
		pm.key, pm.payment_method_name, pm.payment_type as pm_payment_type, pm.currency_code, pm.bank_name 
		FROM payment JOIN invoice ON invoice.invoice_id = payment.invoice_id JOIN account ON account.account_id = invoice.account_id 
		LEFT JOIN payment_method as pm ON payment.payment_method = pm.key `
	}

	if len(baseWhere) > 0 {
		runQuery += "WHERE " + baseWhere
	}

	lib.AppendOrderBy(&runQuery, req.Param.OrderBy, req.Param.OrderDir)
	lib.AppendLimit(&runQuery, req.Param.Page, req.Param.PerPage)

	resultSelect, _, err := conn.DBAppConn.SelectQueryByFieldNameSlice(runQuery, baseParam...)
	log.Info("LihatRes-", resultSelect)
	for _, each := range resultSelect {
		single := make(map[string]interface{})
		single["payment_id"] = each["payment_id"]
		single["invoice_id"] = each["invoice_id"]
		single["payment_date"] = each["payment_date"]
		single["total"] = each["total"]
		single["note"] = each["note"]
		single["last_update_username"] = each["last_update_username"]
		single["payment_method"] = each["payment_method"]
		single["card_number"] = each["card_number"]
		single["clearing_date"] = each["clearing_date"]
		single["status"] = each["status"]
		single["payment_type"] = each["payment_type"]
		var paymentDeductions []map[string]string

		paymentDeductionCounterInt, errconv := strconv.Atoi(each["payment_deduction_counter"])
		if errconv != nil {
			return result, errconv
		}
		log.Info("pydcounter-", paymentDeductionCounterInt)
		if paymentDeductionCounterInt > 0 {
			var resultPaymentDeductions []map[string]string
			qryGetPaymentDeductions := "SELECT payment_deduction_type_id, payment_id, amount, description FROM payment_deduction WHERE payment_id = ?"
			resultPaymentDeductions, _, err = conn.DBAppConn.SelectQueryByFieldNameSlice(qryGetPaymentDeductions, each["payment_id"])

			paymentDeductions = resultPaymentDeductions

		}

		single["payment_deductions"] = paymentDeductions

		invoice := make(map[string]interface{})
		invoice["invoice_no"] = each["invoice_no"]

		account := make(map[string]interface{})
		account["name"] = each["account_name"]
		invoice["account"] = account
		single["invoice"] = invoice

		paymentMethod := make(map[string]interface{})
		paymentMethod["key"] = each["key"]
		paymentMethod["payment_method_name"] = each["payment_method_name"]
		paymentMethod["payment_type"] = each["pm_payment_type"]
		paymentMethod["bank_name"] = each["bank_name"]
		paymentMethod["currency_code"] = each["currency_code"]

		single["payment_method_object"] = paymentMethod

		result = append(result, single)
	}

	log.Info("LihatPayment-", result)
	return result, err
}

func InsertPayment(conn *connections.Connections, req datastruct.PaymentRequest) error {
	var err error

	// -- THIS IS BASIC INSERT EXAMPLE
	var baseIn string
	var baseParam []interface{}

	// var baseInPaymentDeduction string
	// var baseParamPaymentDeduction []interface{}

	lastId, _ := conn.DBAppConn.GetFirstData("SELECT last_id FROM control_id where control_id.key=?", "payment")

	intLastId, err := strconv.Atoi(lastId)
	insertId := intLastId + 1

	insertIdString := strconv.Itoa(insertId)

	lib.AppendComma(&baseIn, &baseParam, "?", insertIdString)
	lib.AppendComma(&baseIn, &baseParam, "?", req.InvoiceID)
	lib.AppendComma(&baseIn, &baseParam, "?", req.Total)
	lib.AppendComma(&baseIn, &baseParam, "?", req.Note)
	lib.AppendComma(&baseIn, &baseParam, "?", req.LastUpdateUsername)
	lib.AppendComma(&baseIn, &baseParam, "?", req.PaymentDate)
	lib.AppendComma(&baseIn, &baseParam, "?", req.PaymentType)
	lib.AppendComma(&baseIn, &baseParam, "?", req.ClearingDate)
	lib.AppendComma(&baseIn, &baseParam, "?", req.CardNumber)
	lib.AppendComma(&baseIn, &baseParam, "?", req.PaymentMethod)
	lib.AppendComma(&baseIn, &baseParam, "?", req.WebUserID)
	lib.AppendCommaRaw(&baseIn, "now()")

	qry := "INSERT INTO payment (payment_id, invoice_id, total, note, last_update_username, payment_date, payment_type, clearing_date, card_number, payment_method, web_user_id, created_at) VALUES (" + baseIn + ")"
	_, _, errInsert := conn.DBAppConn.Exec(qry, baseParam...)
	if errInsert != nil {
		return errInsert
	}

	err = UpdateControlId(conn, insertIdString, "payment")

	if len(req.PaymentDeduction) > 0 {
		bulkInsertQuery := "INSERT INTO payment_deduction (payment_deduction_type_id, invoice_id, payment_id, amount, description, last_update_username) VALUES "
		var paramsBulkInsert []interface{}
		var stringGroup []string
		for _, each := range req.PaymentDeduction {
			partquery := "(?,?,?, ?, ?, ?)"
			paramsBulkInsert = append(paramsBulkInsert, each.PaymentDeductionTypeID)
			paramsBulkInsert = append(paramsBulkInsert, req.InvoiceID)
			paramsBulkInsert = append(paramsBulkInsert, insertIdString)
			paramsBulkInsert = append(paramsBulkInsert, each.Amount)
			paramsBulkInsert = append(paramsBulkInsert, each.Description)
			paramsBulkInsert = append(paramsBulkInsert, req.LastUpdateUsername)
			// paramsBulkInsert = append(paramsBulkInsert, "0")
			// paramsBulkInsert = append(paramsBulkInsert, req.LastUpdateUsername)
			stringGroup = append(stringGroup, partquery)
		}

		final_query_bulk := bulkInsertQuery + strings.Join(stringGroup, ", ")
		_, _, errInsertBulk := conn.DBAppConn.Exec(final_query_bulk, paramsBulkInsert...)
		if errInsertBulk != nil {
			return errInsertBulk
		}
	}

	qryGetSudahDibayar := "SELECT IFNULL(SUM(payment.total),0) FROM payment where invoice_id=? and status = 1"
	sudahDibayar, _ := conn.DBAppConn.GetFirstData(qryGetSudahDibayar, req.InvoiceID)

	// subTotalFloat, err := strconv.ParseFloat(subTotal, 64)
	sudahDibayarFloat, err := strconv.ParseFloat(sudahDibayar, 64)

	qryGetInvoiceData := "SELECT grand_total FROM invoice WHERE invoice.invoice_id = ?"
	resInvoice, _, errGetInvoiceData := conn.DBAppConn.SelectQueryByFieldNameSlice(qryGetInvoiceData, req.InvoiceID)
	if errGetInvoiceData != nil {
		return errGetInvoiceData
	}

	oldInvoice := resInvoice[0]

	grandTotal := oldInvoice["grand_total"]
	grandTotalFloat, err := strconv.ParseFloat(grandTotal, 64)

	qryUpdateInvoicePaymentStatus := "UPDATE invoice SET invoice.paid = ? WHERE invoice.invoice_id = ?"
	if sudahDibayarFloat == grandTotalFloat {
		_, _, errUpdateInvoice := conn.DBAppConn.Exec(qryUpdateInvoicePaymentStatus, "1", req.InvoiceID)
		if errUpdateInvoice != nil {
			return errUpdateInvoice
		}
	} else if sudahDibayarFloat > grandTotalFloat {
		_, _, errUpdateInvoice := conn.DBAppConn.Exec(qryUpdateInvoicePaymentStatus, "2", req.InvoiceID)
		if errUpdateInvoice != nil {
			return errUpdateInvoice
		}
	}

	return err
}

// error if invoice is unpaid
func CheckIsInvoiceAlreadyPaid(conn *connections.Connections, req datastruct.PaymentRequest, invoiceID string) error {
	var err error

	qryGetSudahDibayar := "SELECT IFNULL(SUM(payment.total),0) FROM payment where invoice_id=? and status = 1"
	sudahDibayar, _ := conn.DBAppConn.GetFirstData(qryGetSudahDibayar, invoiceID)

	// subTotalFloat, err := strconv.ParseFloat(subTotal, 64)
	sudahDibayarFloat, err := strconv.ParseFloat(sudahDibayar, 64)

	// qryGetSudahDibayar := "SELECT SUM(total)"
	qryGetInvoiceData := "SELECT grand_total FROM invoice WHERE invoice.invoice_id = ?"
	resInvoice, _, errGetInvoiceData := conn.DBAppConn.SelectQueryByFieldNameSlice(qryGetInvoiceData, invoiceID)
	if errGetInvoiceData != nil {
		return errGetInvoiceData
	}

	oldInvoice := resInvoice[0]
	grandTotal := oldInvoice["grand_total"]
	grandTotalFloat, err := strconv.ParseFloat(grandTotal, 64)

	// sisa := grandTotalFloat - sudahDibayarFloat
	// totalAkanDibayarFloat, err := strconv.ParseFloat(req.Total, 64)
	if sudahDibayarFloat < grandTotalFloat {
		return errors.New("Invoice is unpaid")
	}
	return err
}

func UpdatePayment(conn *connections.Connections, req datastruct.PaymentRequest) error {
	var err error

	// -- THIS IS BASIC UPDATE EXAMPLE
	// var baseUp string
	// var baseParam []interface{}

	// lib.AppendComma(&baseUp, &baseParam, "paymentname = ?", req.PaymentName)
	// qry := "UPDATE payment SET " + baseUp + " WHERE paymentid = ?"
	// baseParam = append(baseParam, req.PaymentID)
	// _, _, err = conn.DBAppConn.Exec(qry, baseParam...)
	return err
}

func DeletePayment(conn *connections.Connections, req datastruct.PaymentRequest) error {
	var err error
	// -- THIS IS DELETE LOGIC EXAMPLE
	qry := "UPDATE payment SET payment.status = 2 WHERE payment_id = ?"
	_, _, err = conn.DBAppConn.Exec(qry, req.PaymentID)

	qryPaymentDeduction := "UPDATE payment_deduction SET status = 'D' WHERE payment_id = ?"
	_, _, err = conn.DBAppConn.Exec(qryPaymentDeduction, req.PaymentID)

	qryGetInvoiceID := "SELECT invoice_id FROM payment where payment_id = ?"
	invoiceID, _ := conn.DBAppConn.GetFirstData(qryGetInvoiceID, req.PaymentID)

	isInvoiceUnpaid := CheckIsInvoiceAlreadyPaid(conn, req, invoiceID)
	if isInvoiceUnpaid != nil {
		log.Error(isInvoiceUnpaid)
		_, _, err = conn.DBAppConn.Exec("UPDATE invoice SET paid = 0 WHERE invoice_id = ?", invoiceID)
	}

	return err
}

func GetPaymentDeductionTypeFromRequest(conn *connections.Connections, req datastruct.PaymentDeductionTypeRequest) ([]map[string]string, error) {
	var result []map[string]string
	var err error

	// -- THIS IS BASIC GET REQUEST EXAMPLE LOGIC
	var baseWhere string
	var baseParam []interface{}

	lib.AppendWhere(&baseWhere, &baseParam, "payment_deduction_type_id = ?", req.PaymentDeductionTypeID)
	lib.AppendWhere(&baseWhere, &baseParam, "description = ?", req.Description)
	lib.AppendWhere(&baseWhere, &baseParam, "category = ?", req.Category)
	if len(req.ListPaymentDeductionTypeID) > 0 {
		var baseIn string
		for _, prid := range req.ListPaymentDeductionTypeID {
			lib.AppendComma(&baseIn, &baseParam, "?", prid)
		}
		lib.AppendWhereRaw(&baseWhere, "payment_deduction_type_id IN ("+baseIn+")")
	}

	runQuery := "SELECT payment_deduction_type_id, description, category, amount, last_update_username, last_update_date FROM payment_deduction_type "
	if len(baseWhere) > 0 {
		runQuery += "WHERE " + baseWhere
	}
	lib.AppendOrderBy(&runQuery, req.Param.OrderBy, req.Param.OrderDir)
	lib.AppendLimit(&runQuery, req.Param.Page, req.Param.PerPage)

	result, _, err = conn.DBAppConn.SelectQueryByFieldNameSlice(runQuery, baseParam...)
	return result, err
}

func GetPaymentDeductionFromRequest(conn *connections.Connections, req datastruct.PaymentDeductionRequest) ([]map[string]string, error) {
	var result []map[string]string
	var err error

	// -- THIS IS BASIC GET REQUEST EXAMPLE LOGIC
	var baseWhere string
	var baseParam []interface{}

	lib.AppendWhere(&baseWhere, &baseParam, "payment_deduction.payment_deduction_type_id = ?", req.PaymentDeductionTypeID)
	lib.AppendWhere(&baseWhere, &baseParam, "payment_id = ?", req.PaymentID)
	lib.AppendWhere(&baseWhere, &baseParam, "invoice_id = ?", req.InvoiceID)
	lib.AppendWhere(&baseWhere, &baseParam, "payment_deduction.status = ?", req.Status)
	// if len(req.ListAccountID) > 0 {
	// 	var baseIn string
	// 	for _, prid := range req.ListAccountID {
	// 		lib.AppendComma(&baseIn, &baseParam, "?", prid)
	// 	}
	// 	lib.AppendWhereRaw(&baseWhere, "account_id IN ("+baseIn+")")
	// }

	runQuery := `SELECT payment_deduction_type.payment_deduction_type_id, payment_id, invoice_id, payment_deduction.amount, payment_deduction.description, payment_deduction_type.unique_in_invoice, 
	payment_deduction.last_update_username, payment_deduction.last_update_date, payment_deduction.status
	FROM payment_deduction 
	JOIN payment_deduction_type ON payment_deduction_type.payment_deduction_type_id = payment_deduction.payment_deduction_type_id `
	if len(baseWhere) > 0 {
		runQuery += "WHERE " + baseWhere
	}
	lib.AppendOrderBy(&runQuery, req.Param.OrderBy, req.Param.OrderDir)
	lib.AppendLimit(&runQuery, req.Param.Page, req.Param.PerPage)

	result, _, err = conn.DBAppConn.SelectQueryByFieldNameSlice(runQuery, baseParam...)
	return result, err
}

func GetAdjustmentReasonFromRequest(conn *connections.Connections, req datastruct.AdjustmentReasonRequest) ([]map[string]string, error) {
	var result []map[string]string
	var err error

	// -- THIS IS BASIC GET REQUEST EXAMPLE LOGIC
	var baseWhere string
	var baseParam []interface{}

	lib.AppendWhere(&baseWhere, &baseParam, "adjustment_reason.key = ?", req.Key)
	lib.AppendWhere(&baseWhere, &baseParam, "payment_method = ?", req.PaymentMethod)
	lib.AppendWhere(&baseWhere, &baseParam, "currency_code = ?", req.CurrencyCode)
	// if len(req.ListAccountID) > 0 {
	// 	var baseIn string
	// 	for _, prid := range req.ListAccountID {
	// 		lib.AppendComma(&baseIn, &baseParam, "?", prid)
	// 	}
	// 	lib.AppendWhereRaw(&baseWhere, "account_id IN ("+baseIn+")")
	// }

	runQuery := `SELECT adjustment_reason.key, description, payment_method, currency_code 
	FROM adjustment_reason `
	if len(baseWhere) > 0 {
		runQuery += "WHERE " + baseWhere
	}
	lib.AppendOrderBy(&runQuery, req.Param.OrderBy, req.Param.OrderDir)
	lib.AppendLimit(&runQuery, req.Param.Page, req.Param.PerPage)

	result, _, err = conn.DBAppConn.SelectQueryByFieldNameSlice(runQuery, baseParam...)
	return result, err
}
