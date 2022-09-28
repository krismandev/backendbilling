package services

import (
	connections "backendbillingdashboard/connections"
	"backendbillingdashboard/core"
	dt "backendbillingdashboard/modules/item/datastruct"
	"backendbillingdashboard/modules/item/models"
	"backendbillingdashboard/modules/item/processors"
	"context"

	log "github.com/sirupsen/logrus"
)

// ItemServices provides operations for endpoint

// ListItem is use for
func ListItem(ctx context.Context, req dt.ItemRequest, conn *connections.Connections) core.GlobalListResponse {
	log.Infof("ItemService.ListItem Request : %+v", req)
	var response = core.DefaultGlobalListResponse(ctx)
	var err error

	listItem, err := processors.GetListItem(conn, req)
	if err != nil {
		core.ErrorGlobalListResponse(&response, core.ErrServer, core.DescServer, err)
		return response
	} else {
		response.Data.Page = req.Param.Page
		response.Data.PerPage = req.Param.PerPage
	}

	// append list data as []interface{}
	for _, ls := range listItem {
		response.Data.List = append(response.Data.List, ls)
	}

	return response
}

// CreateItem is use for
func CreateItem(ctx context.Context, req dt.ItemRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("ItemService.CreateItem Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.ItemName) == 0 || len(req.Category) == 0 || len(req.Route) == 0 || len(req.Operator) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// block request if data is already exists
	err = models.CheckItemDuplicate("", conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, "Item is already exist. Please use other route, operator or category", err)
		return response
	}

	// process input
	err = processors.InsertItem(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrServer, core.DescServer, err)
	}

	return response
}

// UpdateItem is use for
func UpdateItem(ctx context.Context, req dt.ItemRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("ItemService.UpdateItem Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.ItemID) == 0 || len(req.ItemName) == 0 || len(req.Route) == 0 || len(req.Operator) == 0 || len(req.Category) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// block request if old data is not exists
	err = models.CheckItemExists(req.ItemID, conn)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrNoData, core.DescNoData, err)
		return response
	}

	// block request if data is already exists
	err = models.CheckItemDuplicate(req.ItemID, conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, err.Error(), err)
		return response
	}

	// process input
	response.Data, err = processors.UpdateItem(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrServer, core.DescServer, err)
	}

	return response
}

// DeleteItem is use for
func DeleteItem(ctx context.Context, req dt.ItemRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("ItemService.DeleteItem Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.ItemID) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// run
	err = processors.DeleteItem(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, err.Error(), err)
		return response
	}

	return response
}
