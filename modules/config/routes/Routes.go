package routes

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/config/transport"
	"net/http"

	"backendbillingdashboard/core"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func InitRoutes(conn *connections.Connections) {
	ConfigRoute(conn)
}

// ConfigRoute is used for
func ConfigRoute(conn *connections.Connections) {
	configRoute := mux.NewRouter()
	configRoute.Methods("GET").Handler(httptransport.NewServer(
		transport.ListConfigEndpoint(conn),
		transport.ConfigDecodeRequest,
		transport.ConfigListEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))

	http.Handle("/config", configRoute)
}
