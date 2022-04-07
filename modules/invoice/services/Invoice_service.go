package services

import (
	connections "billingdashboard/connections"
	"billingdashboard/core"
	dt "billingdashboard/modules/invoice/datastruct"
	"billingdashboard/modules/invoice/processors"
	"context"

	log "github.com/sirupsen/logrus"
)

// InvoiceServices provides operations for endpoint

// ListInvoice is use for
func ListInvoice(ctx context.Context, req dt.InvoiceRequest, conn *connections.Connections) core.GlobalListResponse {
	log.Infof("InvoiceService.ListInvoice Request : %+v", req)
	var response = core.DefaultGlobalListResponse(ctx)
	var err error

	listInvoice, err := processors.GetListInvoice(conn, req)
	if err != nil {
		core.ErrorGlobalListResponse(&response, core.ErrServer, core.DescServer, err)
		return response
	} else {
		response.Data.Page = req.Param.Page
		response.Data.PerPage = req.Param.PerPage
	}

	// append list data as []interface{}
	for _, ls := range listInvoice {
		response.Data.List = append(response.Data.List, ls)
	}

	return response
}

// CreateInvoice is use for
func CreateInvoice(ctx context.Context, req dt.InvoiceRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("InvoiceService.CreateInvoice Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.AccountID) == 0 || len(req.InvoiceTypeID) == 0 || len(req.MonthUse) == 0 || len(req.TransDate) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// block request if data is already exists
	// err = models.CheckInvoiceDuplicate("", conn, req)
	// if err != nil {
	// 	core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, core.DescDataExists, err)
	// 	return response
	// }

	// process input
	err = processors.InsertInvoice(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrServer, core.DescServer, err)
	}

	return response
}

// UpdateInvoice is use for
func UpdateInvoice(ctx context.Context, req dt.InvoiceRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("InvoiceService.UpdateInvoice Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.InvoiceID) == 0 || len(req.MonthUse) == 0 || len(req.TransDate) == 0 || len(req.InvoiceNo) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// block request if old data is not exists
	// err = models.CheckInvoiceExists(req.InvoiceID, conn)
	// if err != nil {
	// 	core.ErrorGlobalSingleResponse(&response, core.ErrNoData, core.DescNoData, err)
	// 	return response
	// }

	// block request if data is already exists
	// err = models.CheckInvoiceDuplicate(req.InvoiceID, conn, req)
	// if err != nil {
	// 	core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, err.Error(), err)
	// 	return response
	// }

	// process input
	err = processors.UpdateInvoice(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrServer, err.Error(), err)
	}

	return response
}

// DeleteInvoice is use for
func DeleteInvoice(ctx context.Context, req dt.InvoiceRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("InvoiceService.DeleteInvoice Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.InvoiceID) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// run
	err = processors.DeleteInvoice(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, err.Error(), err)
		return response
	}

	return response
}

func CancelInvoice(ctx context.Context, req dt.InvoiceRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("InvoiceService.CancelInvoice Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.InvoiceID) == 0 || len(req.CancelDesc) < 10 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// run
	err = processors.CancelInvoice(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, err.Error(), err)
		return response
	}

	return response
}

func PrintInvoice(ctx context.Context, req dt.InvoiceRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("InvoiceService.PrintInvoice Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.InvoiceID) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// block request if old data is not exists
	// err = models.CheckInvoiceExists(req.InvoiceID, conn)
	// if err != nil {
	// 	core.ErrorGlobalSingleResponse(&response, core.ErrNoData, core.DescNoData, err)
	// 	return response
	// }

	// block request if data is already exists
	// err = models.CheckInvoiceDuplicate(req.InvoiceID, conn, req)
	// if err != nil {
	// 	core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, err.Error(), err)
	// 	return response
	// }

	// process
	err = processors.PrintInvoice(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrServer, core.DescServer, err)
	}

	return response
}
