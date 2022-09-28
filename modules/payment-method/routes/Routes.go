package routes

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/payment-method/transport"
	"net/http"

	"backendbillingdashboard/core"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func InitRoutes(conn *connections.Connections) {
	PaymentMethodRoute(conn)
}

// PaymentMethodRoute is used for
func PaymentMethodRoute(conn *connections.Connections) {
	paymentMethodRoute := mux.NewRouter()
	paymentMethodRoute.Methods("GET").Handler(httptransport.NewServer(
		transport.ListPaymentMethodEndpoint(conn),
		transport.PaymentMethodDecodeRequest,
		transport.PaymentMethodListEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	paymentMethodRoute.Methods("POST").Handler(httptransport.NewServer(
		transport.CreatePaymentMethodEndpoint(conn),
		transport.PaymentMethodDecodeRequest,
		transport.PaymentMethodSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	paymentMethodRoute.Methods("PATCH").Handler(httptransport.NewServer(
		transport.UpdatePaymentMethodEndpoint(conn),
		transport.PaymentMethodDecodeRequest,
		transport.PaymentMethodSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	paymentMethodRoute.Methods("DELETE").Handler(httptransport.NewServer(
		transport.DeletePaymentMethodEndpoint(conn),
		transport.PaymentMethodDecodeRequest,
		transport.PaymentMethodSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	http.Handle("/payment-method", paymentMethodRoute)
}
