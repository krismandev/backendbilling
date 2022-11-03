package models

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/lib"
	"backendbillingdashboard/modules/invoice/datastruct"
	"strconv"
)

func GetInvoiceFromRequest(conn *connections.Connections, req datastruct.InvoiceRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	var resultQuery []map[string]string
	var err error

	// -- THIS IS BASIC GET REQUEST EXAMPLE LOGIC
	var baseWhere string
	var baseParam []interface{}

	lib.AppendWhere(&baseWhere, &baseParam, "invoice_id = ?", req.InvoiceID)
	lib.AppendWhere(&baseWhere, &baseParam, "invoice.invoice_no = ?", req.InvoiceNo)
	lib.AppendWhere(&baseWhere, &baseParam, "invoicestatus = ?", "A")
	lib.AppendWhere(&baseWhere, &baseParam, "invoice.account_id = ?", req.AccountID)
	lib.AppendWhere(&baseWhere, &baseParam, "invoice.paid = ?", req.Paid)
	lib.AppendWhere(&baseWhere, &baseParam, "invoice.invoice_date = ?", req.InvoiceDate)
	lib.AppendWhere(&baseWhere, &baseParam, "invoice.inv_type_id = ?", req.InvoiceTypeID)
	lib.AppendWhere(&baseWhere, &baseParam, "invoice.due_date > curdate() and invoice.due_date <= curdate() + interval ? day", req.ApproachingDueDateInterval)
	if len(req.ListInvoiceID) > 0 {
		var baseIn string
		for _, prid := range req.ListInvoiceID {
			lib.AppendComma(&baseIn, &baseParam, "?", prid)
		}
		lib.AppendWhereRaw(&baseWhere, "invoiceid IN ("+baseIn+")")
	}

	//TERAKHIR NAMBAH QUERY BUAT NAMPILIN ACCOUNT
	runQuery := `SELECT invoice.invoice_id, invoice.invoice_no, invoice.invoice_date, invoice.invoicestatus, invoice.account_id, invoice.month_use, 
	invoice.inv_type_id, invoice.printcounter, invoice.note, invoice.canceldesc, invoice.payment_method ,invoice.last_print_username, 
	invoice.last_print_date, invoice.created_at, invoice.created_by, invoice.last_update_username, invoice.exchange_rate_date, invoice.grand_total, 
	invoice.ppn_amount, invoice.last_update_date, invoice.discount_type, invoice.discount, invoice.ppn ,invoice.paid, invoice.due_date , 
	invoice_type.inv_type_id as tblinvoice_type_inv_type_id, invoice_type.inv_type_name, invoice_type.server_id as tblinvoice_type_server_id, invoice_type.category, 
	invoice_type.load_from_server, invoice_type.currency_code, account.account_id as tblaccount_account_id, account.name as tblaccount_name, account.address1, account.address2, 
	account.city, account.phone, account.contact_person, account.contact_person_phone FROM invoice JOIN invoice_type ON 
	invoice.inv_type_id = invoice_type.inv_type_id JOIN account ON invoice.account_id = account.account_id `
	if len(baseWhere) > 0 {
		runQuery += "WHERE " + baseWhere
	}
	lib.AppendOrderBy(&runQuery, req.Param.OrderBy, req.Param.OrderDir)
	lib.AppendLimit(&runQuery, req.Param.Page, req.Param.PerPage)

	resultQuery, _, err = conn.DBAppConn.SelectQueryByFieldNameSlice(runQuery, baseParam...)
	for _, each := range resultQuery {
		single := make(map[string]interface{})
		single["invoice_id"] = each["invoice_id"]
		single["invoice_no"] = each["invoice_no"]
		single["invoice_date"] = each["invoice_date"]
		single["invoicestatus"] = each["invoicestatus"]
		single["account_id"] = each["account_id"]
		first4 := each["month_use"][0:4]

		last2 := each["month_use"][len(each["month_use"])-2:]

		single["month_use"] = first4 + "-" + last2
		single["inv_type_id"] = each["inv_type_id"]
		single["printcounter"] = each["printcounter"]
		single["note"] = each["note"]
		single["canceldesc"] = each["canceldesc"]
		single["ppn"] = each["ppn"]
		single["last_print_username"] = each["last_print_username"]
		single["last_print_date"] = each["last_print_date"]
		single["created_at"] = each["created_at"]
		single["created_by"] = each["created_by"]
		single["last_update_username"] = each["last_update_username"]
		single["last_update_date"] = each["last_update_date"]
		single["discount_type"] = each["discount_type"]
		single["discount"] = each["discount"]
		single["paid"] = each["paid"]
		single["payment_method"] = each["payment_method"]
		single["exchange_rate_date"] = each["exchange_rate_date"]
		single["due_date"] = each["due_date"]
		single["grand_total"] = each["grand_total"]
		single["ppn_amount"] = each["ppn_amount"]

		invType := make(map[string]interface{})
		invType["inv_type_id"] = each["tblinvoice_type_inv_type_id"]
		invType["inv_type_name"] = each["inv_type_name"]
		invType["server_id"] = each["tblinvoice_type_server_id"]
		invType["category"] = each["category"]
		invType["load_from_server"] = each["load_from_server"]
		invType["currency_code"] = each["currency_code"]

		single["invoice_type"] = invType

		account := make(map[string]interface{})
		account["account_id"] = each["tblaccount_account_id"]
		account["name"] = each["tblaccount_name"]
		account["address1"] = each["address1"]
		account["address2"] = each["address2"]
		account["city"] = each["city"]
		account["phone"] = each["phone"]
		account["contact_person"] = each["contact_person"]
		account["contact_person_phone"] = each["contact_person_phone"]

		single["account"] = account

		qryGetDetail := "SELECT invoice_detail.invoice_detail_id, invoice_detail.invoice_id, invoice_detail.itemid, invoice_detail.qty, invoice_detail.uom, invoice_detail.item_price, tiering, invoice_detail.note, invoice_detail.balance_type, invoice_detail.server_id ,item.item_id as tblitem_item_id, item.item_name, item.operator, item.route, item.category, item.uom as tblitem_uom, server.server_id as tblserver_server_id, server.server_name as tblserver_server_name FROM invoice_detail JOIN item ON invoice_detail.itemid = item.item_id  JOIN server ON invoice_detail.server_id = server.server_id WHERE invoice_id = ?"
		resultGetDetail, _, errGetDetail := conn.DBAppConn.SelectQueryByFieldNameSlice(qryGetDetail, single["invoice_id"])
		if errGetDetail != nil {
			return nil, errGetDetail
		}

		var tampungDetail []map[string]interface{}
		for _, detail := range resultGetDetail {
			singleDetail := make(map[string]interface{})
			singleDetail["invoice_detail_id"] = detail["invoice_detail_id"]
			singleDetail["invoice_id"] = detail["invoice_id"]
			singleDetail["itemid"] = detail["itemid"]
			singleDetail["qty"] = detail["qty"]
			singleDetail["uom"] = detail["uom"]
			singleDetail["item_price"] = detail["item_price"]
			singleDetail["tiering"] = detail["tiering"]
			singleDetail["note"] = detail["note"]
			singleDetail["balance_type"] = detail["balance_type"]
			singleDetail["server_id"] = detail["server_id"]

			item := make(map[string]interface{})
			item["item_id"] = detail["tblitem_item_id"]
			item["item_name"] = detail["item_name"]
			item["operator"] = detail["operator"]
			item["route"] = detail["route"]
			item["category"] = detail["category"]
			item["uom"] = detail["uom"]

			server := make(map[string]interface{})
			server["server_id"] = detail["tblserver_server_id"]
			server["server_name"] = detail["tblserver_server_name"]

			singleDetail["item"] = item
			singleDetail["server"] = server

			tampungDetail = append(tampungDetail, singleDetail)

		}
		single["list_invoice_detail"] = tampungDetail

		// var category
		result = append(result, single)
	}
	return result, err
}

func InsertInvoice(conn *connections.Connections, req datastruct.InvoiceRequest) error {
	var err error

	// -- THIS IS BASIC INSERT EXAMPLE
	var baseIn string
	var baseParam []interface{}

	lastId, _ := conn.DBAppConn.GetFirstData("SELECT last_id FROM control_id where control_id.key=?", "invoice")
	// lastId, _ := conn.DBAppConn.GetFirstData("SELECT last_id FROM control_id where control_id.key=?", "account")

	intLastId, err := strconv.Atoi(lastId)
	insertId := intLastId + 1

	insertIdString := strconv.Itoa(insertId)

	lib.AppendComma(&baseIn, &baseParam, "?", insertIdString)
	lib.AppendComma(&baseIn, &baseParam, "?", req.InvoiceDate)
	lib.AppendComma(&baseIn, &baseParam, "?", "A")
	lib.AppendComma(&baseIn, &baseParam, "?", req.AccountID)
	lib.AppendComma(&baseIn, &baseParam, "?", req.MonthUse)
	lib.AppendComma(&baseIn, &baseParam, "?", req.InvoiceTypeID)
	lib.AppendCommaRaw(&baseIn, "now()")
	lib.AppendComma(&baseIn, &baseParam, "?", req.DiscountType)
	lib.AppendComma(&baseIn, &baseParam, "?", req.Discount)
	lib.AppendComma(&baseIn, &baseParam, "?", req.Note)
	lib.AppendCommaRaw(&baseIn, "now()")
	createdBy := req.LastUpdateUsername
	lib.AppendComma(&baseIn, &baseParam, "?", createdBy)
	lib.AppendComma(&baseIn, &baseParam, "?", "0")
	lib.AppendComma(&baseIn, &baseParam, "?", req.PPN)
	lib.AppendComma(&baseIn, &baseParam, "?", req.PaymentMethod)
	lib.AppendComma(&baseIn, &baseParam, "?", req.ExchangeRateDate)
	lib.AppendComma(&baseIn, &baseParam, "?", req.DueDate)

	// qryCheckControlIdPeriod := "SELECT control_id.key, control_id.period ,last_id FROM control_id WHERE control_id.key = ? AND period = ?"
	// resCheck, countControlIdPeriod, errCheckControlIdPeriod := conn.DBAppConn.SelectQueryByFieldName(qryCheckControlIdPeriod, "invoice_no", req.MonthUse)
	// if errCheckControlIdPeriod != nil {
	// 	return errCheckControlIdPeriod
	// }

	// lastIdPeriod := resCheck[1]["last_id"]

	// if countControlIdPeriod == 0 {
	// 	errInsertControlIdPeriod := InsertControlIdPeriod(conn, req.MonthUse)
	// 	if errInsertControlIdPeriod != nil {
	// 		return errInsertControlIdPeriod
	// 	}
	// }

	qry := "INSERT INTO invoice (invoice_id, invoice_date,invoicestatus, account_id, month_use ,inv_type_id, last_update_date, discount_type, discount, invoice.note, created_at,created_by, printcounter,ppn, payment_method, exchange_rate_date, due_date) VALUES (" + baseIn + ")"
	_, _, errInsert := conn.DBAppConn.Exec(qry, baseParam...)
	if errInsert != nil {
		return errInsert
	}
	_, _, errUpdateId := conn.DBAppConn.Exec("UPDATE control_id set last_id=? where control_id.key=?", insertIdString, "invoice")
	if errUpdateId != nil {
		return errUpdateId
	}

	if len(req.ListInvoiceDetail) > 0 {
		oldLastIdInvoiceDetail, errCheckLastIdInvoiceDetail := CheckControlIdInvoiceDetail(conn)
		if errCheckLastIdInvoiceDetail != nil {
			return errCheckLastIdInvoiceDetail
		}
		tempLastIdInvoiceDetail, errConvLastIdInvoiceDetail := strconv.Atoi(oldLastIdInvoiceDetail)
		if errConvLastIdInvoiceDetail != nil {
			return errCheckLastIdInvoiceDetail
		}
		for _, each := range req.ListInvoiceDetail {
			tempLastIdInvoiceDetail += 1
			qryInsertDetail := "INSERT INTO invoice_detail (invoice_detail.invoice_detail_id, invoice_detail.invoice_id, invoice_detail.itemid, invoice_detail.qty, invoice_detail.uom, invoice_detail.item_price,invoice_detail.note, invoice_detail.balance_type, invoice_detail.server_id ,last_update_username) VALUES (?,?, ?, ?, ?, ?, ?, ?, ?, ?)"
			var invDetailParam []interface{} = []interface{}{
				tempLastIdInvoiceDetail,
				insertIdString,
				each.ItemID,
				each.Qty,
				each.Uom,
				each.ItemPrice,
				each.Note,
				each.BalanceType,
				each.ServerID,
				createdBy,
			}

			_, _, errInsertDetail := conn.DBAppConn.Exec(qryInsertDetail, invDetailParam...)
			if errInsertDetail != nil {
				return errInsertDetail
			}
		}

		newLastIdInvoiceDetailStr := strconv.Itoa(tempLastIdInvoiceDetail)
		errUpdateControlIdInvoiceDetail := UpdateControlId(conn, newLastIdInvoiceDetailStr, "invoice_detail")
		if errUpdateControlIdInvoiceDetail != nil {
			return errUpdateControlIdInvoiceDetail
		}

		// if countControlIdPeriod == 0 {
		// 	invoiceNo := req.MonthUse + "0001"
		// 	logrus.Info("InvoiceNo-", invoiceNo)
		// 	qryUpdateInvoice := "UPDATE invoice set invoice_no = ? WHERE invoice_id = ?"
		// 	_, _, errUpdateInvoice := conn.DBAppConn.Exec(qryUpdateInvoice, invoiceNo, insertIdString)
		// 	if errUpdateInvoice != nil {
		// 		return errUpdateInvoice
		// 	}
		// 	_, _, errUpdateLastIdPeriod := conn.DBAppConn.Exec("UPDATE control_id set last_id = ? WHERE control_id.key=? AND control_id.period = ?", invoiceNo, "invoice_no", req.MonthUse)
		// 	if errUpdateLastIdPeriod != nil {
		// 		return errUpdateLastIdPeriod
		// 	}
		// } else {
		// 	//ambil 6 angka pertama ( atau bisa dibilang month use nya)
		// 	first6 := lastIdPeriod[0:6]
		// 	//4 angka terakhir sebagai id increment nya
		// 	last4 := lastIdPeriod[len(lastIdPeriod)-4:]
		// 	last4Int, errConv := strconv.Atoi(last4)
		// 	if errConv != nil {
		// 		return errConv
		// 	}
		// 	last4Int += 1
		// 	newLast4IdStr := strconv.Itoa(last4Int)
		// 	//digabung kembali
		// 	newInvoiceNo := first6 + lib.StrPadLeft(newLast4IdStr, 4, "0")

		// 	qryUpdateInvoice := "UPDATE invoice set invoice_no = ? WHERE invoice_id = ?"
		// 	_, _, errUpdateInvoice := conn.DBAppConn.Exec(qryUpdateInvoice, newInvoiceNo, insertIdString)
		// 	if errUpdateInvoice != nil {
		// 		return errUpdateInvoice
		// 	}

		// 	_, _, errUpdateLastIdPeriod := conn.DBAppConn.Exec("UPDATE control_id set last_id = ? WHERE control_id.key=? AND control_id.period = ?", newInvoiceNo, "invoice_no", req.MonthUse)
		// 	if errUpdateLastIdPeriod != nil {
		// 		return errUpdateLastIdPeriod
		// 	}
		// }

		qryUpdateServerData := "UPDATE server_data set server_data.invoice_id = ? WHERE server_data.account_id = ? AND server_data.server_id = ? AND DATE_FORMAT(server_data.external_transdate, '%Y%m') = ?"
		var updateServerDataParam []interface{} = []interface{}{
			insertIdString,
			req.AccountID,
			req.ServerID,
			req.MonthUse,
		}
		_, _, errUpdateServerData := conn.DBAppConn.Exec(qryUpdateServerData, updateServerDataParam...)
		if errUpdateServerData != nil {
			return errUpdateServerData
		}

	}

	return err
}

func InsertControlIdPeriod(conn *connections.Connections, monthUse string) error {
	var err error
	qryInserControlIdPeriod := "INSERT INTO control_id (control_id.key,control_id.period,control_id.last_id) VALUES (?,?,?)"
	_, _, err = conn.DBAppConn.Exec(qryInserControlIdPeriod, "invoice_no", monthUse, "0")
	if err != nil {
		return err
	}

	return err
}

func CheckControlIdInvoiceDetail(conn *connections.Connections) (string, error) {
	qryCheckControlIdInvoiceDetail := "SELECT last_id FROM control_id WHERE control_id.key = ?"
	resCheck, countControlIdInvoiceDetail, errCheckControlIdInvoiceDetail := conn.DBAppConn.SelectQueryByFieldName(qryCheckControlIdInvoiceDetail, "invoice_detail")
	if errCheckControlIdInvoiceDetail != nil {
		return "", errCheckControlIdInvoiceDetail
	}

	if countControlIdInvoiceDetail == 0 {
		qryInserControlIdInvoiceDetail := "INSERT INTO control_id (control_id.key,control_id.period,control_id.last_id) VALUES (?,?,?)"
		_, _, errInsert := conn.DBAppConn.Exec(qryInserControlIdInvoiceDetail, "invoice_detail", "0", "0")
		if errInsert != nil {
			return "", errInsert
		}

		return "0", nil
	}

	lastIdInvoiceDetail := resCheck[1]["last_id"]
	return lastIdInvoiceDetail, nil
}

func UpdateControlId(conn *connections.Connections, newId string, key string) error {
	var err error
	_, _, err = conn.DBAppConn.Exec("UPDATE control_id set last_id = ? WHERE control_id.key=?", newId, key)
	return err
}

func UpdateInvoice(conn *connections.Connections, req datastruct.InvoiceRequest) error {
	var err error

	var baseUpInvoice string
	var baseParam []interface{}

	lib.AppendComma(&baseUpInvoice, &baseParam, "invoice_no = ?", req.InvoiceNo)
	lib.AppendComma(&baseUpInvoice, &baseParam, "invoice_date = ?", req.InvoiceDate)
	lib.AppendComma(&baseUpInvoice, &baseParam, "discount_type = ?", req.DiscountType)
	lib.AppendComma(&baseUpInvoice, &baseParam, "discount = ?", req.Discount)
	lib.AppendComma(&baseUpInvoice, &baseParam, "invoice.note = ?", req.Note)
	lib.AppendComma(&baseUpInvoice, &baseParam, "last_update_username = ?", req.LastUpdateUsername)
	lib.AppendComma(&baseUpInvoice, &baseParam, "ppn = ?", req.PPN)
	lib.AppendComma(&baseUpInvoice, &baseParam, "payment_method = ?", req.PaymentMethod)

	errCheckInvoiceNo := CheckInvoiceNoDuplicate(req.InvoiceNo, conn, req)
	if errCheckInvoiceNo != nil {
		return errCheckInvoiceNo
	}

	qry := "UPDATE invoice SET " + baseUpInvoice + " WHERE invoice_id = ?"
	baseParam = append(baseParam, req.InvoiceID)
	// baseParam = append(baseParam, req.InvoiceNo)
	// baseParam = append(baseParam, req.TransDate)
	// baseParam = append(baseParam, req.DiscountType)
	// baseParam = append(baseParam, req.Discount)
	// baseParam = append(baseParam, req.Note)
	// baseParam = append(baseParam, req.LastUpdateUsername)
	// baseParam = append(baseParam, req.PPN)
	// baseParam = append(baseParam, "now()")
	// baseParam = append(baseParam, req.InvoiceID)
	_, _, errUpdateInvoice := conn.DBAppConn.Exec(qry, baseParam...)
	if errUpdateInvoice != nil {
		return errUpdateInvoice
	}

	if len(req.ListInvoiceDetail) > 0 {
		for _, each := range req.ListInvoiceDetail {

			qryUpdateDetail := "UPDATE invoice_detail set invoice_detail.qty = ?, invoice_detail.item_price = ?,invoice_detail.note = ?, invoice_detail.last_update_username = ? WHERE invoice_detail_id = ?"
			var invDetailParam []interface{} = []interface{}{
				each.Qty,
				each.ItemPrice,
				each.Note,
				req.LastUpdateUsername,
				each.InvoiceDetailID,
			}

			_, _, errUpdateDetail := conn.DBAppConn.Exec(qryUpdateDetail, invDetailParam...)
			if errUpdateDetail != nil {
				return errUpdateDetail
			}
		}
	}

	return err
}

func DeleteInvoice(conn *connections.Connections, req datastruct.InvoiceRequest) error {
	var err error
	// -- THIS IS DELETE LOGIC EXAMPLE
	qry := "UPDATE invoice SET invoicestatus = ?, canceldesc = ? WHERE invoice_id = ?"
	_, _, err = conn.DBAppConn.Exec(qry, "D", req.CancelDesc, req.InvoiceID)
	return err
}

func CancelInvoice(conn *connections.Connections, req datastruct.InvoiceRequest) error {
	var err error

	qry := "UPDATE invoice SET invoicestatus = ?, canceldesc = ? WHERE invoice_id = ?"
	qryUpdateServerData := "UPDATE server_data SET invoice_id = NULL where server_data.invoice_id = ?"
	_, _, err = conn.DBAppConn.Exec(qry, "D", req.CancelDesc, req.InvoiceID)

	_, _, err = conn.DBAppConn.Exec(qryUpdateServerData, req.InvoiceID)
	return err
}

func PrintInvoice(conn *connections.Connections, req datastruct.InvoiceRequest) error {
	var err error

	qry := "UPDATE invoice SET printcounter = printcounter + 1, last_print_username = ?, last_print_date = now(), grand_total = ?, ppn_amount = ?  WHERE invoice_id = ?"
	_, _, err = conn.DBAppConn.Exec(qry, req.LastUpdateUsername, req.GrandTotal, req.PPNAmount, req.InvoiceID)
	return err
}
