package services

import (
	connections "backendbillingdashboard/connections"
	"backendbillingdashboard/core"
	dt "backendbillingdashboard/modules/category/datastruct"
	"backendbillingdashboard/modules/category/processors"
	"context"

	log "github.com/sirupsen/logrus"
)

// CategoryServices provides operations for endpoint

// ListCategory is use for
func ListCategory(ctx context.Context, req dt.CategoryRequest, conn *connections.Connections) core.GlobalListResponse {
	log.Infof("CategoryService.ListCategory Request : %+v", req)
	var response = core.DefaultGlobalListResponse(ctx)
	var err error

	listCategory, err := processors.GetListCategory(conn, req)
	if err != nil {
		core.ErrorGlobalListResponse(&response, core.ErrServer, core.DescServer, err)
		return response
	} else {
		response.Data.Page = req.Param.Page
		response.Data.PerPage = req.Param.PerPage
	}

	// append list data as []interface{}
	for _, ls := range listCategory {
		response.Data.List = append(response.Data.List, ls)
	}

	return response
}

// CreateCategory is use for
