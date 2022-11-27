package datastruct

import (
	"backendbillingdashboard/core"
)

//LoginRequest is use for clients login
type ItemRequest struct {
	ListItemID []string            `json:"list_item_id"`
	ListItem   []ItemDataStruct    `json:"list_item"`
	ItemID     string              `json:"item_id"`
	ItemName   string              `json:"item_name"`
	Operator   string              `json:"operator"`
	Route      string              `json:"route"`
	Category   string              `json:"category"`
	UOM        string              `json:"uom"`
	Param      core.DataTableParam `json:"param"`
}

type ItemDataStruct struct {
	ItemID   string `json:"item_id"`
	ItemName string `json:"item_name"`
	Operator string `json:"operator"`
	Route    string `json:"route"`
	Category string `json:"category"`
	UOM      string `json:"uom"`
}

type Category struct {
	CategoryId         string `json:"category_id"`
	CategoryName       string `json:"name"`
	LastUpdateUsername string `json:"last_update_username"`
	LastUdpateDate     string `json:"last_update_date"`
}
