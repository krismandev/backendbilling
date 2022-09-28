package routes

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/invoice/transport"
	"net/http"

	"backendbillingdashboard/core"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func InitRoutes(conn *connections.Connections) {
	InvoiceRoute(conn)
	CancelInvoiceRoute(conn)
	PrintInvoiceRoute(conn)
}

// InvoiceRoute is used for
func InvoiceRoute(conn *connections.Connections) {
	invoiceRoute := mux.NewRouter()
	invoiceRoute.Methods("GET").Handler(httptransport.NewServer(
		transport.ListInvoiceEndpoint(conn),
		transport.InvoiceDecodeRequest,
		transport.InvoiceListEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	invoiceRoute.Methods("POST").Handler(httptransport.NewServer(
		transport.CreateInvoiceEndpoint(conn),
		transport.InvoiceDecodeRequest,
		transport.InvoiceSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	invoiceRoute.Methods("PATCH").Handler(httptransport.NewServer(
		transport.UpdateInvoiceEndpoint(conn),
		transport.InvoiceDecodeRequest,
		transport.InvoiceSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	invoiceRoute.Methods("DELETE").Handler(httptransport.NewServer(
		transport.DeleteInvoiceEndpoint(conn),
		transport.InvoiceDecodeRequest,
		transport.InvoiceSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	http.Handle("/invoice", invoiceRoute)
}

func CancelInvoiceRoute(conn *connections.Connections) {
	cancelInvoiceRoute := mux.NewRouter()
	cancelInvoiceRoute.Methods("GET").Handler(httptransport.NewServer(
		transport.CancelInvoiceEndpoint(conn),
		transport.CancelInvoiceDecodeRequest,
		transport.InvoiceSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))

	http.Handle("/invoice/cancel", cancelInvoiceRoute)
}

func PrintInvoiceRoute(conn *connections.Connections) {
	printInvoiceRoute := mux.NewRouter()
	printInvoiceRoute.Methods("POST").Handler(httptransport.NewServer(
		transport.PrintInvoiceEndpoint(conn),
		transport.InvoiceDecodeRequest,
		transport.InvoiceSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))

	http.Handle("/invoice/print", printInvoiceRoute)
}
