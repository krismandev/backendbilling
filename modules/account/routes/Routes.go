package routes

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/account/transport"
	"net/http"

	"backendbillingdashboard/core"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func InitRoutes(conn *connections.Connections) {
	AccountRoute(conn)
	RootParentAccountRoute(conn)
	RootAccountRoute(conn)
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

func RootParentAccountRoute(conn *connections.Connections) {
	rootParentAccountRoute := mux.NewRouter()
	rootParentAccountRoute.Methods("GET").Handler(httptransport.NewServer(
		transport.ListRootParentAccountEndpoint(conn),
		transport.RootParentAccountDecodeRequest,
		transport.RootParentAccountListEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	http.Handle("/account/root-parent", rootParentAccountRoute)
}

func RootAccountRoute(conn *connections.Connections) {
	rootAccountRoute := mux.NewRouter()
	rootAccountRoute.Methods("GET").Handler(httptransport.NewServer(
		transport.ListRootAccountEndpoint(conn),
		transport.AccountDecodeRequest,
		transport.AccountListEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))

	http.Handle("/account/root", rootAccountRoute)
}
