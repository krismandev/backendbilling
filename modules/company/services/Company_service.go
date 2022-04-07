package services

import (
	connections "billingdashboard/connections"
	"billingdashboard/core"
	dt "billingdashboard/modules/company/datastruct"
	"billingdashboard/modules/company/models"
	"billingdashboard/modules/company/processors"
	"context"

	log "github.com/sirupsen/logrus"
)

// CompanyServices provides operations for endpoint

// ListCompany is use for
func ListCompany(ctx context.Context, req dt.CompanyRequest, conn *connections.Connections) core.GlobalListResponse {
	log.Infof("CompanyService.ListCompany Request : %+v", req)
	var response = core.DefaultGlobalListResponse(ctx)
	var err error

	listCompany, err := processors.GetListCompany(conn, req)
	if err != nil {
		core.ErrorGlobalListResponse(&response, core.ErrServer, core.DescServer, err)
		return response
	} else {
		response.Data.Page = req.Param.Page
		response.Data.PerPage = req.Param.PerPage
	}

	// append list data as []interface{}
	for _, ls := range listCompany {
		response.Data.List = append(response.Data.List, ls)
	}

	return response
}

// CreateCompany is use for
func CreateCompany(ctx context.Context, req dt.CompanyRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("CompanyService.CreateCompany Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.Name) == 0 || len(req.Status) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// block request if data is already exists
	err = models.CheckCompanyDuplicate("", conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, core.DescDataExists, err)
		return response
	}

	// process input
	response.Data, err = processors.InsertCompany(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrServer, core.DescServer, err)
	}

	return response
}

// UpdateCompany is use for
func UpdateCompany(ctx context.Context, req dt.CompanyRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("CompanyService.UpdateCompany Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.CompanyID) == 0 || len(req.Name) == 0 || len(req.Status) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// block request if old data is not exists
	err = models.CheckCompanyExists(req.CompanyID, conn)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrNoData, core.DescNoData, err)
		return response
	}

	// block request if data is already exists
	err = models.CheckCompanyDuplicate(req.CompanyID, conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, err.Error(), err)
		return response
	}

	// process input
	response.Data, err = processors.UpdateCompany(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrServer, core.DescServer, err)
	}

	return response
}

// DeleteCompany is use for
func DeleteCompany(ctx context.Context, req dt.CompanyRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("CompanyService.DeleteCompany Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.CompanyID) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// run
	err = processors.DeleteCompany(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, err.Error(), err)
		return response
	}

	return response
}
