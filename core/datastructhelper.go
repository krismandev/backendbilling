package core

import (
	"context"
)

// ResponseHandlerHelper is used for
func ResponseHandlerHelper(ctx context.Context, responseCode int, responseDesc string, err error) GlobalJSONResponse {
	transid := ctx.Value(ContextTransactionID).(string)
	errorDetail := ""
	if err != nil {
		errorDetail = err.Error()
	}
	return GlobalJSONResponse{
		TransID:      transid,
		ResponseCode: responseCode,
		ResponseDesc: responseDesc,
		ErrorDetail:  errorDetail,
	}
}

// DefaultGlobalListResponse will return default list response struct (in success mode)
func DefaultGlobalListResponse(ctx context.Context) GlobalListResponse {
	var resp GlobalListResponse
	resp.TransID = ctx.Value(ContextTransactionID).(string)
	resp.ResponseCode = ErrSuccess
	resp.ResponseDesc = DescSuccess

	return resp
}

// ErrorGlobalListResponse will return default list response struct (in error mode)
func ErrorGlobalListResponse(resp *GlobalListResponse, errCode int, errDesc string, err error) {
	resp.ResponseCode = errCode
	resp.ResponseDesc = errDesc
	if err != nil {
		resp.ErrorDetail = err.Error()
	} else {
		resp.ErrorDetail = ""
	}
}

// DefaultGlobalSingleResponse will return default single response struct (in success mode)
func DefaultGlobalSingleResponse(ctx context.Context) GlobalSingleResponse {
	var resp GlobalSingleResponse
	resp.TransID = ctx.Value(ContextTransactionID).(string)
	resp.ResponseCode = ErrSuccess
	resp.ResponseDesc = DescSuccess

	return resp
}

// ErrorGlobalSingleResponse will return default single response struct (in error mode)
func ErrorGlobalSingleResponse(resp *GlobalSingleResponse, errCode int, errDesc string, err error) {
	resp.ResponseCode = errCode
	resp.ResponseDesc = errDesc
	if err != nil {
		resp.ErrorDetail = err.Error()
	} else {
		resp.ErrorDetail = ""
	}
}
