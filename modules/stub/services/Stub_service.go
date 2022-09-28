package services

import (
	connections "backendbillingdashboard/connections"
	"backendbillingdashboard/core"
	dt "backendbillingdashboard/modules/stub/datastruct"
	"backendbillingdashboard/modules/stub/models"
	"backendbillingdashboard/modules/stub/processors"
	"context"

	log "github.com/sirupsen/logrus"
)

// StubServices provides operations for endpoint

// ListStub is use for
func ListStub(ctx context.Context, req dt.StubRequest, conn *connections.Connections) core.GlobalListResponse {
	log.Infof("StubService.ListStub Request : %+v", req)
	var response = core.DefaultGlobalListResponse(ctx)
	var err error

	listStub, err := processors.GetListStub(conn, req)
	if err != nil {
		core.ErrorGlobalListResponse(&response, core.ErrServer, core.DescServer, err)
		return response
	} else {
		response.Data.Page = req.Param.Page
		response.Data.PerPage = req.Param.PerPage
	}

	// append list data as []interface{}
	for _, ls := range listStub {
		response.Data.List = append(response.Data.List, ls)
	}

	return response
}

// CreateStub is use for
func CreateStub(ctx context.Context, req dt.StubRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("StubService.CreateStub Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.StubID) == 0 || len(req.StubName) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// block request if data is already exists
	err = models.CheckStubDuplicate("", conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, core.DescDataExists, err)
		return response
	}

	// process input
	response.Data, err = processors.InsertStub(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrServer, core.DescServer, err)
	}

	return response
}

// UpdateStub is use for
func UpdateStub(ctx context.Context, req dt.StubRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("StubService.UpdateStub Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.StubID) == 0 || len(req.StubName) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// block request if old data is not exists
	err = models.CheckStubExists(req.StubID, conn)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrNoData, core.DescNoData, err)
		return response
	}

	// block request if data is already exists
	err = models.CheckStubDuplicate(req.StubID, conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, err.Error(), err)
		return response
	}

	// process input
	response.Data, err = processors.UpdateStub(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrServer, core.DescServer, err)
	}

	return response
}

// DeleteStub is use for
func DeleteStub(ctx context.Context, req dt.StubRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("StubService.DeleteStub Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.StubID) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// run
	err = processors.DeleteStub(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, err.Error(), err)
		return response
	}

	return response
}
