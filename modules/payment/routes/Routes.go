package routes

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/payment/transport"
	"net/http"

	"backendbillingdashboard/core"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func InitRoutes(conn *connections.Connections) {
	PaymentRoute(conn)
	PaymentDeductionTypeRoute(conn)
	PaymentDeductionRoute(conn)
	AdjustmentReasonRoute(conn)
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

func PaymentDeductionTypeRoute(conn *connections.Connections) {
	paymentDeductionTypeRoute := mux.NewRouter()
	paymentDeductionTypeRoute.Methods("GET").Handler(httptransport.NewServer(
		transport.ListPaymentDeductionTypeEndpoint(conn),
		transport.PaymentDeductionTypeDecodeRequest,
		transport.PaymentDeductionTypeListEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	http.Handle("/payment-deduction-type", paymentDeductionTypeRoute)
}

func PaymentDeductionRoute(conn *connections.Connections) {
	paymentDeductionRoute := mux.NewRouter()
	paymentDeductionRoute.Methods("GET").Handler(httptransport.NewServer(
		transport.ListPaymentDeductionEndpoint(conn),
		transport.PaymentDeductionDecodeRequest,
		transport.PaymentDeductionListEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	http.Handle("/payment-deduction", paymentDeductionRoute)
}

func AdjustmentReasonRoute(conn *connections.Connections) {
	adjustmentReasonRoute := mux.NewRouter()
	adjustmentReasonRoute.Methods("GET").Handler(httptransport.NewServer(
		transport.ListAdjustmentReasonEndpoint(conn),
		transport.AdjustmentReasonDecodeRequest,
		transport.AdjustmentReasonListEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	http.Handle("/adjustment-reason", adjustmentReasonRoute)
}
