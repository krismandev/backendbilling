package routes

import (
	"billingdashboard/connections"
	"billingdashboard/modules/currency/transport"
	"net/http"

	"billingdashboard/core"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func InitRoutes(conn *connections.Connections) {
	CurrencyRoute(conn)
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
