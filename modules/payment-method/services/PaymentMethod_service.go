package services

import (
	connections "backendbillingdashboard/connections"
	"backendbillingdashboard/core"
	dt "backendbillingdashboard/modules/payment-method/datastruct"
	"backendbillingdashboard/modules/payment-method/models"
	"backendbillingdashboard/modules/payment-method/processors"
	"context"

	log "github.com/sirupsen/logrus"
)

// PaymentMethodServices provides operations for endpoint

// ListPaymentMethod is use for
func ListPaymentMethod(ctx context.Context, req dt.PaymentMethodRequest, conn *connections.Connections) core.GlobalListResponse {
	log.Infof("PaymentMethodService.ListPaymentMethod Request : %+v", req)
	var response = core.DefaultGlobalListResponse(ctx)
	var err error

	listPaymentMethod, err := processors.GetListPaymentMethod(conn, req)
	if err != nil {
		core.ErrorGlobalListResponse(&response, core.ErrServer, core.DescServer, err)
		return response
	} else {
		response.Data.Page = req.Param.Page
		response.Data.PerPage = req.Param.PerPage
	}

	// append list data as []interface{}
	for _, ls := range listPaymentMethod {
		response.Data.List = append(response.Data.List, ls)
	}

	return response
}

// CreatePaymentMethod is use for
func CreatePaymentMethod(ctx context.Context, req dt.PaymentMethodRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("PaymentMethodService.CreatePaymentMethod Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.Key) == 0 || len(req.PaymentMethodName) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// block request if data is already exists
	err = models.CheckPaymentMethodDuplicate("", conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, core.DescDataExists, err)
		return response
	}

	// process input
	response.Data, err = processors.InsertPaymentMethod(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrServer, core.DescServer, err)
	}

	return response
}

// UpdatePaymentMethod is use for
func UpdatePaymentMethod(ctx context.Context, req dt.PaymentMethodRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("PaymentMethodService.UpdatePaymentMethod Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.Key) == 0 || len(req.PaymentMethodName) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// block request if old data is not exists
	err = models.CheckPaymentMethodExists(req.Key, conn)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrNoData, core.DescNoData, err)
		return response
	}

	// block request if data is already exists
	err = models.CheckPaymentMethodDuplicate(req.Key, conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, err.Error(), err)
		return response
	}

	// process input
	response.Data, err = processors.UpdatePaymentMethod(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrServer, core.DescServer, err)
	}

	return response
}

// DeletePaymentMethod is use for
func DeletePaymentMethod(ctx context.Context, req dt.PaymentMethodRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("PaymentMethodService.DeletePaymentMethod Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.Key) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// run
	err = processors.DeletePaymentMethod(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, err.Error(), err)
		return response
	}

	return response
}
