package services

import (
	connections "backendbillingdashboard/connections"
	"backendbillingdashboard/core"
	dt "backendbillingdashboard/modules/invoice-type/datastruct"
	"backendbillingdashboard/modules/invoice-type/models"
	"backendbillingdashboard/modules/invoice-type/processors"
	"context"

	log "github.com/sirupsen/logrus"
)

// InvoiceTypeServices provides operations for endpoint

// ListInvoiceType is use for
func ListInvoiceType(ctx context.Context, req dt.InvoiceTypeRequest, conn *connections.Connections) core.GlobalListResponse {
	log.Infof("InvoiceTypeService.ListInvoiceType Request : %+v", req)
	var response = core.DefaultGlobalListResponse(ctx)
	var err error

	listInvoiceType, err := processors.GetListInvoiceType(conn, req)
	if err != nil {
		core.ErrorGlobalListResponse(&response, core.ErrServer, core.DescServer, err)
		return response
	} else {
		response.Data.Page = req.Param.Page
		response.Data.PerPage = req.Param.PerPage
	}

	// append list data as []interface{}
	for _, ls := range listInvoiceType {
		response.Data.List = append(response.Data.List, ls)
	}

	return response
}

// CreateInvoiceType is use for
func CreateInvoiceType(ctx context.Context, req dt.InvoiceTypeRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("InvoiceTypeService.CreateInvoiceType Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.InvoiceTypeID) == 0 || len(req.InvoiceTypeName) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// block request if data is already exists
	err = models.CheckInvoiceTypeDuplicate("", conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, core.DescDataExists, err)
		return response
	}

	// process input
	response.Data, err = processors.InsertInvoiceType(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrServer, core.DescServer, err)
	}

	return response
}

// UpdateInvoiceType is use for
func UpdateInvoiceType(ctx context.Context, req dt.InvoiceTypeRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("InvoiceTypeService.UpdateInvoiceType Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.InvoiceTypeID) == 0 || len(req.InvoiceTypeName) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// block request if old data is not exists
	err = models.CheckInvoiceTypeExists(req.InvoiceTypeID, conn)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrNoData, core.DescNoData, err)
		return response
	}

	// block request if data is already exists
	err = models.CheckInvoiceTypeDuplicate(req.InvoiceTypeID, conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, err.Error(), err)
		return response
	}

	// process input
	response.Data, err = processors.UpdateInvoiceType(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrServer, core.DescServer, err)
	}

	return response
}

// DeleteInvoiceType is use for
func DeleteInvoiceType(ctx context.Context, req dt.InvoiceTypeRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("InvoiceTypeService.DeleteInvoiceType Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.InvoiceTypeID) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// run
	err = processors.DeleteInvoiceType(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, err.Error(), err)
		return response
	}

	return response
}
