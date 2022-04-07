package services

import (
	connections "billingdashboard/connections"
	"billingdashboard/core"
	dt "billingdashboard/modules/account/datastruct"
	"billingdashboard/modules/account/models"
	"billingdashboard/modules/account/processors"
	"context"

	log "github.com/sirupsen/logrus"
)

// StubServices provides operations for endpoint

// ListStub is use for
func ListAccount(ctx context.Context, req dt.AccountRequest, conn *connections.Connections) core.GlobalListResponse {
	log.Infof("CompanyService.ListAccount Request : %+v", req)
	var response = core.DefaultGlobalListResponse(ctx)
	var err error

	listAccount, err := processors.GetListAccount(conn, req)
	if err != nil {
		core.ErrorGlobalListResponse(&response, core.ErrServer, core.DescServer, err)
		return response
	} else {
		response.Data.Page = req.Param.Page
		response.Data.PerPage = req.Param.PerPage
	}

	// append list data as []interface{}
	for _, ls := range listAccount {
		response.Data.List = append(response.Data.List, ls)
	}

	return response
}

// CreateCompany is use for
func CreateAccount(ctx context.Context, req dt.AccountRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("AccountService.CreateAccount Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.Name) == 0 || len(req.Status) == 0 || len(req.CompanyID) == 0 || len(req.AccountType) == 0 || len(req.BillingType) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// block request if data is already exists
	err = models.CheckAccountDuplicate("", conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, core.DescDataExists, err)
		return response
	}

	// process input
	response.Data, err = processors.InsertAccount(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrServer, core.DescServer, err)
	}

	return response
}

// UpdateAccount is use for
func UpdateAccount(ctx context.Context, req dt.AccountRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("AccountService.UpdateAccount Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.AccountID) == 0 || len(req.Name) == 0 || len(req.Status) == 0 || len(req.CompanyID) == 0 || len(req.AccountType) == 0 || len(req.BillingType) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// block request if old data is not exists
	err = models.CheckAccountExists(req.AccountID, conn)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrNoData, core.DescNoData, err)
		return response
	}

	// block request if data is already exists
	err = models.CheckAccountDuplicate(req.AccountID, conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, err.Error(), err)
		return response
	}

	// process input
	response.Data, err = processors.UpdateAccount(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrServer, core.DescServer, err)
	}

	return response
}

// DeleteAccount is use for
func DeleteAccount(ctx context.Context, req dt.AccountRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("AccountService.DeleteAccount Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.AccountID) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// run
	err = processors.DeleteAccount(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, err.Error(), err)
		return response
	}

	return response
}
