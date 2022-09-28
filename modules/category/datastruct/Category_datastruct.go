package datastruct

import (
	"backendbillingdashboard/core"
)

//LoginRequest is use for clients login
type CategoryRequest struct {
	ListCategoryID []string            `json:"list_categoryid"`
	CategoryID     string              `json:"categoryid"`
	CategoryName   string              `json:"category_name"`
	Param          core.DataTableParam `json:"param"`
}

type CategoryDataStruct struct {
	CategoryID   string `json:"category_id"`
	CategoryName string `json:"category_name"`
}
