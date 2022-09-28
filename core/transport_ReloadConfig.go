package core

import (
	conf "backendbillingdashboard/config"
	connections "backendbillingdashboard/connections"
	er "backendbillingdashboard/error"
	lib "backendbillingdashboard/lib"
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	log "github.com/sirupsen/logrus"
)

// ReloadConfigDecodeRequest is use for ...
func ReloadConfigDecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	log.WithFields(GetLogFieldValues(ctx, "ReloadConfigDecodeRequest"))
	var request ReloadConfigJSONRequest
	remoteIP := lib.GetRemoteIPAddress(r)
	log.Info("Reload Config Request")
	request.IPAddr = remoteIP
	return request, nil
}

// ReloadConfigEncodeResponse is use for ...
func ReloadConfigEncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	log.WithFields(GetLogFieldValues(ctx, "ReloadConfigEncodeResponse"))
	var body []byte

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Connection", "close")
	body, err := json.Marshal(&response)
	log.Infof("Send Response %s", body[:])
	if err != nil {
		return err
	}

	var e = response.(ReloadConfigJSONResponse).ResponseCode

	if e < 8000 {
		w.WriteHeader(http.StatusOK)
	} else if e < 9000 {
		w.WriteHeader(http.StatusUnauthorized)
	} else if e <= 9999 {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(body)

	return err
}

// ReloadConfigEndpoint is use for
func ReloadConfigEndpoint(svc ReloadConfigServices, conn *connections.Connections) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		log.WithFields(GetLogFieldValues(ctx, "ReloadConfigEndpoint"))
		if conf.Param.UseJWT {
			errNoJWT, errJWT := HandleJWT(ctx)
			if errJWT != nil {
				return ReloadConfigJSONResponse{ResponseCode: errNoJWT, ResponseDesc: errJWT.Error()}, nil
			}
		}

		if req, ok := request.(ReloadConfigJSONRequest); ok {
			return svc.ReloadConfig(ctx, req, conn), nil
		}
		switch request.(type) {
		case *er.AppError:
			{
				if request != nil {
					return ReloadConfigJSONResponse{ResponseCode: request.(*er.AppError).ErrCode, ResponseDesc: request.(*er.AppError).Remark}, nil
				}
			}
		}
		log.Error("Unhandled error occured: request is in unknown format")
		return ReloadConfigJSONResponse{ResponseCode: ErrOthers, ResponseDesc: DescOthers}, nil
	}
}
