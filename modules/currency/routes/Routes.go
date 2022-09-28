package routes

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/currency/transport"
	"net/http"

	"backendbillingdashboard/core"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func InitRoutes(conn *connections.Connections) {
	CurrencyRoute(conn)
	BalanceRoute(conn)
}

// CurrencyRoute is used for
func CurrencyRoute(conn *connections.Connections) {
	currencyRoute := mux.NewRouter()
	currencyRoute.Methods("GET").Handler(httptransport.NewServer(
		transport.ListCurrencyEndpoint(conn),
		transport.CurrencyDecodeRequest,
		transport.CurrencyListEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	currencyRoute.Methods("POST").Handler(httptransport.NewServer(
		transport.CreateCurrencyEndpoint(conn),
		transport.CurrencyDecodeRequest,
		transport.CurrencySingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	currencyRoute.Methods("PATCH").Handler(httptransport.NewServer(
		transport.UpdateCurrencyEndpoint(conn),
		transport.CurrencyDecodeRequest,
		transport.CurrencySingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	currencyRoute.Methods("DELETE").Handler(httptransport.NewServer(
		transport.DeleteCurrencyEndpoint(conn),
		transport.CurrencyDecodeRequest,
		transport.CurrencySingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	http.Handle("/currency", currencyRoute)
}

func BalanceRoute(conn *connections.Connections) {
	BalanceRoute := mux.NewRouter()
	BalanceRoute.Methods("GET").Handler(httptransport.NewServer(
		transport.ListBalanceEndpoint(conn),
		transport.BalanceDecodeRequest,
		transport.BalanceListEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))

	http.Handle("/balance", BalanceRoute)
}
