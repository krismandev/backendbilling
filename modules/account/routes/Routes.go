package routes

import (
	"billingdashboard/connections"
	"billingdashboard/modules/account/transport"
	"net/http"

	"billingdashboard/core"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func InitRoutes(conn *connections.Connections) {
	AccountRoute(conn)
}

// AccountRoute is used for
func AccountRoute(conn *connections.Connections) {
	accountRoute := mux.NewRouter()
	accountRoute.Methods("GET").Handler(httptransport.NewServer(
		transport.ListAccountEndpoint(conn),
		transport.AccountDecodeRequest,
		transport.AccountListEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	accountRoute.Methods("POST").Handler(httptransport.NewServer(
		transport.CreateAccountEndpoint(conn),
		transport.AccountDecodeRequest,
		transport.AccountSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	accountRoute.Methods("PATCH").Handler(httptransport.NewServer(
		transport.UpdateAccountEndpoint(conn),
		transport.AccountDecodeRequest,
		transport.AccountSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	accountRoute.Methods("DELETE").Handler(httptransport.NewServer(
		transport.DeleteAccountEndpoint(conn),
		transport.AccountDecodeRequest,
		transport.AccountSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	http.Handle("/account", accountRoute)
}
