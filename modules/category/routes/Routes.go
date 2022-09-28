package routes

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/category/transport"
	"net/http"

	"backendbillingdashboard/core"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func InitRoutes(conn *connections.Connections) {
	CategoryRoute(conn)
}

// CategoryRoute is used for
func CategoryRoute(conn *connections.Connections) {
	categoryRoute := mux.NewRouter()
	categoryRoute.Methods("GET").Handler(httptransport.NewServer(
		transport.ListCategoryEndpoint(conn),
		transport.CategoryDecodeRequest,
		transport.CategoryListEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	http.Handle("/category", categoryRoute)
}
