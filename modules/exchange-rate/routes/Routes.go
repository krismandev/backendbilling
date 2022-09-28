package routes

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/exchange-rate/transport"
	"net/http"

	"backendbillingdashboard/core"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func InitRoutes(conn *connections.Connections) {
	ExchangeRateRoute(conn)
}

// ExchangeRateRoute is used for
func ExchangeRateRoute(conn *connections.Connections) {
	exchangeRateRoute := mux.NewRouter()
	exchangeRateRoute.Methods("GET").Handler(httptransport.NewServer(
		transport.ListExchangeRateEndpoint(conn),
		transport.ExchangeRateDecodeRequest,
		transport.ExchangeRateListEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	exchangeRateRoute.Methods("POST").Handler(httptransport.NewServer(
		transport.CreateExchangeRateEndpoint(conn),
		transport.ExchangeRateDecodeRequest,
		transport.ExchangeRateSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	exchangeRateRoute.Methods("PATCH").Handler(httptransport.NewServer(
		transport.UpdateExchangeRateEndpoint(conn),
		transport.ExchangeRateDecodeRequest,
		transport.ExchangeRateSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	exchangeRateRoute.Methods("DELETE").Handler(httptransport.NewServer(
		transport.DeleteExchangeRateEndpoint(conn),
		transport.ExchangeRateDecodeRequest,
		transport.ExchangeRateSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	http.Handle("/exchange-rate", exchangeRateRoute)
}
