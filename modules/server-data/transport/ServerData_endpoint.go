package transport

import (
	connections "backendbillingdashboard/connections"
	"backendbillingdashboard/core"
	dt "backendbillingdashboard/modules/server-data/datastruct"
	"backendbillingdashboard/modules/server-data/services"
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	log "github.com/sirupsen/logrus"
)

// ListServerDataEndpoint is as request middleware
func ListServerDataEndpoint(conn *connections.Connections) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var errResp core.GlobalListResponse
		errNoJWT, errJWT := core.HandleJWT(ctx)
		if errJWT != nil {
			core.ErrorGlobalListResponse(&errResp, errNoJWT, errJWT.Error(), errJWT)
			return errResp, nil
		}

		if req, ok := request.(dt.ServerDataRequest); ok {
			return services.ListServerData(ctx, req, conn), nil
		}

		log.Error("Unhandled error occured: request is in unknown format")
		core.ErrorGlobalListResponse(&errResp, core.ErrOthers, core.DescOthers, errors.New("Request is in unknown format"))
		return errResp, nil
	}
}

// LoadServerDataEndpoint is as request middleware
func LoadServerDataEndpoint(conn *connections.Connections) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var errResp core.GlobalListResponse
		errNoJWT, errJWT := core.HandleJWT(ctx)
		if errJWT != nil {
			core.ErrorGlobalListResponse(&errResp, errNoJWT, errJWT.Error(), errJWT)
			return errResp, nil
		}

		if req, ok := request.(dt.ServerDataRequest); ok {
			return services.LoadServerData(ctx, req, conn), nil
		}

		log.Error("Unhandled error occured: request is in unknown format")
		core.ErrorGlobalListResponse(&errResp, core.ErrOthers, core.DescOthers, errors.New("Request is in unknown format"))
		return errResp, nil
	}
}

// CreateServerDataEndpoint is as request middleware
func CreateServerDataEndpoint(conn *connections.Connections) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var errResp core.GlobalSingleResponse
		errNoJWT, errJWT := core.HandleJWT(ctx)
		if errJWT != nil {
			core.ErrorGlobalSingleResponse(&errResp, errNoJWT, errJWT.Error(), errJWT)
			return errResp, nil
		}

		if req, ok := request.(dt.ServerDataRequest); ok {
			return services.CreateServerData(ctx, req, conn), nil
		}

		log.Error("Unhandled error occured: request is in unknown format")
		core.ErrorGlobalSingleResponse(&errResp, core.ErrOthers, core.DescOthers, errors.New("Request is in unknown format"))
		return errResp, nil

	}
}

// UpdateServerDataEndpoint is as request middleware
func UpdateServerDataEndpoint(conn *connections.Connections) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var errResp core.GlobalSingleResponse
		errNoJWT, errJWT := core.HandleJWT(ctx)
		if errJWT != nil {
			core.ErrorGlobalSingleResponse(&errResp, errNoJWT, errJWT.Error(), errJWT)
			return errResp, nil
		}

		if req, ok := request.(dt.ServerDataRequest); ok {
			return services.UpdateServerData(ctx, req, conn), nil
		}

		log.Error("Unhandled error occured: request is in unknown format")
		core.ErrorGlobalSingleResponse(&errResp, core.ErrOthers, core.DescOthers, errors.New("Request is in unknown format"))
		return errResp, nil
	}
}

// DeleteServerDataEndpoint is as request middleware
func DeleteServerDataEndpoint(conn *connections.Connections) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var errResp core.GlobalSingleResponse
		errNoJWT, errJWT := core.HandleJWT(ctx)
		if errJWT != nil {
			core.ErrorGlobalSingleResponse(&errResp, errNoJWT, errJWT.Error(), errJWT)
			return errResp, nil
		}

		if req, ok := request.(dt.ServerDataRequest); ok {
			return services.DeleteServerData(ctx, req, conn), nil
		}

		log.Error("Unhandled error occured: request is in unknown format")
		core.ErrorGlobalSingleResponse(&errResp, core.ErrOthers, core.DescOthers, errors.New("Request is in unknown format"))
		return errResp, nil
	}
}
