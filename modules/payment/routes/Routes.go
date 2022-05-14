package routes

import (
	"billingdashboard/connections"
	"billingdashboard/modules/payment/transport"
	"net/http"

	"billingdashboard/core"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func InitRoutes(conn *connections.Connections) {
	PaymentRoute(conn)
}

// PaymentRoute is used for
func PaymentRoute(conn *connections.Connections) {
	paymentRoute := mux.NewRouter()
	paymentRoute.Methods("GET").Handler(httptransport.NewServer(
		transport.ListPaymentEndpoint(conn),
		transport.PaymentDecodeRequest,
		transport.PaymentListEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	paymentRoute.Methods("POST").Handler(httptransport.NewServer(
		transport.CreatePaymentEndpoint(conn),
		transport.PaymentDecodeRequest,
		transport.PaymentSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	paymentRoute.Methods("PATCH").Handler(httptransport.NewServer(
		transport.UpdatePaymentEndpoint(conn),
		transport.PaymentDecodeRequest,
		transport.PaymentSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	paymentRoute.Methods("DELETE").Handler(httptransport.NewServer(
		transport.DeletePaymentEndpoint(conn),
		transport.PaymentDecodeRequest,
		transport.PaymentSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	http.Handle("/payment", paymentRoute)
}
