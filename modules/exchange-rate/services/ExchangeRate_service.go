package services

import (
	connections "backendbillingdashboard/connections"
	"backendbillingdashboard/core"
	dt "backendbillingdashboard/modules/exchange-rate/datastruct"
	"backendbillingdashboard/modules/exchange-rate/processors"
	"context"

	log "github.com/sirupsen/logrus"
)

// ExchangeRateServices provides operations for endpoint

// ListExchangeRate is use for
func ListExchangeRate(ctx context.Context, req dt.ExchangeRateRequest, conn *connections.Connections) core.GlobalListResponse {
	log.Infof("ExchangeRateService.ListExchangeRate Request : %+v", req)
	var response = core.DefaultGlobalListResponse(ctx)
	var err error

	listExchangeRate, err := processors.GetListExchangeRate(conn, req)
	if err != nil {
		core.ErrorGlobalListResponse(&response, core.ErrServer, core.DescServer, err)
		return response
	} else {
		response.Data.Page = req.Param.Page
		response.Data.PerPage = req.Param.PerPage
	}

	// append list data as []interface{}
	for _, ls := range listExchangeRate {
		response.Data.List = append(response.Data.List, ls)
	}

	return response
}

// CreateExchangeRate is use for
func CreateExchangeRate(ctx context.Context, req dt.ExchangeRateRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("ExchangeRateService.CreateExchangeRate Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.Date) == 0 || len(req.FromCurrency) == 0 || len(req.ToCurrency) == 0 || len(req.ConvertValue) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// // block request if data is already exists
	// err = models.CheckExchangeRateDuplicate("", conn, req)
	// if err != nil {
	// 	core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, core.DescDataExists, err)
	// 	return response
	// }

	// process input
	err = processors.InsertExchangeRate(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrServer, err.Error(), err)
	}

	return response
}

// UpdateExchangeRate is use for
func UpdateExchangeRate(ctx context.Context, req dt.ExchangeRateRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("ExchangeRateService.UpdateExchangeRate Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.Date) == 0 || len(req.FromCurrency) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// // block request if old data is not exists
	// err = models.CheckExchangeRateExists(req.ExchangeRateID, conn)
	// if err != nil {
	// 	core.ErrorGlobalSingleResponse(&response, core.ErrNoData, core.DescNoData, err)
	// 	return response
	// }

	// // block request if data is already exists
	// err = models.CheckExchangeRateDuplicate(req.ExchangeRateID, conn, req)
	// if err != nil {
	// 	core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, err.Error(), err)
	// 	return response
	// }

	// process input
	err = processors.UpdateExchangeRate(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrServer, core.DescServer, err)
	}

	return response
}

// DeleteExchangeRate is use for
func DeleteExchangeRate(ctx context.Context, req dt.ExchangeRateRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("ExchangeRateService.DeleteExchangeRate Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.Date) == 0 || len(req.FromCurrency) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// run
	err = processors.DeleteExchangeRate(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, err.Error(), err)
		return response
	}

	return response
}
