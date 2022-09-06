package models

import (
	"billingdashboard/connections"
	"billingdashboard/lib"
	"billingdashboard/modules/payment/datastruct"
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
	if len(req.ListPaymentID) > 0 {
		var baseIn string
		for _, prid := range req.ListPaymentID {
			lib.AppendComma(&baseIn, &baseParam, "?", prid)
		}
		lib.AppendWhereRaw(&baseWhere, "payment_id IN ("+baseIn+")")
	}

	var runQuery string
	if len(req.AccountID) > 0 {
		runQuery = `SELECT payment_id, payment.invoice_id, payment.payment_date, payment.total, payment.note, payment.created_by, username,  
		payment.payment_method, payment.card_number, invoice.account_id, invoice.invoice_no, account.name as account_name, payment.payment_method, payment.clearing_date,
		payment.card_number, payment.status, (select count(payment_deduction.payment_id) from payment_deduction where payment_id = payment.payment_id) as payment_deduction_counter FROM payment JOIN invoice ON invoice.invoice_id = payment.invoice_id JOIN account ON account.account_id = invoice.account_id `
	} else {
		runQuery = `SELECT payment_id, payment.invoice_id, payment.payment_date, payment.total, payment.note, payment.created_by, username,  
		payment.payment_method, payment.card_number, invoice.account_id, invoice.invoice_no, account.name as account_name, payment.payment_method, payment.clearing_date, 
		payment.card_number, payment.status, (select count(payment_deduction.payment_id) from payment_deduction where payment_id = payment.payment_id) as payment_deduction_counter FROM payment JOIN invoice ON invoice.invoice_id = payment.invoice_id JOIN account ON account.account_id = invoice.account_id `
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
		single["created_by"] = each["created_by"]
		single["username"] = each["username"]
		single["payment_method"] = each["payment_method"]
		single["card_number"] = each["card_number"]
		single["clearing_date"] = each["clearing_date"]
		single["status"] = each["status"]
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
	// lastId, _ := conn.DBAppConn.GetFirstData("SELECT last_id FROM control_id where control_id.key=?", "account")

	// lastIdPaymentDeduction, _ := conn.DBAppConn.GetFirstData("SELECT last_id FROM control_id where control_id.key=?", "payment_deduction")

	intLastId, err := strconv.Atoi(lastId)
	insertId := intLastId + 1

	insertIdString := strconv.Itoa(insertId)

	// intLastIdPaymentDeduction, err := strconv.Atoi(lastIdPaymentDeduction)
	// insertIdPaymentDeduction := intLastIdPaymentDeduction + 1

	// insertIdStringPaymentDeduction := strconv.Itoa(insertIdPaymentDeduction)

	lib.AppendComma(&baseIn, &baseParam, "?", insertIdString)
	lib.AppendComma(&baseIn, &baseParam, "?", req.InvoiceID)
	lib.AppendComma(&baseIn, &baseParam, "?", req.Total)
	lib.AppendComma(&baseIn, &baseParam, "?", req.Note)
	lib.AppendComma(&baseIn, &baseParam, "?", req.LastUpdateUsername)
	lib.AppendComma(&baseIn, &baseParam, "?", req.LastUpdateUsername)
	lib.AppendComma(&baseIn, &baseParam, "?", req.PaymentDate)
	lib.AppendComma(&baseIn, &baseParam, "?", req.PaymentType)
	lib.AppendComma(&baseIn, &baseParam, "?", req.ClearingDate)
	lib.AppendComma(&baseIn, &baseParam, "?", req.CardNumber)
	lib.AppendComma(&baseIn, &baseParam, "?", req.PaymentMethod)
	lib.AppendCommaRaw(&baseIn, "now()")

	//insert to payment_deduction
	// lib.AppendComma(&baseInPaymentDeduction, &baseParamPaymentDeduction, "?", insertIdStringPaymentDeduction)
	// lib.AppendComma(&baseInPaymentDeduction, &baseParamPaymentDeduction, "?", insertIdString)
	// lib.AppendComma(&baseInPaymentDeduction, &baseParamPaymentDeduction, "?", req.PaymentDeduction.PaymentDeductionTypeID)
	// lib.AppendComma(&baseInPaymentDeduction, &baseParamPaymentDeduction, "?", req.PaymentDeduction.PPH)
	// lib.AppendComma(&baseInPaymentDeduction, &baseParamPaymentDeduction, "?", req.PaymentDeduction.AdminFee)

	// else if len(req.ListItemPrice) > 0 {

	// 	bulkInsertQuery := "INSERT INTO item_price (item_price.item_id, item_price.account_id,item_price.price,item_price.server_id,item_price.tiering, item_price.last_update_username) VALUES "
	// 	var paramsBulkInsert []interface{}
	// 	var stringGroup []string
	// 	for _, each := range req.ListItemPrice {
	// 		partquery := "(?, ?, ?, ?, ?, ?)"
	// 		paramsBulkInsert = append(paramsBulkInsert, each.ItemID)
	// 		paramsBulkInsert = append(paramsBulkInsert, each.AccountID)
	// 		paramsBulkInsert = append(paramsBulkInsert, each.Price)
	// 		paramsBulkInsert = append(paramsBulkInsert, each.ServerID)
	// 		paramsBulkInsert = append(paramsBulkInsert, "0")
	// 		paramsBulkInsert = append(paramsBulkInsert, req.LastUpdateUsername)
	// 		stringGroup = append(stringGroup, partquery)
	// 	}

	// 	final_query := bulkInsertQuery + strings.Join(stringGroup, ", ") + " ON DUPLICATE KEY UPDATE price = VALUES(item_price.price), last_update_username = VALUES(item_price.last_update_username)"
	// 	_, _, errInsert := conn.DBAppConn.Exec(final_query, paramsBulkInsert...)

	qry := "INSERT INTO payment (payment_id, invoice_id, total, note, created_by, username, payment_date, payment_type, clearing_date, card_number, payment_method, created_at) VALUES (" + baseIn + ")"
	_, _, errInsert := conn.DBAppConn.Exec(qry, baseParam...)
	if errInsert != nil {
		return errInsert
	}

	if len(req.PaymentDeduction) > 0 {
		bulkInsertQuery := "INSERT INTO payment_deduction (payment_deduction_type_id, payment_id, amount, description) VALUES "
		var paramsBulkInsert []interface{}
		var stringGroup []string
		for _, each := range req.PaymentDeduction {
			partquery := "(?, ?, ?, ?)"
			paramsBulkInsert = append(paramsBulkInsert, each.PaymentDeductionTypeID)
			paramsBulkInsert = append(paramsBulkInsert, insertIdString)
			paramsBulkInsert = append(paramsBulkInsert, each.Amount)
			paramsBulkInsert = append(paramsBulkInsert, each.Description)
			// paramsBulkInsert = append(paramsBulkInsert, "0")
			// paramsBulkInsert = append(paramsBulkInsert, req.LastUpdateUsername)
			stringGroup = append(stringGroup, partquery)
		}

		final_query_bulk := bulkInsertQuery + strings.Join(stringGroup, ", ")
		_, _, errInsert := conn.DBAppConn.Exec(final_query_bulk, paramsBulkInsert...)
		if errInsert != nil {

		}
	}

	// qryPaymentDeduction := "INSERT INTO payment_deduction (payment_deduction_type_id, payment_id, amount, description) VALUES (" + baseInPaymentDeduction + ")"
	// _, _, errInsertPaymentDeduction := conn.DBAppConn.Exec(qryPaymentDeduction, baseParamPaymentDeduction...)
	// if errInsertPaymentDeduction != nil {
	// 	return errInsertPaymentDeduction
	// }

	// _, _, errUpdateId := conn.DBAppConn.Exec("UPDATE control_id set last_id=? where control_id.key=?", insertIdString, "payment")
	// if errUpdateId != nil {
	// 	return errUpdateId
	// }
	err = UpdateControlId(conn, insertIdString, "payment")

	// err = UpdateControlId(conn, insertIdStringPaymentDeduction, "payment_deduction")

	// qryGetSubTotal := "SELECT IFNULL(SUM(invoice_detail.item_price * invoice_detail.qty),0) FROM invoice_detail where invoice_id=?"
	// subTotal, _ := conn.DBAppConn.GetFirstData(qryGetSubTotal, req.InvoiceID)

	qryGetSudahDibayar := "SELECT IFNULL(SUM(payment.total),0) FROM payment where invoice_id=?"
	sudahDibayar, _ := conn.DBAppConn.GetFirstData(qryGetSudahDibayar, req.InvoiceID)

	// subTotalFloat, err := strconv.ParseFloat(subTotal, 64)
	sudahDibayarFloat, err := strconv.ParseFloat(sudahDibayar, 64)

	// qryGetSudahDibayar := "SELECT SUM(total)"
	// qryGetInvoiceData := "SELECT discount, discount_type, ppn FROM invoice WHERE invoice.invoice_id = ?"
	// resInvoice, _, errGetInvoiceData := conn.DBAppConn.SelectQueryByFieldNameSlice(qryGetInvoiceData, req.InvoiceID)
	// if errGetInvoiceData != nil {
	// 	return errGetInvoiceData
	// }

	qryGetInvoiceData := "SELECT grand_total FROM invoice WHERE invoice.invoice_id = ?"
	resInvoice, _, errGetInvoiceData := conn.DBAppConn.SelectQueryByFieldNameSlice(qryGetInvoiceData, req.InvoiceID)
	if errGetInvoiceData != nil {
		return errGetInvoiceData
	}

	oldInvoice := resInvoice[0]
	// discountType := oldInvoice["discount_type"]
	grandTotal := oldInvoice["grand_total"]
	grandTotalFloat, err := strconv.ParseFloat(grandTotal, 64)
	// discount, _ := strconv.ParseFloat(oldInvoice["discount"], 64)
	// ppn, _ := strconv.ParseFloat(oldInvoice["ppn"], 64)
	// if discountType == "percent" {
	// 	discount = discount * subTotalFloat / 100
	// 	logrus.Info("DiscountSebelum- ", discount)
	// 	discount = math.Ceil(discount*100) / 100
	// 	logrus.Info("FinalDiscountPercent-", discount)

	// }

	// logrus.Info("Calculate Discount- ", discount)
	// newSubTotal := subTotalFloat - discount
	// logrus.Info("NewSubTotal-", newSubTotal)
	// ppn = math.Ceil((ppn*newSubTotal/100)*100) / 100

	// grandTotal := math.Ceil((newSubTotal+ppn)*100) / 100
	// logrus.Info("GrandTotal-", grandTotal)

	if sudahDibayarFloat >= grandTotalFloat {
		qryUpdateInvoicePaymentStatus := "UPDATE invoice SET invoice.paid = ? WHERE invoice.invoice_id = ?"
		_, _, errUpdateInvoice := conn.DBAppConn.Exec(qryUpdateInvoicePaymentStatus, "1", req.InvoiceID)
		if errUpdateInvoice != nil {
			return errUpdateInvoice
		}
	}

	return err
}

func CheckPaymentNominal(conn *connections.Connections, req datastruct.PaymentRequest) error {
	var err error
	// qryGetSubTotal := "SELECT IFNULL(SUM(invoice_detail.item_price * invoice_detail.qty),0) FROM invoice_detail where invoice_id=?"
	// subTotal, _ := conn.DBAppConn.GetFirstData(qryGetSubTotal, req.InvoiceID)

	qryGetSudahDibayar := "SELECT IFNULL(SUM(payment.total),0) FROM payment where invoice_id=?"
	sudahDibayar, _ := conn.DBAppConn.GetFirstData(qryGetSudahDibayar, req.InvoiceID)

	// subTotalFloat, err := strconv.ParseFloat(subTotal, 64)
	sudahDibayarFloat, err := strconv.ParseFloat(sudahDibayar, 64)

	// qryGetSudahDibayar := "SELECT SUM(total)"
	qryGetInvoiceData := "SELECT grand_total FROM invoice WHERE invoice.invoice_id = ?"
	resInvoice, _, errGetInvoiceData := conn.DBAppConn.SelectQueryByFieldNameSlice(qryGetInvoiceData, req.InvoiceID)
	if errGetInvoiceData != nil {
		return errGetInvoiceData
	}

	oldInvoice := resInvoice[0]
	grandTotal := oldInvoice["grand_total"]
	grandTotalFloat, err := strconv.ParseFloat(grandTotal, 64)
	// discountType := oldInvoice["discount_type"]
	// discount, _ := strconv.ParseFloat(oldInvoice["discount"], 64)
	// ppn, _ := strconv.ParseFloat(oldInvoice["ppn"], 64)
	// if discountType == "percent" {
	// 	discount = discount * subTotalFloat / 100
	// 	logrus.Info("DiscountSebelum- ", discount)
	// 	discount = math.Ceil(discount*100) / 100

	// }

	// newSubTotal := subTotalFloat - discount

	// ppn = math.Ceil((ppn*newSubTotal/100)*100) / 100

	// grandTotal := math.Ceil((newSubTotal+ppn)*100) / 100

	sisa := grandTotalFloat - sudahDibayarFloat
	totalAkanDibayarFloat, err := strconv.ParseFloat(req.Total, 64)
	if totalAkanDibayarFloat > sisa {
		return errors.New("The payment amount exceeds the remaining bill")
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
	qry := "UPDATE payment SET payment.status = 1 WHERE payment_id = ?"
	_, _, err = conn.DBAppConn.Exec(qry, req.PaymentID)
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
