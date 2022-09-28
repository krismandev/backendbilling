package core

import (
	"backendbillingdashboard/lib"
	"context"

	logger "github.com/sirupsen/logrus"
)

type contextKey string

var (
	//ContextTransactionID for context transaction ID
	ContextTransactionID = contextKey("IMSTransactionID")
)

// GlobalJSONResponse is used for
type GlobalJSONResponse struct {
	TransID      string `json:"trans_id"`
	ResponseCode int    `json:"response_code"`
	ResponseDesc string `json:"response_desc"`
	ErrorDetail  string `json:"error_detail"`
	LastInsertID string `json:"last_insert_id,omitempty"`
}

// GlobalListResponse use as global listing response format
type GlobalListResponse struct {
	TransID      string                      `json:"trans_id"`
	ResponseCode int                         `json:"response_code"`
	ResponseDesc string                      `json:"response_desc"`
	ErrorDetail  string                      `json:"error_detail"`
	Data         GlobalListDataTableResponse `json:"data,omitempty"`
}

type GlobalListDataTableResponse struct {
	// TotalData int64 `json:"total_data"`
	PerPage int `json:"per_page"`
	Page    int `json:"page"`
	// TotalPage int           `json:"total_page"`
	List []interface{} `json:"list"`
}

// GlobalSingleResponse used as global single data response
type GlobalSingleResponse struct {
	TransID      string      `json:"trans_id"`
	ResponseCode int         `json:"response_code"`
	ResponseDesc string      `json:"response_desc"`
	ErrorDetail  string      `json:"error_detail"`
	Data         interface{} `json:"data,omitempty"`
}

// DataTableParam is used for
type DataTableParam struct {
	LastID   string `json:"last_id"`
	PerPage  int    `json:"per_page"`
	Page     int    `json:"page"`
	OrderBy  string `json:"order_by"`
	OrderDir string `json:"order_dir"`
}

//GetLogFieldValues is  use for Set Log Fields
func GetLogFieldValues(ctx context.Context, Module string) logger.Fields {
	var fields logger.Fields = make(logger.Fields)
	fields["Node"] = lib.GetHostName()
	fields["Module"] = Module
	if ctx != nil {
		fields["TxID"] = ctx.Value(ContextTransactionID).(string)
	}
	return fields
}

//ReloadConfigJSONRequest is use for
type ReloadConfigJSONRequest struct {
	OriginalRequest string
	IPAddr          string
}

//ReloadConfigJSONResponse is  use for
type ReloadConfigJSONResponse struct {
	ResponseCode int    `json:"responseCode"`
	ResponseDesc string `json:"responseDesc"`
	IPAddr       string `json:"-"`
}

//ReloadConfigType is  use for
type ReloadConfigType struct {
	Event        string `json:"event,omitempty"`
	DeviceNumber string `json:"device_number,omitempty"`
	Timestamp    string `json:"timestamp,omitempty"`
}
