package datastruct

import (
	"billingdashboard/core"
)

//LoginRequest is use for clients login
type ItemPriceRequest struct {
	ListItemPriceID []string              `json:"list_item-price_id"`
	ListItemPrice   []ItemPriceDataStruct `json:"list_item_price"`
	ItemID          string                `json:"item_id"`
	AccountID       string                `json:"account_id"`
	ServerID        string                `json:"server_id"`
	Price           string                `json:"price"`
	Category        string                `json:"category"`

	ListAccountID []string            `json:"list_account_id"`
	ListServerID  []string            `json:"list_server_id"`
	Param         core.DataTableParam `json:"param"`
}

type ItemPriceDataStruct struct {
	ItemID    string `json:"item_id"`
	AccountID string `json:"account_id"`
	ServerID  string `json:"server_id"`
	Price     string `json:"price"`
	Category  string `json:"category"`

	Account AccountDataStruct `json:"account"`
	Server  ServerDataStruct  `json:"server"`
	Item    ItemDataStruct    `json:"item"`
}

type AccountDataStruct struct {
	AccountID          string `json:"account_id"`
	Name               string `json:"name"`
	Status             string `json:"status"`
	CompanyID          string `json:"company_id"`
	AccountType        string `json:"account_type"`
	BillingType        string `json:"billing_type"`
	Desc               string `json:"desc"`
	Address1           string `json:"address1"`
	Address2           string `json:"address2"`
	City               string `json:"city"`
	Phone              string `json:"phone"`
	ContactPerson      string `json:"contact_person"`
	ContactPersonPhone string `json:"contact_person_phone"`
	LastUpdateUsername string `json:"last_update_username"`
	LastUpdateDate     string `json:"last_update_date"`
}

type ServerDataStruct struct {
	ServerID   string `json:"server_id"`
	ServerName string `json:"server_name"`
	ServerUrl  string `json:"server_url"`
}

type ItemDataStruct struct {
	ItemID   string `json:"item_id"`
	ItemName string `json:"item_name"`
	Operator string `json:"operator"`
	Route    string `json:"route"`
	Category string `json:"category"`
	UOM      string `json:"uom"`
}
