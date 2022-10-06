package services

import (
	connections "backendbillingdashboard/connections"
	"backendbillingdashboard/core"
	dt "backendbillingdashboard/modules/server-data/datastruct"
	"backendbillingdashboard/modules/server-data/models"
	"backendbillingdashboard/modules/server-data/processors"
	"context"
	"time"

	log "github.com/sirupsen/logrus"
)

// ServerDataServices provides operations for endpoint

// ListServerData is use for
func ListServerData(ctx context.Context, req dt.ServerDataRequest, conn *connections.Connections) core.GlobalListResponse {
	log.Infof("ServerDataService.ListServerData Request : %+v", req)
	var response = core.DefaultGlobalListResponse(ctx)
	var err error

	if len(req.CurrencyCode) == 0 || len(req.MonthUse) == 0 {
		core.ErrorGlobalListResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	listServerData, err := processors.GetListServerData(conn, req)
	if err != nil {
		core.ErrorGlobalListResponse(&response, core.ErrServer, core.DescServer, err)
		return response
	} else {
		response.Data.Page = req.Param.Page
		response.Data.PerPage = req.Param.PerPage
	}

	// append list data as []interface{}
	for _, ls := range listServerData {
		response.Data.List = append(response.Data.List, ls)
	}

	return response
}

func LoadServerData(ctx context.Context, req dt.ServerDataRequest, conn *connections.Connections) core.GlobalListResponse {
	log.Infof("ServerDataService.LoadServerData Request : %+v", req)
	var response = core.DefaultGlobalListResponse(ctx)

	// jakartaTime, _ := time.LoadLocation("Asia/Jakarta")
	// scheduler := cron.New(cron.WithLocation(jakartaTime))

	// defer scheduler.Stop()

	// scheduler.AddFunc("*/2 * * * *", TestCron)
	// go scheduler.Start()
	// // trap SIGINT untuk trigger shutdown.
	// sig := make(chan os.Signal, 1)
	// signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	// <-sig
	go func() {
		for true {
			// log.Info("Hello !!")
			models.LoadServerData(conn, req)
			time.Sleep(3600 * time.Second)

		}
	}()
	// ticker := time.NewTicker(1 * time.Second)
	// go func() {
	// 	for range ticker.C {
	// 		fmt.Println("Hello !!")
	// 	}
	// }()

	// // wait for 10 seconds
	// time.Sleep(10 * time.Second)
	// ticker.Stop()
	return response
}

func TestCron() {
	log.Info("TestHallo")
}

// CreateServerData is use for
func CreateServerData(ctx context.Context, req dt.ServerDataRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("ServerDataService.CreateServerData Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.ServerDataID) == 0 || len(req.ServerID) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// block request if data is already exists
	err = models.CheckServerDataDuplicate("", conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, core.DescDataExists, err)
		return response
	}

	// process input
	response.Data, err = processors.InsertServerData(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrServer, core.DescServer, err)
	}

	return response
}

// UpdateServerData is use for
func UpdateServerData(ctx context.Context, req dt.ServerDataRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("ServerDataService.UpdateServerData Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.ServerDataID) == 0 || len(req.ServerID) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// block request if old data is not exists
	err = models.CheckServerDataExists(req.ServerDataID, conn)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrNoData, core.DescNoData, err)
		return response
	}

	// block request if data is already exists
	err = models.CheckServerDataDuplicate(req.ServerDataID, conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, err.Error(), err)
		return response
	}

	// process input
	response.Data, err = processors.UpdateServerData(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrServer, core.DescServer, err)
	}

	return response
}

// DeleteServerData is use for
func DeleteServerData(ctx context.Context, req dt.ServerDataRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("ServerDataService.DeleteServerData Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.ServerDataID) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// run
	err = processors.DeleteServerData(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, err.Error(), err)
		return response
	}

	return response
}
