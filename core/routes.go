package core

import (
	"backendbillingdashboard/config"
	"backendbillingdashboard/connections"
	"backendbillingdashboard/lib"
	"fmt"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

// InitRoutes handle all routes in this application
func InitRoutes(conn *connections.Connections, version, builddate string) {
	var svcReloadConfig ReloadConfigServices
	svcReloadConfig = ReloadConfigService{}
	ReloadConfigHandler := httptransport.NewServer(
		ReloadConfigEndpoint(svcReloadConfig, conn),
		ReloadConfigDecodeRequest,
		ReloadConfigEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(GetRequestInformation),
	)
	http.Handle("/reload", ReloadConfigHandler)

	http.HandleFunc("/ping", PingHandler)
	http.HandleFunc("/ver", VersionHandler(version, builddate))
	http.HandleFunc("/get-access-token", TokenHandler(conn))

}

// PingHandler Liveness
func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

// VersionHandler for Version
func VersionHandler(version, builddate string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprint("{\"Version\":\"" + version + "\",\"BuildDate\":\"" + builddate + "\"}")))
	}
	return fn
}

// TokenHandler for JWT TOken
func TokenHandler(conn *connections.Connections) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ipaddr := lib.GetRemoteIPAddress(r)
		bearer, err := lib.GenerateToken(config.Param.AppName, ipaddr, conn.JWTSecretKey, 1)
		if err != nil {
			w.Write([]byte(fmt.Sprint("{\"error\":\"" + err.Error() + "\"}")))
		}
		w.Write([]byte(fmt.Sprint("{\"access_token\":\"" + bearer + "\"}")))
	}
	return fn
}
