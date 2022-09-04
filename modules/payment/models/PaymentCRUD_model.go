package models

import (
	"billingdashboard/connections"
	"billingdashboard/lib"
	"billingdashboard/modules/payment/datastruct"
	"errors"
	"strconv"
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
		runQuery = "SELECT payment_id, payment.invoice_id, payment.payment_date, payment.total, payment.note, payment.created_by, username,  payment.payment_method, payment.card_number, invoice.account_id, invoice.invoice_no, account.name as account_name, payment.payment_method, payment.card_number FROM payment JOIN invoice ON invoice.invoice_id = payment.invoice_id JOIN account ON account.account_id = invoice.account_id "
	} else {
		runQuery = "SELECT payment_id, payment.invoice_id, payment.payment_date, payment.total, payment.note, payment.created_by, username,  payment.payment_method, payment.card_number, invoice.account_id, invoice.invoice_no, account.name as account_name, payment.payment_method, payment.card_number FROM payment JOIN invoice ON invoice.invoice_id = payment.invoice_id JOIN account ON account.account_id = invoice.account_id "
	}

	if len(baseWhere) > 0 {
		runQuery += "WHERE " + baseWhere
	}

	lib.AppendOrderBy(&runQuery, req.Param.OrderBy, req.Param.OrderDir)
	lib.AppendLimit(&runQuery, req.Param.Page, req.Param.PerPage)

	resultSelect, _, err := conn.DBAppConn.SelectQueryByFieldNameSlice(runQuery, baseParam...)
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

		invoice := make(map[string]interface{})
		invoice["invoice_no"] = each["invoice_no"]

		account := make(map[string]interface{})
		account["name"] = each["account_name"]
		invoice["account"] = account
		single["invoice"] = invoice
		result = append(result, single)
	}

	return result, err
}

func InsertPayment(conn *connections.Connections, req datastruct.PaymentRequest) error {
	var err error

	// -- THIS IS BASIC INSERT EXAMPLE
	var baseIn string
	var baseParam []interface{}

	var baseInPaymentDeduction string
	var baseParamPaymentDeduction []interface{}

	lastId, _ := conn.DBAppConn.GetFirstData("SELECT last_id FROM control_id where control_id.key=?", "payment")
	// lastId, _ := conn.DBAppConn.GetFirstData("SELECT last_id FROM control_id where control_id.key=?", "account")

	lastIdPaymentDeduction, _ := conn.DBAppConn.GetFirstData("SELECT last_id FROM control_id where control_id.key=?", "payment_deduction")

	intLastId, err := strconv.Atoi(lastId)
	insertId := intLastId + 1

	insertIdString := strconv.Itoa(insertId)

	intLastIdPaymentDeduction, err := strconv.Atoi(lastIdPaymentDeduction)
	insertIdPaymentDeduction := intLastIdPaymentDeduction + 1

	insertIdStringPaymentDeduction := strconv.Itoa(insertIdPaymentDeduction)

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
	lib.AppendComma(&baseInPaymentDeduction, &baseParamPaymentDeduction, "?", insertIdStringPaymentDeduction)
	lib.AppendComma(&baseInPaymentDeduction, &baseParamPaymentDeduction, "?", insertIdString)
	lib.AppendComma(&baseInPaymentDeduction, &baseParamPaymentDeduction, "?", req.PaymentDeduction.PPN)
	lib.AppendComma(&baseInPaymentDeduction, &baseParamPaymentDeduction, "?", req.PaymentDeduction.PPH)
	lib.AppendComma(&baseInPaymentDeduction, &baseParamPaymentDeduction, "?", req.PaymentDeduction.AdminFee)

	qry := "INSERT INTO payment (payment_id, invoice_id, total, note, created_by, username, payment_date, payment_type, clearing_date, card_number, payment_method, created_at) VALUES (" + baseIn + ")"
	_, _, errInsert := conn.DBAppConn.Exec(qry, baseParam...)
	if errInsert != nil {
		return errInsert
	}

	qryPaymentDeduction := "INSERT INTO payment_deduction (payment_deduction_id, payment_id, ppn, pph, admin_fee) VALUES (" + baseInPaymentDeduction + ")"
	_, _, errInsertPaymentDeduction := conn.DBAppConn.Exec(qryPaymentDeduction, baseParamPaymentDeduction...)
	if errInsertPaymentDeduction != nil {
		return errInsertPaymentDeduction
	}

	// _, _, errUpdateId := conn.DBAppConn.Exec("UPDATE control_id set last_id=? where control_id.key=?", insertIdString, "payment")
	// if errUpdateId != nil {
	// 	return errUpdateId
	// }
	err = UpdateControlId(conn, insertIdString, "payment")

	err = UpdateControlId(conn, insertIdStringPaymentDeduction, "payment_deduction")

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
	// qry := "DELETE FROM payment WHERE paymentid = ?"
	// _, _, err = conn.DBAppConn.Exec(qry, req.PaymentID)
	return err
}
