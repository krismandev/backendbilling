package services

import (
	connections "billingdashboard/connections"
	"billingdashboard/core"
	dt "billingdashboard/modules/item-price/datastruct"
	"billingdashboard/modules/item-price/processors"
	"context"

	log "github.com/sirupsen/logrus"
)

// ItemPriceServices provides operations for endpoint

// ListItemPrice is use for
func ListItemPrice(ctx context.Context, req dt.ItemPriceRequest, conn *connections.Connections) core.GlobalListResponse {
	log.Infof("ItemPriceService.ListItemPrice Request : %+v", req)
	var response = core.DefaultGlobalListResponse(ctx)
	var err error

	listItemPrice, err := processors.GetListItemPrice(conn, req)
	if err != nil {
		core.ErrorGlobalListResponse(&response, core.ErrServer, core.DescServer, err)
		return response
	} else {
		response.Data.Page = req.Param.Page
		response.Data.PerPage = req.Param.PerPage
	}

	// append list data as []interface{}
	for _, ls := range listItemPrice {
		response.Data.List = append(response.Data.List, ls)
	}

	return response
}

// CreateItemPrice is use for
func CreateItemPrice(ctx context.Context, req dt.ItemPriceRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("ItemPriceService.CreateItemPrice Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.ListItemPrice) == 0 {
		if len(req.ItemID) == 0 || len(req.AccountID) == 0 || len(req.ServerID) == 0 || len(req.Price) == 0 {
			core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
			return response
		}
	}

	// block request if data is already exists
	// err = models.CheckItemPriceDuplicate("", conn, req)
	// if err != nil {
	// 	core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, core.DescDataExists, err)
	// 	return response
	// }

	// errCheck := models.CheckItemPriceExists(req.ItemID, conn)
	// if errCheck != nil {
	// 	// core.ErrorGlobalSingleResponse(&response, core.ErrNoData, core.DescNoData, err)
	// 	// return response
	// 	response.Data, err = processors.InsertItemPrice(conn, req)
	// 	if err != nil {
	// 		core.ErrorGlobalSingleResponse(&response, core.ErrServer, core.DescServer, err)
	// 	}
	// } else {
	// 	response.Data, err = processors.UpdateItemPrice(conn, req)
	// 	if err != nil {
	// 		core.ErrorGlobalSingleResponse(&response, core.ErrServer, core.DescServer, err)
	// 	}
	// }

	response.Data, err = processors.InsertItemPrice(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrServer, core.DescServer, err)
	}

	// process input

	return response
}

// UpdateItemPrice is use for
func UpdateItemPrice(ctx context.Context, req dt.ItemPriceRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("ItemPriceService.UpdateItemPrice Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	// if len(req.ItemID) == 0 || len(req.AccountID) == 0 || len(req.ServerID) == 0 || len(req.Price) == 0 {
	// 	core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
	// 	return response
	// }

	if len(req.ListItemPrice) == 0 {
		if len(req.ItemID) == 0 || len(req.AccountID) == 0 || len(req.ServerID) == 0 || len(req.Price) == 0 {
			core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
			return response
		}
	}

	// block request if old data is not exists
	// err = models.CheckItemPriceExists(req.ItemID, conn)
	// if err != nil {
	// 	core.ErrorGlobalSingleResponse(&response, core.ErrNoData, core.DescNoData, err)
	// 	return response
	// }

	// block request if data is already exists
	// err = models.CheckItemPriceDuplicate(req.ItemID, conn, req)
	// if err != nil {
	// 	core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, err.Error(), err)
	// 	return response
	// }

	// process input
	response.Data, err = processors.UpdateItemPrice(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrServer, core.DescServer, err)
	}

	return response
}

// DeleteItemPrice is use for
func DeleteItemPrice(ctx context.Context, req dt.ItemPriceRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("ItemPriceService.DeleteItemPrice Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.ItemID) == 0 || len(req.AccountID) == 0 || len(req.ServerID) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// run
	err = processors.DeleteItemPrice(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, err.Error(), err)
		return response
	}

	return response
}

// UpdateItemPrice is use for
func BulkUpdateItemPrice(ctx context.Context, req dt.ItemPriceRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("ItemPriceService.BulkUpdateItemPrice Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	// if len(req.ItemID) == 0 || len(req.AccountID) == 0 || len(req.ServerID) == 0 || len(req.Price) == 0 {
	// 	core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
	// 	return response
	// }

	if len(req.ListItemPrice) == 0 {
		if len(req.ItemID) == 0 || len(req.AccountID) == 0 || len(req.ServerID) == 0 || len(req.Price) == 0 {
			core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
			return response
		}
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// block request if old data is not exists
	// err = models.CheckItemPriceExists(req.ItemID, conn)
	// if err != nil {
	// 	core.ErrorGlobalSingleResponse(&response, core.ErrNoData, core.DescNoData, err)
	// 	return response
	// }

	// block request if data is already exists
	// err = models.CheckItemPriceDuplicate(req.ItemID, conn, req)
	// if err != nil {
	// 	core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, err.Error(), err)
	// 	return response
	// }

	// process input
	response.Data, err = processors.BulkUpdateItemPrice(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrServer, core.DescServer, err)
	}

	return response
}
