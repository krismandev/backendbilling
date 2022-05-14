package datastruct

import (
	"billingdashboard/core"
)

//LoginRequest is use for clients login
type PaymentMethodRequest struct {
	ListKey           []string            `json:"list_key"`
	Key               string              `json:"key"`
	PaymentMethodName string              `json:"payment_method_name"`
	NeedClearingDate  string              `json:"need_clearing_date"`
	NeedCardNumber    string              `json:"need_card_number"`
	Param             core.DataTableParam `json:"param"`
}

type PaymentMethodDataStruct struct {
	Key               string `json:"key"`
	PaymentMethodName string `json:"payment_method_name"`
	NeedClearingDate  string `json:"need_clearing_date"`
	NeedCardNumber    string `json:"need_card_number"`
}
