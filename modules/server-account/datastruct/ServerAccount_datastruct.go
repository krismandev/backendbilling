package datastruct

import (
	"billingdashboard/core"
)

type ServerAccountRequest struct {
	ServerID          string                    `json:"server_id"`
	AccountID         string                    `json:"account_id"`
	ServerAccount     string                    `json:"server_account"`
	ListAccountID     []string                  `json:"list_account_id"`
	ListServerAccount []ServerAccountDataStruct `json:"list_server_account"`
	Param             core.DataTableParam       `json:"param"`
}

type ServerAccountDataStruct struct {
	ServerID           string   `json:"server_id"`
	AccountID          string   `json:"account_id"`
	ServerAccount      string   `json:"server_account"`
	ListAccountID      []string `json:"list_account_id"`
	LastUpdateUsername string   `json:"last_update_username"`
	LastUpdateDate     string   `json:"last_update_date"`
}
