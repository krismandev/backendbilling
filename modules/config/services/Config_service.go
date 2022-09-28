package services

import (
	connections "backendbillingdashboard/connections"
	"backendbillingdashboard/core"
	dt "backendbillingdashboard/modules/config/datastruct"
	"backendbillingdashboard/modules/config/processors"
	"context"

	log "github.com/sirupsen/logrus"
)

// ConfigServices provides operations for endpoint

// ListConfig is use for
func ListConfig(ctx context.Context, req dt.ConfigRequest, conn *connections.Connections) core.GlobalListResponse {
	log.Infof("ConfigService.ListConfig Request : %+v", req)
	var response = core.DefaultGlobalListResponse(ctx)
	var err error

	listConfig, err := processors.GetListConfig(conn, req)
	if err != nil {
		core.ErrorGlobalListResponse(&response, core.ErrServer, core.DescServer, err)
		return response
	} else {
		response.Data.Page = req.Param.Page
		response.Data.PerPage = req.Param.PerPage
	}

	// append list data as []interface{}
	for _, ls := range listConfig {
		response.Data.List = append(response.Data.List, ls)
	}

	return response
}

// CreateConfig is use for
