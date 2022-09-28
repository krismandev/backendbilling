package services

import (
	connections "backendbillingdashboard/connections"
	"backendbillingdashboard/core"
	dt "backendbillingdashboard/modules/currency/datastruct"
	"backendbillingdashboard/modules/currency/models"
	"backendbillingdashboard/modules/currency/processors"
	"context"

	log "github.com/sirupsen/logrus"
)

// CurrencyServices provides operations for endpoint

// ListCurrency is use for
func ListCurrency(ctx context.Context, req dt.CurrencyRequest, conn *connections.Connections) core.GlobalListResponse {
	log.Infof("CurrencyService.ListCurrency Request : %+v", req)
	var response = core.DefaultGlobalListResponse(ctx)
	var err error

	listCurrency, err := processors.GetListCurrency(conn, req)
	if err != nil {
		core.ErrorGlobalListResponse(&response, core.ErrServer, core.DescServer, err)
		return response
	} else {
		response.Data.Page = req.Param.Page
		response.Data.PerPage = req.Param.PerPage
	}

	// append list data as []interface{}
	for _, ls := range listCurrency {
		response.Data.List = append(response.Data.List, ls)
	}

	return response
}

// CreateCurrency is use for
func CreateCurrency(ctx context.Context, req dt.CurrencyRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("CurrencyService.CreateCurrency Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.CurrencyName) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// block request if data is already exists
	err = models.CheckCurrencyDuplicate("", conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, core.DescDataExists, err)
		return response
	}

	// process input
	response.Data, err = processors.InsertCurrency(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrServer, core.DescServer, err)
	}

	return response
}

// UpdateCurrency is use for
func UpdateCurrency(ctx context.Context, req dt.CurrencyRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("CurrencyService.UpdateCurrency Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.CurrencyName) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// block request if old data is not exists
	err = models.CheckCurrencyExists(req.CurrencyCode, conn)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrNoData, core.DescNoData, err)
		return response
	}

	// block request if data is already exists
	err = models.CheckCurrencyDuplicate(req.CurrencyCode, conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, err.Error(), err)
		return response
	}

	// process input
	response.Data, err = processors.UpdateCurrency(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrServer, core.DescServer, err)
	}

	return response
}

// DeleteCurrency is use for
func DeleteCurrency(ctx context.Context, req dt.CurrencyRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("CurrencyService.DeleteCurrency Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.CurrencyCode) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// run
	err = processors.DeleteCurrency(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, err.Error(), err)
		return response
	}

	return response
}

// ListBalance is use for
func ListBalance(ctx context.Context, req dt.BalanceRequest, conn *connections.Connections) core.GlobalListResponse {
	log.Infof("BalanceService.ListBalance Request : %+v", req)
	var response = core.DefaultGlobalListResponse(ctx)
	var err error

	listBalance, err := processors.GetListBalance(conn, req)
	if err != nil {
		core.ErrorGlobalListResponse(&response, core.ErrServer, core.DescServer, err)
		return response
	} else {
		response.Data.Page = req.Param.Page
		response.Data.PerPage = req.Param.PerPage
	}

	// append list data as []interface{}
	for _, ls := range listBalance {
		response.Data.List = append(response.Data.List, ls)
	}

	return response
}
