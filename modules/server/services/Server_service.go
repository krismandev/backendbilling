package services

import (
	connections "backendbillingdashboard/connections"
	"backendbillingdashboard/core"
	dt "backendbillingdashboard/modules/server/datastruct"
	"backendbillingdashboard/modules/server/models"
	"backendbillingdashboard/modules/server/processors"
	"context"

	log "github.com/sirupsen/logrus"
)

// ServerServices provides operations for endpoint

// ListServer is use for
func ListServer(ctx context.Context, req dt.ServerRequest, conn *connections.Connections) core.GlobalListResponse {
	log.Infof("ServerService.ListServer Request : %+v", req)
	var response = core.DefaultGlobalListResponse(ctx)
	var err error

	listServer, err := processors.GetListServer(conn, req)
	if err != nil {
		core.ErrorGlobalListResponse(&response, core.ErrServer, core.DescServer, err)
		return response
	} else {
		response.Data.Page = req.Param.Page
		response.Data.PerPage = req.Param.PerPage
	}

	// append list data as []interface{}
	for _, ls := range listServer {
		response.Data.List = append(response.Data.List, ls)
	}

	return response
}

// CreateServer is use for
func CreateServer(ctx context.Context, req dt.ServerRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("ServerService.CreateServer Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.ServerName) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// block request if data is already exists
	err = models.CheckServerDuplicate("", conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, core.DescDataExists, err)
		return response
	}

	// process input
	err = processors.InsertServer(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrServer, core.DescServer, err)
	}

	return response
}

// UpdateServer is use for
func UpdateServer(ctx context.Context, req dt.ServerRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("ServerService.UpdateServer Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.ServerID) == 0 || len(req.ServerName) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// block request if old data is not exists
	err = models.CheckServerExists(req.ServerID, conn)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrNoData, core.DescNoData, err)
		return response
	}

	// process input
	err = processors.UpdateServer(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrServer, core.DescServer, err)
	}

	return response
}

// DeleteServer is use for
func DeleteServer(ctx context.Context, req dt.ServerRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("ServerService.DeleteServer Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.ServerID) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// run
	err = processors.DeleteServer(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, err.Error(), err)
		return response
	}

	return response
}

func ListServerAccount(ctx context.Context, req dt.ServerAccountRequest, conn *connections.Connections) core.GlobalListResponse {
	log.Infof("ServerService.ListServerAccount Request : %+v", req)
	var response = core.DefaultGlobalListResponse(ctx)
	var err error

	listServerAccount, err := processors.GetServerAccount(conn, req)
	if err != nil {
		core.ErrorGlobalListResponse(&response, core.ErrServer, core.DescServer, err)
		return response
	} else {
		response.Data.Page = req.Param.Page
		response.Data.PerPage = req.Param.PerPage
	}

	// append list data as []interface{}
	for _, ls := range listServerAccount {
		response.Data.List = append(response.Data.List, ls)
	}

	return response
}
