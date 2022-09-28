package routes

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/invoice-type/transport"
	"net/http"

	"backendbillingdashboard/core"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func InitRoutes(conn *connections.Connections) {
	InvoiceTypeRoute(conn)
}

// InvoiceTypeRoute is used for
func InvoiceTypeRoute(conn *connections.Connections) {
	invoiceTypeRoute := mux.NewRouter()
	invoiceTypeRoute.Methods("GET").Handler(httptransport.NewServer(
		transport.ListInvoiceTypeEndpoint(conn),
		transport.InvoiceTypeDecodeRequest,
		transport.InvoiceTypeListEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	invoiceTypeRoute.Methods("POST").Handler(httptransport.NewServer(
		transport.CreateInvoiceTypeEndpoint(conn),
		transport.InvoiceTypeDecodeRequest,
		transport.InvoiceTypeSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	invoiceTypeRoute.Methods("PATCH").Handler(httptransport.NewServer(
		transport.UpdateInvoiceTypeEndpoint(conn),
		transport.InvoiceTypeDecodeRequest,
		transport.InvoiceTypeSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	invoiceTypeRoute.Methods("DELETE").Handler(httptransport.NewServer(
		transport.DeleteInvoiceTypeEndpoint(conn),
		transport.InvoiceTypeDecodeRequest,
		transport.InvoiceTypeSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	http.Handle("/invoice-type", invoiceTypeRoute)
}
