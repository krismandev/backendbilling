package services

import (
	connections "billingdashboard/connections"
	"billingdashboard/core"
	"billingdashboard/lib"
	dt "billingdashboard/modules/account/datastruct"
	"billingdashboard/modules/account/models"
	"billingdashboard/modules/account/processors"
	"context"
	"encoding/json"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

// AccountServices provides operations for endpoint

func ListAccount(ctx context.Context, req dt.AccountRequest, conn *connections.Connections) core.GlobalListResponse {
	log.Infof("AccountService.ListAccount Request : %+v", req)
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
		if len(req.AccountID) > 0 {
			jsonParam := fmt.Sprintf(`{"accountid":"%v"}`, ls.AccountID)
			log.Info("JsonParam-", jsonParam)
			resp, httpStatus := lib.CallRestAPI(conn.ManagementUrl+"/account", "GET", jsonParam, time.Duration(30)*time.Second)
			if httpStatus != 200 {
				core.ErrorGlobalListResponse(&response, core.ErrServer, "Error Call Rest API", err)
				return response
			}
			var tmpAccounts dt.AccountJsonResponse
			err = json.Unmarshal([]byte(resp), &tmpAccounts)
			if tmpAccounts.ResponseData.TotalData > 0 {
				ls.BalanceList = tmpAccounts.ResponseData.Lists[0].BalanceList
			}
		}
		response.Data.List = append(response.Data.List, ls)
	}

	if err != nil {
		core.ErrorGlobalListResponse(&response, core.ErrServer, err.Error(), err)
		return response
	}

	return response
}

// CreateAccount is use for
func CreateAccount(ctx context.Context, req dt.AccountRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("AccountService.CreateAccount Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.Name) == 0 || len(req.Status) == 0 || len(req.AccountType) == 0 || len(req.BillingType) == 0 {
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

// ListStub is use for
func ListRootParentAccount(ctx context.Context, req dt.RootParentAccountRequest, conn *connections.Connections) core.GlobalListResponse {
	log.Infof("CompanyService.ListRootParentAccount Request : %+v", req)
	var response = core.DefaultGlobalListResponse(ctx)
	var err error

	listRootParentAccount, err := processors.GetListRootParentAccount(conn, req)
	if err != nil {
		core.ErrorGlobalListResponse(&response, core.ErrServer, core.DescServer, err)
		return response
	} else {
		response.Data.Page = req.Param.Page
		response.Data.PerPage = req.Param.PerPage
	}

	// append list data as []interface{}
	for _, ls := range listRootParentAccount {
		response.Data.List = append(response.Data.List, ls)
	}

	return response
}

func ListRootAccount(ctx context.Context, req dt.AccountRequest, conn *connections.Connections) core.GlobalListResponse {
	log.Infof("AccouuntService.ListRootAccount Request : %+v", req)
	var response = core.DefaultGlobalListResponse(ctx)
	var err error

	listRootAccount, err := processors.GetListRootAccount(conn, req)
	if err != nil {
		core.ErrorGlobalListResponse(&response, core.ErrServer, core.DescServer, err)
		return response
	} else {
		response.Data.Page = req.Param.Page
		response.Data.PerPage = req.Param.PerPage
	}

	// append list data as []interface{}
	for _, ls := range listRootAccount {
		response.Data.List = append(response.Data.List, ls)
	}

	return response
}
