package processors

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/payment-method/datastruct"
	"backendbillingdashboard/modules/payment-method/models"
)

func GetListPaymentMethod(conn *connections.Connections, req datastruct.PaymentMethodRequest) ([]datastruct.PaymentMethodDataStruct, error) {
	var output []datastruct.PaymentMethodDataStruct
	var err error

	// grab mapping data from model
	paymentMethodList, err := models.GetPaymentMethodFromRequest(conn, req)
	if err != nil {
		return output, err
	}

	for _, stub := range paymentMethodList {
		single := CreateSinglePaymentMethodStruct(stub)
		output = append(output, single)
	}

	return output, err
}

func CreateSinglePaymentMethodStruct(stub map[string]string) datastruct.PaymentMethodDataStruct {
	var single datastruct.PaymentMethodDataStruct
	single.Key, _ = stub["key"]
	single.PaymentMethodName, _ = stub["payment_method_name"]
	single.NeedClearingDate, _ = stub["need_clearing_date"]
	single.NeedCardNumber, _ = stub["need_card_number"]
	single.BankName, _ = stub["bank_name"]
	single.Branch, _ = stub["branch"]
	single.AccountName, _ = stub["account_name"]
	single.AccountNo, _ = stub["account_no"]
	single.Code, _ = stub["code"]
	single.Status, _ = stub["status"]
	single.PayementType, _ = stub["payment_type"]
	single.CurrencyCode, _ = stub["currency_code"]

	return single
}

func InsertPaymentMethod(conn *connections.Connections, req datastruct.PaymentMethodRequest) (datastruct.PaymentMethodDataStruct, error) {
	var output datastruct.PaymentMethodDataStruct
	var err error

	err = models.InsertPaymentMethod(conn, req)
	if err != nil {
		return output, err
	}

	// jika tidak ada error, return single instance of single stub
	single, err := models.GetSinglePaymentMethod(req.Key, conn, req)
	if err != nil {
		return output, err
	}

	output = CreateSinglePaymentMethodStruct(single)
	return output, err
}

func UpdatePaymentMethod(conn *connections.Connections, req datastruct.PaymentMethodRequest) (datastruct.PaymentMethodDataStruct, error) {
	var output datastruct.PaymentMethodDataStruct
	var err error

	err = models.UpdatePaymentMethod(conn, req)
	if err != nil {
		return output, err
	}

	// jika tidak ada error, return single instance of single stub
	single, err := models.GetSinglePaymentMethod(req.Key, conn, req)
	if err != nil {
		return output, err
	}

	output = CreateSinglePaymentMethodStruct(single)
	return output, err
}

func DeletePaymentMethod(conn *connections.Connections, req datastruct.PaymentMethodRequest) error {
	err := models.DeletePaymentMethod(conn, req)
	return err
}
