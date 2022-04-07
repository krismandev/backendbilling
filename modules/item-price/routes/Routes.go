package routes

import (
	"billingdashboard/connections"
	"billingdashboard/modules/item-price/transport"
	"net/http"

	"billingdashboard/core"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func InitRoutes(conn *connections.Connections) {
	ItemPriceRoute(conn)
}

// ItemPriceRoute is used for
func ItemPriceRoute(conn *connections.Connections) {
	stubRoute := mux.NewRouter()
	stubRoute.Methods("GET").Handler(httptransport.NewServer(
		transport.ListItemPriceEndpoint(conn),
		transport.ItemPriceDecodeRequest,
		transport.ItemPriceListEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	stubRoute.Methods("POST").Handler(httptransport.NewServer(
		transport.CreateItemPriceEndpoint(conn),
		transport.ItemPriceDecodeRequest,
		transport.ItemPriceSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	stubRoute.Methods("PATCH").Handler(httptransport.NewServer(
		transport.UpdateItemPriceEndpoint(conn),
		transport.ItemPriceDecodeRequest,
		transport.ItemPriceSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	stubRoute.Methods("DELETE").Handler(httptransport.NewServer(
		transport.DeleteItemPriceEndpoint(conn),
		transport.ItemPriceDecodeRequest,
		transport.ItemPriceSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	http.Handle("/item-price", stubRoute)
}
