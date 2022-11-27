package processors

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/invoice/datastruct"
	"backendbillingdashboard/modules/invoice/models"
	dtServer "backendbillingdashboard/modules/server/datastruct"
)

func GetListInvoice(conn *connections.Connections, req datastruct.InvoiceRequest) ([]datastruct.InvoiceDataStruct, error) {
	var output []datastruct.InvoiceDataStruct
	var err error

	// grab mapping data from model
	invoiceList, err := models.GetInvoiceFromRequest(conn, req)
	if err != nil {
		return output, err
	}

	for _, invoice := range invoiceList {
		single := CreateSingleInvoiceStruct(invoice)
		output = append(output, single)
	}

	return output, err
}

func CreateSingleInvoiceStruct(invoice map[string]interface{}) datastruct.InvoiceDataStruct {
	var single datastruct.InvoiceDataStruct
	single.InvoiceID, _ = invoice["invoice_id"].(string)
	single.InvoiceNo, _ = invoice["invoice_no"].(string)
	single.InvoiceDate, _ = invoice["invoice_date"].(string)
	single.InvoiceStatus, _ = invoice["invoicestatus"].(string)
	single.AccountID, _ = invoice["account_id"].(string)
	single.MonthUse, _ = invoice["month_use"].(string)
	single.InvoiceTypeID, _ = invoice["inv_type_id"].(string)
	single.PrintCounter, _ = invoice["printcounter"].(string)
	single.PPN, _ = invoice["ppn"].(string)
	single.Note, _ = invoice["note"].(string)
	single.CancelDesc, _ = invoice["canceldesc"].(string)
	single.LastPrintUsername, _ = invoice["last_print_username"].(string)
	single.LastPrintDate, _ = invoice["last_print_date"].(string)
	single.CreatedAt, _ = invoice["created_at"].(string)
	single.CreatedBy, _ = invoice["created_by"].(string)
	single.LastUpdateUsername, _ = invoice["last_update_username"].(string)
	single.LastUpdateDate, _ = invoice["last_update_date"].(string)
	single.DiscountType, _ = invoice["discount_type"].(string)
	single.Discount, _ = invoice["discount"].(string)
	single.Paid, _ = invoice["paid"].(string)
	single.PaymentMethod, _ = invoice["payment_method"].(string)
	single.ExchangeRateDate, _ = invoice["exchange_rate_date"].(string)
	single.DueDate, _ = invoice["due_date"].(string)
	single.GrandTotal, _ = invoice["grand_total"].(string)
	single.PPNAmount, _ = invoice["ppn_amount"].(string)
	single.Sender, _ = invoice["sender"].(string)

	var invoiceType datastruct.InvoiceTypeDataStruct
	invoiceType.InvoiceTypeID = invoice["invoice_type"].(map[string]interface{})["inv_type_id"].(string)
	invoiceType.InvoiceTypeName = invoice["invoice_type"].(map[string]interface{})["inv_type_name"].(string)
	invoiceType.ServerID = invoice["invoice_type"].(map[string]interface{})["server_id"].(string)
	invoiceType.Category = invoice["invoice_type"].(map[string]interface{})["category"].(string)
	invoiceType.LoadFromServer = invoice["invoice_type"].(map[string]interface{})["load_from_server"].(string)
	invoiceType.CurrencyCode = invoice["invoice_type"].(map[string]interface{})["currency_code"].(string)

	single.InvoiceType = invoiceType

	var account datastruct.AccountDataStruct
	account.AccountID = invoice["account"].(map[string]interface{})["account_id"].(string)
	account.Name = invoice["account"].(map[string]interface{})["name"].(string)
	account.Address1 = invoice["account"].(map[string]interface{})["address1"].(string)
	account.Address2 = invoice["account"].(map[string]interface{})["address2"].(string)
	account.City = invoice["account"].(map[string]interface{})["city"].(string)
	account.ContactPerson = invoice["account"].(map[string]interface{})["contact_person"].(string)
	account.ContactPersonPhone = invoice["account"].(map[string]interface{})["contact_person_phone"].(string)

	single.Account = account

	var tampungDetail []datastruct.InvoiceDetailStruct
	for _, eachDetail := range invoice["list_invoice_detail"].([]map[string]interface{}) {
		var detail datastruct.InvoiceDetailStruct
		detail.InvoiceDetailID = eachDetail["invoice_detail_id"].(string)
		detail.InvoiceID = eachDetail["invoice_id"].(string)
		detail.ItemID = eachDetail["itemid"].(string)
		detail.ItemPrice = eachDetail["item_price"].(string)
		detail.Qty = eachDetail["qty"].(string)
		detail.Uom = eachDetail["uom"].(string)
		detail.Note = eachDetail["note"].(string)
		detail.BalanceType = eachDetail["balance_type"].(string)
		detail.ServerID = eachDetail["server_id"].(string)

		var item datastruct.ItemDataStruct
		item.ItemID = eachDetail["item"].(map[string]interface{})["item_id"].(string)
		item.ItemName = eachDetail["item"].(map[string]interface{})["item_name"].(string)
		item.Operator = eachDetail["item"].(map[string]interface{})["operator"].(string)
		item.Route = eachDetail["item"].(map[string]interface{})["route"].(string)
		item.Category = eachDetail["item"].(map[string]interface{})["category"].(string)
		item.UOM = eachDetail["item"].(map[string]interface{})["uom"].(string)

		var server dtServer.ServerDataStruct
		server.ServerID = eachDetail["server"].(map[string]interface{})["server_id"].(string)
		server.ServerName = eachDetail["server"].(map[string]interface{})["server_name"].(string)

		detail.Item = item
		detail.Server = server

		tampungDetail = append(tampungDetail, detail)
	}
	single.ListInvoiceDetail = tampungDetail

	return single
}

func InsertInvoice(conn *connections.Connections, req datastruct.InvoiceRequest) error {
	var err error

	err = models.InsertInvoice(conn, req)
	if err != nil {
		return err
	}

	// jika tidak ada error, return single instance of single invoice
	// single, err := models.GetSingleInvoice(req.InvoiceID, conn, req)
	// if err != nil {
	// 	return output, err
	// }

	// output = CreateSingleInvoiceStruct(single)
	return err
}

func UpdateInvoice(conn *connections.Connections, req datastruct.InvoiceRequest) error {
	var err error

	err = models.UpdateInvoice(conn, req)
	if err != nil {
		return err
	}

	// jika tidak ada error, return single instance of single invoice
	// single, err := models.GetSingleInvoice(req.InvoiceID, conn, req)
	// if err != nil {
	// 	return output, err
	// }

	// output = CreateSingleInvoiceStruct(single)
	return err
}

func DeleteInvoice(conn *connections.Connections, req datastruct.InvoiceRequest) error {
	err := models.DeleteInvoice(conn, req)
	return err
}

func CancelInvoice(conn *connections.Connections, req datastruct.InvoiceRequest) error {
	err := models.CancelInvoice(conn, req)
	return err
}

func PrintInvoice(conn *connections.Connections, req datastruct.InvoiceRequest) error {
	var err error

	err = models.PrintInvoice(conn, req)
	if err != nil {
		return err
	}

	return err
}
