package routes

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/server/transport"
	"net/http"

	"backendbillingdashboard/core"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func InitRoutes(conn *connections.Connections) {
	ServerRoute(conn)
	ServerAccountRoute(conn)
}

// ServerRoute is used for
func ServerRoute(conn *connections.Connections) {
	serverRoute := mux.NewRouter()
	serverRoute.Methods("GET").Handler(httptransport.NewServer(
		transport.ListServerEndpoint(conn),
		transport.ServerDecodeRequest,
		transport.ServerListEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	serverRoute.Methods("POST").Handler(httptransport.NewServer(
		transport.CreateServerEndpoint(conn),
		transport.ServerDecodeRequest,
		transport.ServerSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	serverRoute.Methods("PATCH").Handler(httptransport.NewServer(
		transport.UpdateServerEndpoint(conn),
		transport.ServerDecodeRequest,
		transport.ServerSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	serverRoute.Methods("DELETE").Handler(httptransport.NewServer(
		transport.DeleteServerEndpoint(conn),
		transport.ServerDecodeRequest,
		transport.ServerSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	http.Handle("/server", serverRoute)
}

func ServerAccountRoute(conn *connections.Connections) {
	serverAccountRoute := mux.NewRouter()
	serverAccountRoute.Methods("GET").Handler(httptransport.NewServer(
		transport.ListServerAccountEndpoint(conn),
		transport.ServerAccountDecodeRequest,
		transport.ServerListEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	http.Handle("/server/account", serverAccountRoute)
}
