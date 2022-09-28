package transport

import (
	connections "backendbillingdashboard/connections"
	"backendbillingdashboard/core"
	dt "backendbillingdashboard/modules/currency/datastruct"
	"backendbillingdashboard/modules/currency/services"
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	log "github.com/sirupsen/logrus"
)

// ListCurrencyEndpoint is as request middleware
func ListCurrencyEndpoint(conn *connections.Connections) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var errResp core.GlobalListResponse
		errNoJWT, errJWT := core.HandleJWT(ctx)
		if errJWT != nil {
			core.ErrorGlobalListResponse(&errResp, errNoJWT, errJWT.Error(), errJWT)
			return errResp, nil
		}

		if req, ok := request.(dt.CurrencyRequest); ok {
			return services.ListCurrency(ctx, req, conn), nil
		}

		log.Error("Unhandled error occured: request is in unknown format")
		core.ErrorGlobalListResponse(&errResp, core.ErrOthers, core.DescOthers, errors.New("Request is in unknown format"))
		return errResp, nil
	}
}

// CreateCurrencyEndpoint is as request middleware
func CreateCurrencyEndpoint(conn *connections.Connections) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var errResp core.GlobalSingleResponse
		errNoJWT, errJWT := core.HandleJWT(ctx)
		if errJWT != nil {
			core.ErrorGlobalSingleResponse(&errResp, errNoJWT, errJWT.Error(), errJWT)
			return errResp, nil
		}

		if req, ok := request.(dt.CurrencyRequest); ok {
			return services.CreateCurrency(ctx, req, conn), nil
		}

		log.Error("Unhandled error occured: request is in unknown format")
		core.ErrorGlobalSingleResponse(&errResp, core.ErrOthers, core.DescOthers, errors.New("Request is in unknown format"))
		return errResp, nil

	}
}

// UpdateCurrencyEndpoint is as request middleware
func UpdateCurrencyEndpoint(conn *connections.Connections) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var errResp core.GlobalSingleResponse
		errNoJWT, errJWT := core.HandleJWT(ctx)
		if errJWT != nil {
			core.ErrorGlobalSingleResponse(&errResp, errNoJWT, errJWT.Error(), errJWT)
			return errResp, nil
		}

		if req, ok := request.(dt.CurrencyRequest); ok {
			return services.UpdateCurrency(ctx, req, conn), nil
		}

		log.Error("Unhandled error occured: request is in unknown format")
		core.ErrorGlobalSingleResponse(&errResp, core.ErrOthers, core.DescOthers, errors.New("Request is in unknown format"))
		return errResp, nil
	}
}

// DeleteCurrencyEndpoint is as request middleware
func DeleteCurrencyEndpoint(conn *connections.Connections) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var errResp core.GlobalSingleResponse
		errNoJWT, errJWT := core.HandleJWT(ctx)
		if errJWT != nil {
			core.ErrorGlobalSingleResponse(&errResp, errNoJWT, errJWT.Error(), errJWT)
			return errResp, nil
		}

		if req, ok := request.(dt.CurrencyRequest); ok {
			return services.DeleteCurrency(ctx, req, conn), nil
		}

		log.Error("Unhandled error occured: request is in unknown format")
		core.ErrorGlobalSingleResponse(&errResp, core.ErrOthers, core.DescOthers, errors.New("Request is in unknown format"))
		return errResp, nil
	}
}

func ListBalanceEndpoint(conn *connections.Connections) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var errResp core.GlobalListResponse
		errNoJWT, errJWT := core.HandleJWT(ctx)
		if errJWT != nil {
			core.ErrorGlobalListResponse(&errResp, errNoJWT, errJWT.Error(), errJWT)
			return errResp, nil
		}

		if req, ok := request.(dt.BalanceRequest); ok {
			return services.ListBalance(ctx, req, conn), nil
		}

		log.Error("Unhandled error occured: request is in unknown format")
		core.ErrorGlobalListResponse(&errResp, core.ErrOthers, core.DescOthers, errors.New("Request is in unknown format"))
		return errResp, nil
	}
}
