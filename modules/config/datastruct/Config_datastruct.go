package datastruct

import (
	"backendbillingdashboard/core"
)

//LoginRequest is use for clients login
type ConfigRequest struct {
	ListConfig []string            `json:"list_config"`
	Key        string              `json:"key"`
	Type       string              `json:"type"`
	Value      string              `json:"value"`
	Param      core.DataTableParam `json:"param"`
}

type ConfigDataStruct struct {
	Key   string `json:"key"`
	Type  string `json:"type"`
	Value string `json:"value"`
}
