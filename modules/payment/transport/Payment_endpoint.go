package transport

import (
	connections "backendbillingdashboard/connections"
	"backendbillingdashboard/core"
	dt "backendbillingdashboard/modules/payment/datastruct"
	"backendbillingdashboard/modules/payment/services"
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	log "github.com/sirupsen/logrus"
)

// ListPaymentEndpoint is as request middleware
func ListPaymentEndpoint(conn *connections.Connections) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var errResp core.GlobalListResponse
		errNoJWT, errJWT := core.HandleJWT(ctx)
		if errJWT != nil {
			core.ErrorGlobalListResponse(&errResp, errNoJWT, errJWT.Error(), errJWT)
			return errResp, nil
		}

		if req, ok := request.(dt.PaymentRequest); ok {
			return services.ListPayment(ctx, req, conn), nil
		}

		log.Error("Unhandled error occured: request is in unknown format")
		core.ErrorGlobalListResponse(&errResp, core.ErrOthers, core.DescOthers, errors.New("Request is in unknown format"))
		return errResp, nil
	}
}

// CreatePaymentEndpoint is as request middleware
func CreatePaymentEndpoint(conn *connections.Connections) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var errResp core.GlobalSingleResponse
		errNoJWT, errJWT := core.HandleJWT(ctx)
		if errJWT != nil {
			core.ErrorGlobalSingleResponse(&errResp, errNoJWT, errJWT.Error(), errJWT)
			return errResp, nil
		}

		if req, ok := request.(dt.PaymentRequest); ok {
			return services.CreatePayment(ctx, req, conn), nil
		}

		log.Error("Unhandled error occured: request is in unknown format")
		core.ErrorGlobalSingleResponse(&errResp, core.ErrOthers, core.DescOthers, errors.New("Request is in unknown format"))
		return errResp, nil

	}
}

// UpdatePaymentEndpoint is as request middleware
func UpdatePaymentEndpoint(conn *connections.Connections) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var errResp core.GlobalSingleResponse
		errNoJWT, errJWT := core.HandleJWT(ctx)
		if errJWT != nil {
			core.ErrorGlobalSingleResponse(&errResp, errNoJWT, errJWT.Error(), errJWT)
			return errResp, nil
		}

		if req, ok := request.(dt.PaymentRequest); ok {
			return services.UpdatePayment(ctx, req, conn), nil
		}

		log.Error("Unhandled error occured: request is in unknown format")
		core.ErrorGlobalSingleResponse(&errResp, core.ErrOthers, core.DescOthers, errors.New("Request is in unknown format"))
		return errResp, nil
	}
}

// DeletePaymentEndpoint is as request middleware
func DeletePaymentEndpoint(conn *connections.Connections) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var errResp core.GlobalSingleResponse
		errNoJWT, errJWT := core.HandleJWT(ctx)
		if errJWT != nil {
			core.ErrorGlobalSingleResponse(&errResp, errNoJWT, errJWT.Error(), errJWT)
			return errResp, nil
		}

		if req, ok := request.(dt.PaymentRequest); ok {
			return services.DeletePayment(ctx, req, conn), nil
		}

		log.Error("Unhandled error occured: request is in unknown format")
		core.ErrorGlobalSingleResponse(&errResp, core.ErrOthers, core.DescOthers, errors.New("Request is in unknown format"))
		return errResp, nil
	}
}

// ListPaymentEndpoint is as request middleware
func ListPaymentDeductionTypeEndpoint(conn *connections.Connections) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var errResp core.GlobalListResponse
		errNoJWT, errJWT := core.HandleJWT(ctx)
		if errJWT != nil {
			core.ErrorGlobalListResponse(&errResp, errNoJWT, errJWT.Error(), errJWT)
			return errResp, nil
		}

		if req, ok := request.(dt.PaymentDeductionTypeRequest); ok {
			return services.ListPaymentDeductionType(ctx, req, conn), nil
		}

		log.Error("Unhandled error occured: request is in unknown format")
		core.ErrorGlobalListResponse(&errResp, core.ErrOthers, core.DescOthers, errors.New("Request is in unknown format"))
		return errResp, nil
	}
}

func ListPaymentDeductionEndpoint(conn *connections.Connections) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var errResp core.GlobalListResponse
		errNoJWT, errJWT := core.HandleJWT(ctx)
		if errJWT != nil {
			core.ErrorGlobalListResponse(&errResp, errNoJWT, errJWT.Error(), errJWT)
			return errResp, nil
		}

		if req, ok := request.(dt.PaymentDeductionRequest); ok {
			return services.ListPaymentDeduction(ctx, req, conn), nil
		}

		log.Error("Unhandled error occured: request is in unknown format")
		core.ErrorGlobalListResponse(&errResp, core.ErrOthers, core.DescOthers, errors.New("Request is in unknown format"))
		return errResp, nil
	}
}

func ListAdjustmentReasonEndpoint(conn *connections.Connections) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var errResp core.GlobalListResponse
		errNoJWT, errJWT := core.HandleJWT(ctx)
		if errJWT != nil {
			core.ErrorGlobalListResponse(&errResp, errNoJWT, errJWT.Error(), errJWT)
			return errResp, nil
		}

		if req, ok := request.(dt.AdjustmentReasonRequest); ok {
			return services.ListAdjustmentReason(ctx, req, conn), nil
		}

		log.Error("Unhandled error occured: request is in unknown format")
		core.ErrorGlobalListResponse(&errResp, core.ErrOthers, core.DescOthers, errors.New("Request is in unknown format"))
		return errResp, nil
	}
}
