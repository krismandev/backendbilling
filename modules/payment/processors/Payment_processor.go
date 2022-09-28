package processors

import (
	"backendbillingdashboard/connections"
	paymentMethodDt "backendbillingdashboard/modules/payment-method/datastruct"
	"backendbillingdashboard/modules/payment/datastruct"
	"backendbillingdashboard/modules/payment/models"
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
	single.LastUpdateUsername, _ = payment["last_update_username"].(string)
	single.PaymentMethod, _ = payment["payment_method"].(string)
	single.CardNumber, _ = payment["card_number"].(string)
	single.ClearingDate, _ = payment["clearing_date"].(string)
	single.Status, _ = payment["status"].(string)
	single.PaymentType, _ = payment["payment_type"].(string)

	var account datastruct.AccountDataStruct
	account.Name = payment["invoice"].(map[string]interface{})["account"].(map[string]interface{})["name"].(string)

	var invoice datastruct.InvoiceDataStruct
	invoice.InvoiceNo = payment["invoice"].(map[string]interface{})["invoice_no"].(string)
	invoice.Account = account
	single.Invoice = invoice

	var paymentMethod paymentMethodDt.PaymentMethodDataStruct
	paymentMethod.Key = payment["payment_method_object"].(map[string]interface{})["key"].(string)
	paymentMethod.PaymentMethodName = payment["payment_method_object"].(map[string]interface{})["payment_method_name"].(string)
	paymentMethod.BankName = payment["payment_method_object"].(map[string]interface{})["bank_name"].(string)
	paymentMethod.PayementType = payment["payment_method_object"].(map[string]interface{})["payment_type"].(string)
	paymentMethod.CurrencyCode = payment["payment_method_object"].(map[string]interface{})["currency_code"].(string)
	single.PaymentMethodObject = paymentMethod

	var paymentDeductions []datastruct.PaymentDeductionDataStruct
	for _, each := range payment["payment_deductions"].([]map[string]string) {
		var paymentDeduction datastruct.PaymentDeductionDataStruct
		paymentDeduction.PaymentDeductionTypeID = each["payment_deduction_type_id"]
		paymentDeduction.PaymentID = each["payment_id"]
		paymentDeduction.Amount = each["amount"]
		paymentDeduction.Description = each["description"]
		paymentDeductions = append(paymentDeductions, paymentDeduction)
	}
	single.PaymentDeduction = paymentDeductions
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

func GetListPaymentDeductionType(conn *connections.Connections, req datastruct.PaymentDeductionTypeRequest) ([]datastruct.PaymentDeductionTypeDataStruct, error) {
	var output []datastruct.PaymentDeductionTypeDataStruct
	var err error

	// grab mapping data from model
	paymentDeductionTypeList, err := models.GetPaymentDeductionTypeFromRequest(conn, req)
	if err != nil {
		return output, err
	}

	for _, paymentDeductionType := range paymentDeductionTypeList {
		single := CreateSinglePaymentDeductionTypeStruct(paymentDeductionType)
		output = append(output, single)
	}

	return output, err
}

func CreateSinglePaymentDeductionTypeStruct(paymentDeductionType map[string]string) datastruct.PaymentDeductionTypeDataStruct {
	var single datastruct.PaymentDeductionTypeDataStruct
	single.PaymentDeductionTypeID, _ = paymentDeductionType["payment_deduction_type_id"]
	single.Description, _ = paymentDeductionType["description"]
	single.Category, _ = paymentDeductionType["category"]
	single.Amount, _ = paymentDeductionType["amount"]
	single.LastUpdateUsername, _ = paymentDeductionType["last_update_username"]
	single.LastUpdateDate, _ = paymentDeductionType["last_update_date"]
	return single
}

func GetListPaymentDeduction(conn *connections.Connections, req datastruct.PaymentDeductionRequest) ([]datastruct.PaymentDeductionDataStruct, error) {
	var output []datastruct.PaymentDeductionDataStruct
	var err error

	// grab mapping data from model
	paymentDeductionList, err := models.GetPaymentDeductionFromRequest(conn, req)
	if err != nil {
		return output, err
	}

	for _, paymentDeduction := range paymentDeductionList {
		single := CreateSinglePaymentDeductionStruct(paymentDeduction)
		output = append(output, single)
	}

	return output, err
}

func CreateSinglePaymentDeductionStruct(paymentDeduction map[string]string) datastruct.PaymentDeductionDataStruct {
	var single datastruct.PaymentDeductionDataStruct
	single.PaymentDeductionTypeID, _ = paymentDeduction["payment_deduction_type_id"]
	single.PaymentID, _ = paymentDeduction["payment_id"]
	single.InvoiceID, _ = paymentDeduction["invoice_id"]
	single.Amount, _ = paymentDeduction["amount"]
	single.Description, _ = paymentDeduction["description"]
	single.UniqueInInvoice, _ = paymentDeduction["unique_in_invoice"]
	single.Status, _ = paymentDeduction["status"]
	return single
}

func GetListAdjustmentReason(conn *connections.Connections, req datastruct.AdjustmentReasonRequest) ([]datastruct.AdjustmentReasonDataStruct, error) {
	var output []datastruct.AdjustmentReasonDataStruct
	var err error

	// grab mapping data from model
	adjustmentReasonList, err := models.GetAdjustmentReasonFromRequest(conn, req)
	if err != nil {
		return output, err
	}

	for _, adjustmentReason := range adjustmentReasonList {
		single := CreateSingleAdjustmentReasonStruct(adjustmentReason)
		output = append(output, single)
	}

	return output, err
}

func CreateSingleAdjustmentReasonStruct(adjustmentReason map[string]string) datastruct.AdjustmentReasonDataStruct {
	var single datastruct.AdjustmentReasonDataStruct
	single.Key, _ = adjustmentReason["key"]
	single.Description, _ = adjustmentReason["description"]
	single.PaymentMethod, _ = adjustmentReason["payment_method"]
	single.CurrencyCode, _ = adjustmentReason["currency_code"]
	return single
}
