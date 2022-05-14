package services

import (
	connections "billingdashboard/connections"
	"billingdashboard/core"
	dt "billingdashboard/modules/payment/datastruct"
	"billingdashboard/modules/payment/models"
	"billingdashboard/modules/payment/processors"
	"context"

	log "github.com/sirupsen/logrus"
)

// PaymentServices provides operations for endpoint

// ListPayment is use for
func ListPayment(ctx context.Context, req dt.PaymentRequest, conn *connections.Connections) core.GlobalListResponse {
	log.Infof("PaymentService.ListPayment Request : %+v", req)
	var response = core.DefaultGlobalListResponse(ctx)
	var err error

	listPayment, err := processors.GetListPayment(conn, req)
	if err != nil {
		core.ErrorGlobalListResponse(&response, core.ErrServer, core.DescServer, err)
		return response
	} else {
		response.Data.Page = req.Param.Page
		response.Data.PerPage = req.Param.PerPage
	}

	// append list data as []interface{}
	for _, ls := range listPayment {
		response.Data.List = append(response.Data.List, ls)
	}

	return response
}

// CreatePayment is use for
func CreatePayment(ctx context.Context, req dt.PaymentRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("PaymentService.CreatePayment Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.InvoiceID) == 0 || len(req.Total) == 0 || len(req.LastUpdateUsername) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// block request if data is already exists
	err = models.CheckIsItPaid(req.InvoiceID, conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, "This invoice already paid", err)
		return response
	}

	errCheckNominal := models.CheckPaymentNominal(conn, req)
	if errCheckNominal != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, errCheckNominal.Error(), errCheckNominal)
		return response
	}
	// process input
	err = processors.InsertPayment(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrServer, core.DescServer, err)
	}

	return response
}

// UpdatePayment is use for
func UpdatePayment(ctx context.Context, req dt.PaymentRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("PaymentService.UpdatePayment Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.PaymentID) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// block request if old data is not exists
	err = models.CheckPaymentExists(req.PaymentID, conn)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrNoData, core.DescNoData, err)
		return response
	}

	// block request if data is already exists
	err = models.CheckPaymentDuplicate(req.PaymentID, conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, err.Error(), err)
		return response
	}

	// process input
	err = processors.UpdatePayment(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrServer, core.DescServer, err)
	}

	return response
}

// DeletePayment is use for
func DeletePayment(ctx context.Context, req dt.PaymentRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("PaymentService.DeletePayment Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.PaymentID) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// run
	err = processors.DeletePayment(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, err.Error(), err)
		return response
	}

	return response
}
