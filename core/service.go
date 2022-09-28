package core

import (
	connections "backendbillingdashboard/connections"
	"context"

	//"strings"
	conf "backendbillingdashboard/config"

	log "github.com/sirupsen/logrus"
)

// ReloadConfigServices provides operations for endpoint
type ReloadConfigServices interface {
	ReloadConfig(context.Context, ReloadConfigJSONRequest, *connections.Connections) ReloadConfigJSONResponse
}

// ReloadConfigService is use for
type ReloadConfigService struct{}

const ()

//ReloadConfigValidateRequest Validate Draw Prize Request
func ReloadConfigValidateRequest(conn *connections.Connections, req ReloadConfigJSONRequest) *ReloadConfigJSONResponse {
	return nil
}

//ReLoadConfig to load config from config.yml file
func ReLoadConfig(configPath string) {
	conf.LoadConfig(&configPath)
	log.Infof("Configuration Reloaded")
}

// ReloadConfig service is use for
func (ReloadConfigService) ReloadConfig(ctx context.Context, req ReloadConfigJSONRequest, conn *connections.Connections) ReloadConfigJSONResponse {
	//transid:=ctx.Value(ContextTransactionID).(string)
	logger := log.WithFields(GetLogFieldValues(ctx, "ReloadConfig"))
	logger.Info("ReloadConfig Triggered")

	ReLoadConfig("config/config.yml")

	return ReloadConfigJSONResponse{
		ResponseCode: ErrSuccess,
		ResponseDesc: DescSuccess,
		IPAddr:       req.IPAddr,
	}
}
