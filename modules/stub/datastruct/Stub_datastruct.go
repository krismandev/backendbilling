package datastruct

import (
	"backendbillingdashboard/core"
)

//LoginRequest is use for clients login
type StubRequest struct {
	ListStubID []string            `json:"list_stubid"`
	StubID     string              `json:"stubid"`
	StubName   string              `json:"stubname"`
	Param      core.DataTableParam `json:"param"`
}

type StubDataStruct struct {
	StubID   string `json:"stubid"`
	StubName string `json:"stubname"`
}
