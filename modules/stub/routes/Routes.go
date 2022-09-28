package routes

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/stub/transport"
	"net/http"

	"backendbillingdashboard/core"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func InitRoutes(conn *connections.Connections) {
	StubRoute(conn)
}

// StubRoute is used for
func StubRoute(conn *connections.Connections) {
	stubRoute := mux.NewRouter()
	stubRoute.Methods("GET").Handler(httptransport.NewServer(
		transport.ListStubEndpoint(conn),
		transport.StubDecodeRequest,
		transport.StubListEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	stubRoute.Methods("POST").Handler(httptransport.NewServer(
		transport.CreateStubEndpoint(conn),
		transport.StubDecodeRequest,
		transport.StubSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	stubRoute.Methods("PATCH").Handler(httptransport.NewServer(
		transport.UpdateStubEndpoint(conn),
		transport.StubDecodeRequest,
		transport.StubSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	stubRoute.Methods("DELETE").Handler(httptransport.NewServer(
		transport.DeleteStubEndpoint(conn),
		transport.StubDecodeRequest,
		transport.StubSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	http.Handle("/stub", stubRoute)
}
