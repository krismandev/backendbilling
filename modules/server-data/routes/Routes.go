package routes

import (
	"billingdashboard/connections"
	"billingdashboard/modules/server-data/transport"
	"net/http"

	"billingdashboard/core"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func InitRoutes(conn *connections.Connections) {
	ServerDataRoute(conn)
}

// ServerDataRoute is used for
func ServerDataRoute(conn *connections.Connections) {
	serverDataRoute := mux.NewRouter()
	serverDataRoute.Methods("GET").Handler(httptransport.NewServer(
		transport.ListServerDataEndpoint(conn),
		transport.ServerDataDecodeRequest,
		transport.ServerDataListEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	serverDataRoute.Methods("POST").Handler(httptransport.NewServer(
		transport.CreateServerDataEndpoint(conn),
		transport.ServerDataDecodeRequest,
		transport.ServerDataSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	serverDataRoute.Methods("PATCH").Handler(httptransport.NewServer(
		transport.UpdateServerDataEndpoint(conn),
		transport.ServerDataDecodeRequest,
		transport.ServerDataSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	serverDataRoute.Methods("DELETE").Handler(httptransport.NewServer(
		transport.DeleteServerDataEndpoint(conn),
		transport.ServerDataDecodeRequest,
		transport.ServerDataSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	http.Handle("/server-data", serverDataRoute)
}
