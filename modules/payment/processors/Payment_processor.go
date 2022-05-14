package processors

import (
	"billingdashboard/connections"
	"billingdashboard/modules/payment/datastruct"
	"billingdashboard/modules/payment/models"
)

func GetListPayment(conn *connections.Connections, req datastruct.PaymentRequest) ([]datastruct.PaymentDataStruct, error) {
	var output []datastruct.PaymentDataStruct
	var err error

	// grab mapping data from model
	paymentList, err := models.GetPaymentFromRequest(conn, req)
	if err != nil {
		return output, err
	}

	for _, payment := range paymentList {
		single := CreateSinglePaymentStruct(payment)
		output = append(output, single)
	}

	return output, err
}

func CreateSinglePaymentStruct(payment map[string]interface{}) datastruct.PaymentDataStruct {
	var single datastruct.PaymentDataStruct
	single.PaymentID, _ = payment["payment_id"].(string)
	single.InvoiceID, _ = payment["invoice_id"].(string)
	single.PaymentDate, _ = payment["payment_date"].(string)
	single.Total, _ = payment["total"].(string)
	single.Note, _ = payment["note"].(string)
	single.CreatedBy, _ = payment["created_by"].(string)
	single.UserName, _ = payment["username"].(string)
	single.PaymentMethod, _ = payment["payment_method"].(string)
	single.CardNumber, _ = payment["card_number"].(string)

	var account datastruct.AccountDataStruct
	account.Name = payment["invoice"].(map[string]interface{})["account"].(map[string]interface{})["name"].(string)

	var invoice datastruct.InvoiceDataStruct
	invoice.InvoiceNo = payment["invoice"].(map[string]interface{})["invoice_no"].(string)
	invoice.Account = account
	single.Invoice = invoice
	return single
}

func InsertPayment(conn *connections.Connections, req datastruct.PaymentRequest) error {
	var err error

	err = models.InsertPayment(conn, req)
	if err != nil {
		return err
	}

	return err
}

func UpdatePayment(conn *connections.Connections, req datastruct.PaymentRequest) error {
	var err error

	err = models.UpdatePayment(conn, req)
	if err != nil {
		return err
	}

	return err
}

func DeletePayment(conn *connections.Connections, req datastruct.PaymentRequest) error {
	err := models.DeletePayment(conn, req)
	return err
}
