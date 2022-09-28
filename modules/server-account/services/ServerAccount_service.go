package services

import (
	connections "backendbillingdashboard/connections"
	"backendbillingdashboard/core"
	dt "backendbillingdashboard/modules/server-account/datastruct"
	"backendbillingdashboard/modules/server-account/models"
	"backendbillingdashboard/modules/server-account/processors"
	"context"

	log "github.com/sirupsen/logrus"
)

// ServerAccountServices provides operations for endpoint

// ListServerAccount is use for
func ListServerAccount(ctx context.Context, req dt.ServerAccountRequest, conn *connections.Connections) core.GlobalListResponse {
	log.Infof("ServerAccountService.ListServerAccount Request : %+v", req)
	var response = core.DefaultGlobalListResponse(ctx)
	var err error

	listServerAccount, err := processors.GetListServerAccount(conn, req)
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

// CreateServerAccount is use for
func CreateServerAccount(ctx context.Context, req dt.ServerAccountRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("ServerAccountService.CreateServerAccount Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	// if len(req.AccountID) == 0 || len(req.ServerID) == 0 || len(req.ServerAccount) == 0 {
	// 	core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
	// 	return response
	// }

	// block request if data is already exists
	// err = models.CheckServerAccountDuplicate("", conn, req)
	// if err != nil {
	// 	core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, core.DescDataExists, err)
	// 	return response
	// }

	// process input
	response.Data, err = processors.InsertServerAccount(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrServer, err.Error(), err)
	}

	return response
}

// UpdateServerAccount is use for
func UpdateServerAccount(ctx context.Context, req dt.ServerAccountRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("ServerAccountService.UpdateServerAccount Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.AccountID) == 0 || len(req.ServerID) == 0 || len(req.ExternalAccountID) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// block request if old data is not exists
	err = models.CheckServerAccountExists(req.AccountID, conn)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrNoData, core.DescNoData, err)
		return response
	}

	// block request if data is already exists
	err = models.CheckServerAccountDuplicate(req.AccountID, conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, err.Error(), err)
		return response
	}

	// process input
	response.Data, err = processors.UpdateServerAccount(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrServer, core.DescServer, err)
	}

	return response
}

// DeleteServerAccount is use for
func DeleteServerAccount(ctx context.Context, req dt.ServerAccountRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("ServerAccountService.DeleteServerAccount Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.AccountID) == 0 || len(req.ServerID) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// run
	err = processors.DeleteServerAccount(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, err.Error(), err)
		return response
	}

	return response
}
