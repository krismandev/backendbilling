package routes

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/server-account/transport"
	"net/http"

	"backendbillingdashboard/core"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func InitRoutes(conn *connections.Connections) {
	ServerAccountRoute(conn)
}

// ServerAccountRoute is used for
func ServerAccountRoute(conn *connections.Connections) {
	ServerAccountRoute := mux.NewRouter()
	ServerAccountRoute.Methods("GET").Handler(httptransport.NewServer(
		transport.ListServerAccountEndpoint(conn),
		transport.ServerAccountDecodeRequest,
		transport.ServerAccountListEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	ServerAccountRoute.Methods("POST").Handler(httptransport.NewServer(
		transport.CreateServerAccountEndpoint(conn),
		transport.ServerAccountDecodeRequest,
		transport.ServerAccountSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	ServerAccountRoute.Methods("PATCH").Handler(httptransport.NewServer(
		transport.UpdateServerAccountEndpoint(conn),
		transport.ServerAccountDecodeRequest,
		transport.ServerAccountSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	ServerAccountRoute.Methods("DELETE").Handler(httptransport.NewServer(
		transport.DeleteServerAccountEndpoint(conn),
		transport.ServerAccountDecodeRequest,
		transport.ServerAccountSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	http.Handle("/server-account", ServerAccountRoute)
}
