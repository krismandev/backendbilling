package routes

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/item/transport"
	"net/http"

	"backendbillingdashboard/core"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func InitRoutes(conn *connections.Connections) {
	ItemRoute(conn)
}

// ItemRoute is used for
func ItemRoute(conn *connections.Connections) {
	itemRoute := mux.NewRouter()
	itemRoute.Methods("GET").Handler(httptransport.NewServer(
		transport.ListItemEndpoint(conn),
		transport.ItemDecodeRequest,
		transport.ItemListEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	itemRoute.Methods("POST").Handler(httptransport.NewServer(
		transport.CreateItemEndpoint(conn),
		transport.ItemDecodeRequest,
		transport.ItemSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	itemRoute.Methods("PATCH").Handler(httptransport.NewServer(
		transport.UpdateItemEndpoint(conn),
		transport.ItemDecodeRequest,
		transport.ItemSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	itemRoute.Methods("DELETE").Handler(httptransport.NewServer(
		transport.DeleteItemEndpoint(conn),
		transport.ItemDecodeRequest,
		transport.ItemSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))

	http.Handle("/item", itemRoute)
}
